package gfx

import (
	"github.com/aykevl/tinygl/image"
	"tinygo.org/x/drivers/pixel"
)

// Image wraps an image.Image object to be drawn on a canvas.
type Image[T pixel.Color] struct {
	baseObject[T]
	hidden bool
	scale  uint8
	img    image.Image[T]
}

// NewImage creates a new image not yet attached to a canvas.
func NewImage[T pixel.Color](img image.Image[T], x, y int) *Image[T] {
	return &Image[T]{
		baseObject: baseObject[T]{
			x: int16(x),
			y: int16(y),
		},
		scale: 1,
		img:   img,
	}
}

// Draw implements the gfx.Object interface.
func (obj *Image[T]) Draw(bufX, bufY int, buf pixel.Image[T]) {
	if obj.hidden {
		return
	}
	obj.img.Draw(buf, bufX-int(obj.x), bufY-int(obj.y), int(obj.scale))
}

// Bounds returns the rectangle coordinates relative to (0, 0) of the canvas.
func (obj *Image[T]) Bounds() (x, y, width, height int) {
	imageWidth, imageHeight := obj.img.Size()
	scale := int(obj.scale)
	return int(obj.x), int(obj.y), imageWidth * scale, imageHeight * scale
}

// Move the rectangle to the new coordinates.
func (obj *Image[T]) Move(x, y int) {
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

// SetImage replaces the image to be drawn on screen.
func (obj *Image[T]) SetImage(img image.Image[T]) {
	if obj.img == img {
		return
	}
	if !obj.hidden {
		obj.markDirty()
	}
	obj.img = img
	if !obj.hidden {
		obj.markDirty()
	}
}

// SetScale changes the default scale from 1 (100%) to a different value.
func (obj *Image[T]) SetScale(scale int) {
	if int(uint8(scale)) != scale {
		panic("Image.SetScale: scale out of range")
	}
	if scale == int(obj.scale) {
		return
	}
	if int(obj.scale) > scale && !obj.hidden {
		obj.markDirty()
	}
	obj.scale = uint8(scale)
	if int(obj.scale) < scale && !obj.hidden {
		obj.markDirty()
	}
}

func (obj *Image[T]) markDirty() {
	imageWidth, imageHeight := obj.img.Size()
	obj.canvas.markDirty(int(obj.x), int(obj.y), imageWidth, imageHeight)
}

// Hidden returns whether this object is currently hidden.
func (obj *Image[T]) Hidden() bool {
	return obj.hidden
}

// SetHidden implements gfx.Object. It sets the visibility status of the
// rectangle on screen.
func (obj *Image[T]) SetHidden(hidden bool) {
	if obj.hidden != hidden {
		obj.hidden = hidden
		obj.markDirty()
	}
}
