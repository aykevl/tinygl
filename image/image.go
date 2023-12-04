// Package image wraps read-only image buffers and allows drawing them on pixel
// buffers.
package image

import (
	"tinygo.org/x/drivers/pixel"
)

// Image interface that is implemented by all the image types in this package.
type Image[T pixel.Color] interface {
	// Size of the image.
	Size() (width, height int)

	// Draw the image to the given pixel buffer.
	Draw(buf pixel.Image[T], x, y, scale int)
}

// SetPixel updates the pixel at the given coordinates if it lies within the
// buffer. Nothing happens when the pixel lies outside the buffer.
func setPixel[T pixel.Color](buf pixel.Image[T], x, y int, color T) {
	bufWidth, bufHeight := buf.Size()
	if uint(x) >= uint(bufWidth) || uint(y) >= uint(bufHeight) {
		return
	}
	buf.Set(x, y, color)
}
