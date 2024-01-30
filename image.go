package tinygl

import (
	"github.com/aykevl/tinygl/image"
	"tinygo.org/x/drivers/pixel"
)

type Image[T pixel.Color] struct {
	Rect[T]
	img    image.Image[T]
	imageX int16
	imageY int16
}

// NewImage creates a new image object that will display the given image. By
// default, the image is centered in the available space, with a black
// background.
func NewImage[T pixel.Color](img image.Image[T]) *Image[T] {
	return &Image[T]{
		img: img,
	}
}

// MinSize returns the minimal size of the image.
func (obj *Image[T]) MinSize() (width, height int) {
	return obj.img.Size()
}

// SetBackground sets the background color for the image.
func (obj *Image[T]) SetBackground(color T) {
	obj.background = color
	obj.RequestUpdate()
}

// Layout implements tinygl.Object.
func (obj *Image[T]) Layout(width, height int) {
	imageWidth, imageHeight := obj.MinSize()
	obj.imageX = int16(width/2 - imageWidth/2)
	obj.imageY = int16(height/2 - imageHeight/2)
	obj.flags &^= flagNeedsLayout
}

// Update implements tinygl.Object.
func (obj *Image[T]) Update(screen *Screen[T], displayX, displayY, displayWidth, displayHeight, x, y int) {
	if obj.flags&flagNeedsUpdate == 0 {
		return // nothing to do
	}

	linesPerChunk := screen.buffer.Len() / displayWidth
	lineStart := 0
	for lineStart < displayHeight {
		lines := linesPerChunk
		if lineStart+lines > displayHeight {
			lines = displayHeight - lineStart
		}
		subimg := screen.buffer.Rescale(displayWidth, lines)
		subimg.FillSolidColor(obj.background) // always necessary because of transparency
		obj.img.Draw(subimg, x-int(obj.imageX), y+lineStart-int(obj.imageY), 1)
		screen.Send(x, y+lineStart, subimg)
		lineStart += lines
	}
}
