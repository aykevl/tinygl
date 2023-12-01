package gfx

import "tinygo.org/x/drivers/pixel"

// Circle is a simple solid color circle.
type Circle[T pixel.Color] struct {
	baseObject[T]
	radius int16
	color  T
	hidden bool
}

// NewCircle implements a circle object, with antialiased edges.
//
// The (x, y) coordinates are of the center of the circle, right in the middle
// of the innermost 4 pixels. This means for example that if you draw a circle
// at (0, 0) of a large enough canvas, exactly a quarter of the circle is shown.
//
// Warning: this API might change once circles with holes and/or circles with
// borders are added.
func NewCircle[T pixel.Color](color T, x, y, radius int) *Circle[T] {
	return &Circle[T]{
		baseObject: baseObject[T]{
			x: int16(x),
			y: int16(y),
		},
		radius: int16(radius),
		color:  color,
	}
}

// Draw implements the gfx.Object interface.
func (obj *Circle[T]) Draw(imgX, imgY int, img pixel.Image[T]) {
	if obj.hidden {
		return
	}

	// Don't draw circles that are too small to see.
	// (This isn't just for performance: the logic below assumes obj.radius is
	// non-zero).
	if int(obj.radius) <= 0 {
		return
	}

	// If the to-be-drawn area is outside the img area, don't paint.
	// TODO: we might also want to do the same thing in the X direction. But at
	// least for whole-screen updates, it doesn't matter as img will usually be
	// the width of the screen anyway.
	_, imgHeight := img.Size()
	if imgY > int(obj.y)+int(obj.radius) {
		return
	}
	if imgY+imgHeight < int(obj.y)-int(obj.radius) {
		return
	}

	// Draw the entire circle in one go.
	// We do this by treating the circle as two halves split horizontally in the
	// middle. We draw line by line, starting at the middle most lines working
	// towards the top and the bottom at the same time.
	x := int(obj.x) - imgX
	y := int(obj.y) - imgY
	cx := int(obj.radius) - 1
	r2 := int(obj.radius) * int(obj.radius) // r²
	aaCutoff := int(obj.radius) * 2         // why this value??
	intensityDiv := (256 * 255 / aaCutoff)  // do this division only once instead of for every antialiased pixel
	for cy := 1; cy <= int(obj.radius); cy++ {
		var d2 int
		for {
			// Paint a single antialiased pixel.
			// I honestly don't know why this works exactly, but it does result
			// in nice looking antialiased circles. It would be nice to look
			// into using a "real" algorithm at some point which might be faster
			// and/or more correct.
			//
			// I did a quick test and it appears that most of the time (around
			// three quarters) spent here is in drawing/blending the pixels,
			// _not_ in calculating the AA value. So I don't think a different
			// algorithm could significantly speed up AA. That said, perhaps
			// blending itself could be improved a bit.
			aaDistance := (cx+1)*(cx+1) + cy*cy
			extra := r2 - aaDistance + aaCutoff
			if extra > 0 {
				// Calculate the intensity in the following way, but do the
				// division only once per Draw call:
				//    intensity := extra * 255 / offset
				intensity := extra * intensityDiv / 256
				blendPixel(img, x+cx, y+cy-1, obj.color, uint8(intensity))
				blendPixel(img, x+cx, y-cy, obj.color, uint8(intensity))
				blendPixel(img, x-cx-1, y-cy, obj.color, uint8(intensity))
				blendPixel(img, x-cx-1, y+cy-1, obj.color, uint8(intensity))
			}

			// Calculate the radius from the center of the circle:
			//
			//    cx² + cy²
			//
			// I believe this is slightly incorrect: it should actually be the
			// following formula:
			//
			//    (cx-0.5)² + (cy-0.5)²
			//
			// (Remeber that the circle center is at the intersection of four
			// pixels, and we're calculating the distance to a pixel which is
			// best represented at the middle of that pixel.)
			// But even though this might be slightly wrong, it works well
			// enough for us.
			d2 = cx*cx + cy*cy
			if d2 <= r2 {
				// Make sure the pixels we'll be drawing are exactly within the
				// circle: antialiasing pixels on the edge of the circle is done
				// above.
				break
			}
			cx--
		}

		// Fill the inside of the circle using two lines (for the top half and
		// the bottom half of the circle).
		drawLine(img, x-cx, x+cx, y-cy, obj.color)
		drawLine(img, x-cx, x+cx, y+cy-1, obj.color)
	}
}

func (obj *Circle[T]) markDirty() {
	x := int(obj.x) - int(obj.radius)
	y := int(obj.y) - int(obj.radius)
	w := int(obj.radius) * 2
	h := int(obj.radius) * 2
	obj.canvas.markDirty(x, y, w, h)
}

// Hidden returns whether this object is currently hidden.
func (obj *Circle[T]) Hidden() bool {
	return obj.hidden
}

// SetHidden implements gfx.Object. It sets the visibility status of the
// circle on screen.
func (obj *Circle[T]) SetHidden(hidden bool) {
	if obj.hidden != hidden {
		obj.hidden = hidden
		obj.markDirty()
	}
}
