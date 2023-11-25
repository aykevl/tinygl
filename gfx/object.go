package gfx

import "tinygo.org/x/drivers/pixel"

// Object on a canvas.
type Object[T pixel.Color] interface {
	// Draw the object on the given image.
	// The X and Y coordinates are the offsets of img from the top left (0, 0)
	// of the canvas.
	Draw(x, y int, img pixel.Image[T])

	// SetHidden changes the visibility of an object.
	SetHidden(bool)

	// Set the canvas property of an object. This can only be done once.
	setCanvas(canvas *Canvas[T])

	// markDirty marks the blocks that this object occupies as dirty so that
	// they will be redrawn on the next update.
	markDirty()
}

// baseObject implements common methods on objects.
type baseObject[T pixel.Color] struct {
	canvas *Canvas[T]
	x      int16
	y      int16
}

func (obj *baseObject[T]) setCanvas(canvas *Canvas[T]) {
	if obj.canvas != nil {
		panic("gfx: object added twice to canvas")
	}
	obj.canvas = canvas
}
