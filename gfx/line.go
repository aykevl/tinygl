package gfx

import (
	"tinygo.org/x/drivers/pixel"
)

// Line is a line from a start coordinate to an end coordinate.
type Line[T pixel.Color] struct {
	baseObject[T]
	x2, y2      int16 // x1, y1 are x, y
	strokeWidth int16 // line width in pixels
	color       T
	hidden      bool
	polygon     polygon
	edges       [4]polygonEdge
}

// NewLine creates a new line object, with antialiased edges.
//
// The two coordinates (x1, y1) and (x2, y2) describe the starting and ending
// coordinate. The strokeWidth describes the width of the line, like the SVG
// property. Lines by default have an ending like SVG stroke-linecap="butt".
func NewLine[T pixel.Color](color T, x1, y1, x2, y2, strokeWidth int) *Line[T] {
	line := &Line[T]{
		baseObject: baseObject[T]{
			x: int16(x1),
			y: int16(y1),
		},
		x2:          int16(x2),
		y2:          int16(y2),
		strokeWidth: int16(strokeWidth),
		color:       color,
	}
	line.polygon.edges = line.edges[:]
	line.updatePolygon()
	return line
}

// Update the polygon that's used to render the line.
// This must be called after every change, like updating the coordinates or
// updating the line width.
func (obj *Line[T]) updatePolygon() {
	thickness := int(obj.strokeWidth)
	if thickness == 0 {
		// The line has a width of 0, which means it's not visible.
		obj.polygon.boundX1 = 0
		obj.polygon.boundX2 = 0
		obj.polygon.boundY1 = 0
		obj.polygon.boundY2 = 0
		return
	}

	x1 := int(obj.x)
	y1 := int(obj.y)
	x2 := int(obj.x2)
	y2 := int(obj.y2)

	if y1 == y2 {
		// Straight horizontal line.
		if x1 > x2 {
			x2, x1 = x1, x2
		}
		obj.polygon.boundX1 = int16(x1)
		obj.polygon.boundX2 = int16(x2 + 1)
		obj.polygon.boundY1 = int16(y1 - thickness/2)
		obj.polygon.boundY2 = obj.polygon.boundY1 + int16(thickness)
		return
	}

	if x1 == x2 {
		// Straight vertical line.
		if y1 > y2 {
			y2, y1 = y1, y2
		}
		obj.polygon.boundX1 = int16(x1 - thickness/2)
		obj.polygon.boundX2 = obj.polygon.boundX1 + int16(thickness)
		obj.polygon.boundY1 = int16(y1)
		obj.polygon.boundY2 = int16(y2 + 1)
		return
	}

	// Make sure the line is pointing downwards.
	if y1 > y2 {
		y2, y1 = y1, y2
		x2, x1 = x1, x2
	}

	// The following code makes heavy use of fixed point coordinates.
	// To make it easier to follow the number of fractional bits, the number of
	// fractional bits are given in a comment. For example, in this example:
	//
	//     var foo = x2 << 4 // .4
	//
	// the variable foo contains x2 with 4 fractional bits.

	// TODO: fix the line start/end coordinates.
	// Right now the start/end coordinate is assumed to be the middle of the
	// start/end pixel, which means it's not consistent with the way straight
	// lines are drawn. Either we should draw straight lines like we draw
	// non-straight lines (with fuzzy edges), or we should extend the range of
	// straight lines so that the start/end pixel are fully colored.
	//
	// A particularly interesting case is lines with a lineWidth of 1, where
	// users probably expect the first and last pixel to be fully (or mostly)
	// drawn instead of only being drawn for around 50%.

	xslope := int32((x2 - x1) << 16 / (y2 - y1)) // .16
	yslope := int32((y2 - y1) << 16 / (x2 - x1)) // .16

	// Determine the length of the line, using sqrt(w*w + h*h).
	length := int(isqrt32(uint32((x2-x1)*(x2-x1)*16 + (y2-y1)*(y2-y1)*16))) // .2

	// Split the line thickness into the X and Y components.
	thickness_w := thickness * (y2 - y1) * 1024 / length // .8
	thickness_h := thickness * (x2 - x1) * 1024 / length // .8
	thickness_h_abs := thickness_h
	if thickness_h_abs < 0 {
		thickness_h_abs = -thickness_h_abs
	}

	// Determine the Y coordinates that will be used for sampling sub-scanlines.
	// There are four sub-scanlines per regular scanline.
	//
	// TODO: we should sample the polygon not from the (0, 0) position within
	// the pixel, but from X=1/16 and Y=1/8. Otherwise we're not really sampling
	// the pixel, but a small offset from within the pixel.
	var (
		y1top    = int32(y1<<16+1<<15-thickness_h_abs<<7) >> 14 // .2
		y1bottom = int32(y1<<16+1<<15+thickness_h_abs<<7) >> 14 // .2
		y2top    = int32(y2<<16+1<<15-thickness_h_abs<<7) >> 14 // .2
		y2bottom = int32(y2<<16+1<<15+thickness_h_abs<<7) >> 14 // .2
	)

	// Define the polygon itself.
	if thickness_h > 0 {
		// A line that goes down and to the right.
		topEdge := polygonEdge{
			ytop:    y1top,
			ybottom: y1bottom,
			xstart:  int32(x1<<16 + thickness_w<<7 + 1<<15),
			xinc:    -yslope,
		}
		leftEdge := polygonEdge{
			ytop:    y1bottom,
			ybottom: y2bottom,
			xstart:  int32(x1<<16 - thickness_w<<7 + 1<<15),
			xinc:    xslope,
		}
		rightEdge := polygonEdge{
			ytop:    y1top,
			ybottom: y2top,
			xstart:  int32(x1<<16 + thickness_w<<7 + 1<<15),
			xinc:    xslope,
		}
		bottomEdge := polygonEdge{
			ytop:    y2top,
			ybottom: y2bottom,
			xstart:  int32(x2<<16 + thickness_w<<7 + 1<<15),
			xinc:    -yslope,
		}
		obj.edges = [4]polygonEdge{leftEdge, topEdge, rightEdge, bottomEdge}
	} else {
		// A line that goes down and to the left.
		topEdge := polygonEdge{
			ytop:    y1top,
			ybottom: y1bottom,
			xstart:  int32(x1<<16 - thickness_w<<7 + 1<<15),
			xinc:    -yslope,
		}
		leftEdge := polygonEdge{
			ytop:    y1top,
			ybottom: y2top,
			xstart:  int32(x1<<16 - thickness_w<<7 + 1<<15),
			xinc:    xslope,
		}
		rightEdge := polygonEdge{
			ytop:    y1bottom,
			ybottom: y2bottom,
			xstart:  int32(x1<<16 + thickness_w<<7 + 1<<15),
			xinc:    xslope,
		}
		bottomEdge := polygonEdge{
			ytop:    y2top,
			ybottom: y2bottom,
			xstart:  int32(x2<<16 - thickness_w<<7 + 1<<15),
			xinc:    -yslope,
		}
		obj.edges = [4]polygonEdge{leftEdge, bottomEdge, topEdge, rightEdge}
	}

	obj.polygon.updateBounds()
}

