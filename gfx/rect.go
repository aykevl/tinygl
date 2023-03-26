package gfx

import "github.com/aykevl/tinygl/pixel"

// Rect is a simple solid color rectangle.
type Rect[T pixel.Color] struct {
	canvas *Canvas[T]
	x      int16
	y      int16
	width  int16
	height int16
	color  T
	hidden bool
}

// Draw implements the gfx.Object interface.
func (obj *Rect[T]) Draw(imgX, imgY int, img pixel.Image[T]) {
	if obj.hidden {
		return
	}
	imgWidth, imgHeight := img.Size()
	startX := int(obj.x) - imgX
	startY := int(obj.y) - imgY
	stopX := startX + int(obj.width)
	stopY := startY + int(obj.height)
	if stopX < 0 || stopY < 0 || startX >= imgWidth || startY >= imgHeight {
		return // out of bounds
	}
	if startX < 0 {
		startX = 0
	}
	if startY < 0 {
		startY = 0
	}
	if stopX >= imgWidth {
		stopX = imgWidth
	}
	if stopY >= imgHeight {
		stopY = imgHeight
	}
	for y := startY; y < stopY; y++ {
		for x := startX; x < stopX; x++ {
			img.Set(x, y, obj.color)
		}
	}
}

// Bounds returns the rectangle coordinates relative to (0, 0) of the canvas.
func (obj *Rect[T]) Bounds() (x, y, width, height int) {
	return int(obj.x), int(obj.y), int(obj.width), int(obj.height)
}

// Move the rectangle to the new coordinates.
func (obj *Rect[T]) Move(x, y int) {
	// TODO: be smarter and only invalidate the area that is actually dirty (the
	// rect is a solid color, so if it only moved a little bit most of the
	// screen doesn't need to update).
	if x == int(obj.x) && y == int(obj.y) {
		return
	}
	if !obj.hidden {
		obj.markDirty()
	}
	obj.x = int16(x)
	obj.y = int16(y)
	if !obj.hidden {
		obj.markDirty()
	}
}

func (obj *Rect[T]) markDirty() {
	obj.canvas.markDirty(int(obj.x), int(obj.y), int(obj.width), int(obj.height))
}

// Hidden returns whether this object is currently hidden.
func (obj *Rect[T]) Hidden() bool {
	return obj.hidden
}

// SetHidden implements gfx.Object. It sets the visibility status of the
// rectangle on screen.
func (obj *Rect[T]) SetHidden(hidden bool) {
	if obj.hidden != hidden {
		obj.hidden = hidden
		obj.markDirty()
	}
}
