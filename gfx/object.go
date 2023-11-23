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
}
