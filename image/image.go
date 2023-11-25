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
