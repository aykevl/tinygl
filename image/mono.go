package image

import (
	"unsafe"

	"tinygo.org/x/drivers/pixel"
)

// Monochrome image where each pixel can be either a foreground or a background
// color.
type Mono[T pixel.Color] struct {
	data   unsafe.Pointer // use a raw pointer for more efficiency
	width  int16
	height int16

	// Foreground color, used for the set bits.
	Foreground T

	// Background color, used for the clear bits.
	Background T
}

// Make a new monochrome image with the image data in the 'data' string. The
// string is treated as a byte array, not an actual string, but it's the only
// truly read-only value in Go. The data string must be long enough to contain
// all the bits in the image.
//
// You can create image data this way using ImageMagick, and include them in the
// program using "embed":
//
//	convert image.png -channel RGB -negate -monochrome -depth 1 MONO:image.raw
func MakeMono[T pixel.Color](foreground, background T, width, height int, data string) Mono[T] {
	if len(data) < (width*height+7)/8 {
		panic("ImageMono: data too short")
	}
	return Mono[T]{
		data:       unsafe.Pointer(unsafe.StringData(data)),
		width:      int16(width),
		height:     int16(height),
		Foreground: foreground,
		Background: background,
	}
}

// Size returns the image size of this image.
func (img Mono[T]) Size() (width, height int) {
	return int(img.width), int(img.height)
}

// Draw the image to the pixel buffer at the given coordinates. Anything drawn
// outside the buffer is skipped.
func (img Mono[T]) Draw(buf pixel.Image[T], x, y, scale int) {
	if scale == 0 {
		return
	}
	if scale != 1 {
		// TODO: implement scale other than the default scale.
		panic("todo: scale != 1")
	}
	// TODO: rewrite this all with proper bounds (no useless loops) and scaling
	// in mind.
	bufWidth, bufHeight := buf.Size()
	for drawY := 0; drawY < bufHeight; drawY++ {
		for drawX := 0; drawX < bufWidth; drawX++ {
			imageX := drawX + x
			imageY := drawY + y
			// TODO: don't skip pixels here, instead adjust the for loop
			// start/end.
			if imageX < 0 || imageX >= int(img.width) {
				continue
			}
			if imageY < 0 || imageY >= int(img.height) {
				continue
			}
			bitIndex := imageY*int(img.width) + imageX
			color := img.Background
			if *(*uint8)(unsafe.Add(img.data, bitIndex/8))&(uint8(1)<<(bitIndex%8)) != 0 {
				color = img.Foreground
			}
			buf.Set(drawX, drawY, color)
		}
	}
}
