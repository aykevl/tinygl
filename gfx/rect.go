package gfx

import "tinygo.org/x/drivers/pixel"

// Rect is a simple solid color rectangle.
type Rect[T pixel.Color] struct {
	baseObject[T]
	width  int16
	height int16
	color  T
	hidden bool
}

// NewRect returns a new rectangle of the given size to be added to the canvas.
func NewRect[T pixel.Color](color T, x, y, width, height int) *Rect[T] {
	return &Rect[T]{
		baseObject: baseObject[T]{
			x: int16(x),
			y: int16(y),
		},
		width:  int16(width),
		height: int16(height),
		color:  color,
	}
}

// Draw implements the gfx.Object interface.
func (obj *Rect[T]) Draw(imgX, imgY int, img pixel.Image[T]) {
	if obj.hidden {
		return
	}
	drawRect(img, obj.color, int(obj.x)-imgX, int(obj.y)-imgY, int(obj.width), int(obj.height))
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

// SetColor changes the color of the rectangle.
func (obj *Rect[T]) SetColor(color T) {
	if color == obj.color {
		return
	}
	obj.color = color
	if !obj.hidden {
		obj.markDirty()
	}
}

func drawRect[T pixel.Color](img pixel.Image[T], color T, x, y, width, height int) {
	imgWidth, imgHeight := img.Size()
	startX := x
	startY := y
	stopX := x + width
	stopY := y + height
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
			img.Set(x, y, color)
		}
	}
}
