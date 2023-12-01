package gfx

// This file implements a polygon filling algorithm. Originally written for
// drawing lines, but it should be extensible to any arbitrary polygon shape.
//
// This algorithm uses a lot of fixed-point math. To indicate how many
// fractional bits each value has, a comment is added. For example:
//
//     var foo = x2 << 4 // .4
//
// ...means that 'foo' has 4 fractional bits.

import (
	"math/bits"

	"tinygo.org/x/drivers/pixel"
)

// Polygon shape. This is a struct, for convenience.
type polygon struct {
	edges            []polygonEdge
	boundX1, boundY1 int16
	boundX2, boundY2 int16
}

// Polygon edge. You can also think of this as a trapezoid, that starts at the
// edge on the left and extends to the right until it's past the bounding box.
type polygonEdge struct {
	// ytop and ybottom indicate the horizontal start and end lines of the
	// polygon, as in [ytop, ybottom). Each horizontal scanline contains 4 lines
	// for antialiasing, so the coordinates use 2 fractional bits.
	ytop    int32 // .2
	ybottom int32 // .2

	// xstart is the X coordinate at the ytop sub-scanline.
	xstart int32 // .16

	// xinc is how much X increases for each Y pixel. That is, X at a certain
	// scanline can be calculated using:
	//
	//   xstart + xinc*(y-ytop)
	xinc int32 // .16
}

// xleft returns the leftmost whole pixel that covers the polygon edge at the
// given Y scanline (for the whole scanline height).
func (e polygonEdge) xleft(y int) int32 {
	edgeY := y*4 - int(e.ytop)
	if e.xinc > 0 {
		return (e.xstart + int32(edgeY)*e.xinc/4) >> 16 // .0
	} else {
		return (e.xstart + int32(edgeY+3)*e.xinc/4) >> 16 // .0
	}
}

// xright returns the rightmost pixel that covers the polygon edge at the given
// Y scanline (for the whole scanline height).
// To be clear, this is done while viewing the edge as an edge, not as a
// trapezoid (in which case xright wouldn't make much sense).
func (e polygonEdge) xright(y int) int32 {
	edgeY := y*4 - int(e.ytop)
	if e.xinc > 0 {
		return (e.xstart + int32(edgeY+3)*e.xinc/4) >> 16 // .0
	} else {
		return (e.xstart + int32(edgeY)*e.xinc/4) >> 16 // .0
	}
}

// inrange returns whether the polygon edge covers the given scanline Y (even if
// only partially). This can be used to quickly skip over edges that are not
// relevant to calculating the pixels at a given scanline.
func (e polygonEdge) inrange(y int) bool {
	return y >= int(e.ytop/4) && y < int((e.ybottom+3)/4)
}