// Draw implements the gfx.Object interface.
func (obj *Line[T]) Draw(imgX, imgY int, img pixel.Image[T]) {
	if obj.hidden {
		return
	}

	thickness := int(obj.strokeWidth)
	if thickness == 0 {
		// Nothing to paint.
		return
	}

	x1 := int(obj.x) - imgX
	y1 := int(obj.y) - imgY
	x2 := int(obj.x2) - imgX
	y2 := int(obj.y2) - imgY

	if y1 == y2 {
		// Fast path: draw horizontal line.
		if x1 > x2 {
			x2, x1 = x1, x2
		}
		drawRect(img, obj.color, x1, y1-thickness/2, x2-x1+1, thickness)
		return
	}

	if x1 == x2 {
		// Fast path: draw vertical line.
		if y1 > y2 {
			y2, y1 = y1, y2
		}
		drawRect(img, obj.color, x1-thickness/2, y1, thickness, y2-y1+1)
		return
	}

	// Draw the full polygon.
	drawPolygon(&obj.polygon, img, imgX, imgY, obj.color)
}

func (obj *Line[T]) markDirty() {
	thickness := int(obj.strokeWidth)
	if thickness == 0 {
		// The line is invisible.
		return
	}

	// TODO: this invalidates the entire bounding box of the line.
	// It's possible to optimize this, essentially by running the polygon
	// filling algorithm algorithm again but at blockSize-sized pixels (the
	// canvas tile size). But for now, this works.
	obj.canvas.markDirty(int(obj.polygon.boundX1), int(obj.polygon.boundY1), int(obj.polygon.boundX2-obj.polygon.boundX1), int(obj.polygon.boundY2-obj.polygon.boundY1))
}

// Set the line position (start and end point).
func (obj *Line[T]) SetPosition(x1, y1, x2, y2 int) {
	if int(obj.x) == x1 && int(obj.y) == y1 && int(obj.x2) == x2 && int(obj.y2) == y2 {
		return
	}
	if !obj.hidden {
		obj.markDirty()
	}
	obj.x = int16(x1)
	obj.y = int16(y1)
	obj.x2 = int16(x2)
	obj.y2 = int16(y2)
	obj.updatePolygon()
	if !obj.hidden {
		obj.markDirty()
	}
}

// Hidden returns whether this object is currently hidden.
func (obj *Line[T]) Hidden() bool {
	return obj.hidden
}

// SetHidden implements gfx.Object. It sets the visibility status of the object
// on screen.
func (obj *Line[T]) SetHidden(hidden bool) {
	if obj.hidden != hidden {
		obj.hidden = hidden
		obj.markDirty()
	}
}