// drawPolygon draws the polygon at the given coordinates into the given image.
func drawPolygon[T pixel.Color](p *polygon, img pixel.Image[T], imgX, imgY int, color T) {
	// This is a scanline algorithm, so it considers every horizontal line
	// independenty. Which is convenient, as it's easy to clip the polygon to
	// the (vertical) bounding box of the image.

	// Clip the polygon bounding box to the image bounding box.
	imgWidth, imgHeight := img.Size()
	ystart := imgY
	if ystart < int(p.boundY1) {
		ystart = int(p.boundY1)
	}
	yend := imgY + imgHeight
	if yend >= int(p.boundY2) {
		yend = int(p.boundY2)
	}
	xstart := imgX
	if xstart < int(p.boundX1) {
		xstart = int(p.boundX1)
	}
	xend := imgX + imgWidth
	if xend >= int(p.boundX2) {
		xend = int(p.boundX2)
	}

	// Iterate over every scanline (y is the scanline).
	for y := ystart; y < yend; y++ {
		// Determine the first edge that crosses the scanline.
		edgeStart := 0
		for edgeStart < len(p.edges) && !p.edges[edgeStart].inrange(y) {
			edgeStart++
		}
		if edgeStart >= len(p.edges) {
			// Apparently this scanline doesn't cross any edges, which means we
			// can leave the whole line empty.
			continue
		}

		// Determine the last X pixel coordinate this leftmost edge crosses.
		// We precalculate this so that we can more quickly check whether to
		// increase edgeStart later on. So essentially, edgeStartRight is the
		// cached value of p.edges[edgeStart].xright(y).
		edgeStartRight := p.edges[edgeStart].xright(y)

		// Set the initial value of the rightmost edge that crosses the scanline
		// at a given pixel X coordinate. This will be incremented on the first
		// iteration inside the scanline.
		edgeEnd := edgeStart
		edgeEndLeft := p.edges[edgeEnd].xleft(y)

		// The active mask that we'll use for the scanline. It is updated with
		// each pixel we calculate for the scanline.
		// It is an 8x4 grid of subpixels, where each subpixel is a single bit.
		// For more information, take a look at the A-buffer algorithm:
		// https://bucior.com/antialising-polygon-edges-for-scanline-rendering/
		mask := uint32(0)

		// Go through the scanline from left to right.
		// Note: X is incremented manually at the end of the loop, to better
		// handle long stretches of solid color pixels.
		for x := int(p.edges[edgeStart].xleft(y)); x < xend; {
			// Update edgeStart if needed: make sure edgeStart points to the
			// first edge that crosses the current pixel.
			for edgeStartRight < int32(x) {
				edgeStart++
				for edgeStart < len(p.edges) && !p.edges[edgeStart].inrange(y) {
					edgeStart++
				}
				if edgeStart < len(p.edges) {
					edgeStartRight = p.edges[edgeStart].xright(y)
				} else {
					edgeStartRight = 0x7fff_ffff
				}
			}
			if edgeStart >= len(p.edges) {
				// We're past the last edge on this scanline, so stop drawing.
				break
			}

			// Update edgeEnd if needed: make sure edgeEnd points to the last
			// edge that crosses the current pixel.
			for edgeEndLeft <= int32(x) {
				edgeEnd++
				for edgeEnd < len(p.edges) && !p.edges[edgeEnd].inrange(y) {
					edgeEnd++
				}
				if edgeEnd < len(p.edges) {
					edgeEndLeft = p.edges[edgeEnd].xleft(y)
				} else {
					edgeEndLeft = 0x7fff_ffff
				}
			}

			// Update the mask according to every edge that crosses the current
			// pixel in the scanline.
			for i := edgeStart; i < edgeEnd; i++ {
				edge := p.edges[i]
				if !edge.inrange(y) {
					// It's possible that there are edges that are between the
					// start and end edge that do not cross the scanline.
					// Ignore these.
					continue
				}

				// Calculate the intersection of the edge at four sub-scanlines.
				// This enables antialiasing in the vertical direction (though
				// not as well as antialiasing in the vertical direction which
				// uses 8 subpixel positions instead of 4).
				for subline := 0; subline < 4; subline++ {
					// Calculate the number of sub-scanlines from the top of the
					// edge/trapezoid.
					sublineY := int32(y*4+subline) - edge.ytop // .2
					if uint32(sublineY) >= uint32(edge.ybottom-edge.ytop) {
						// Edge doesn't intersect sub-scanline so ignore.
						continue
					}

					// Calculate edge intersection with the sub-scanline.
					xoffset := edge.xstart + sublineY*edge.xinc/4

					// Calculate distance from the left side of the current
					// pixel.
					bitOffsetX := (xoffset - int32(x)<<16) >> 13 // .3
					if uint32(bitOffsetX) >= 8 {
						// The intersection is outside the pixel we're interested
						// in.
						continue
					}
					// bitOffsetX is now in the range 0..7

					// Calculate the 8 subpixel bits in one go.
					bits := (uint32(1) << (8 - bitOffsetX)) - 1

					// XOR the horizontal subpixel mask (8 bits) with the
					// current mask.
					mask ^= bits << (subline * 8)
				}
			}

			if mask != 0 {
				if mask == 0xffff_ffff {
					// Special case: efficiently paint a row of pixels inside
					// the polygon.
					pixelXStart := x
					if pixelXStart < xstart {
						pixelXStart = xstart
					}
					pixelXEnd := int(edgeEndLeft)
					if pixelXEnd > xend {
						pixelXEnd = xend
					}
					drawLine(img, pixelXStart-imgX, pixelXEnd-imgX, y-imgY, color)
					x = int(edgeEndLeft)
					continue
				}

				// Calculate the alpha channel by counting the number of set
				// bits in the mask.
				alpha := bits.OnesCount32(mask) << 3 // 0..255 (actualy 0..248)
				blendPixel(img, x-imgX, y-imgY, color, uint8(alpha))

				// Extend the subpixels that are at the right of the pixel mask
				// to the whole subpixel line using bitwise tricks.
				mask = (mask & 0x0101_0101) * 255
			}

			// Loop increment.
			// This is done here (instead of in the loop header) so that we can
			// easily skip ranges of pixels.
			x++
		}
	}
}

// Update bounding box of this polygon. Must be called after changing the
// polygon shape.
func (p *polygon) updateBounds() {
	var polygonX1 int32 = 0x7fff_ffff
	var polygonY1 int32 = 0x7fff_ffff
	var polygonX2 int32 = -0x8000_0000
	var polygonY2 int32 = -0x8000_0000
	for _, edge := range p.edges {
		ytop := edge.ytop / 4
		if ytop < polygonY1 {
			polygonY1 = ytop
		}
		ybottom := (edge.ybottom + 3) / 4
		if ybottom > polygonY2 {
			polygonY2 = ybottom
		}
		edgeX1 := edge.xstart
		edgeX2 := edge.xstart + (ybottom-ytop)*edge.xinc
		if edgeX1 > edgeX2 {
			edgeX2, edgeX1 = edgeX1, edgeX2
		}
		if edgeX1 < polygonX1 {
			polygonX1 = edgeX1
		}
		if edgeX2 > polygonX2 {
			polygonX2 = edgeX2
		}
	}
	p.boundX1 = int16(polygonX1 >> 16)
	p.boundY1 = int16(polygonY1)
	p.boundX2 = int16((polygonX2 + 0xffff) >> 16)
	p.boundY2 = int16(polygonY2)
}
