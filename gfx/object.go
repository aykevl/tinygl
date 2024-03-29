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

// True if the target architecture has pointers that are 16 bits or smaller.
// This is mainly used to detect AVR.
const is16bit = ^uintptr(0)>>16 == 0

// Naive linear blend of two pixel values.
// Naive, because blending assumes pixels are linear while they aren't (they use
// the usual gamma encoding of sRGB). It's good enough for our purposes though,
// and doing a correct blend would be more computationally expensive.
func naiveBlend[T pixel.Color](bottom, top T, alpha uint8) T {
	bottomColor := bottom.RGBA()
	topColor := top.RGBA()
	r := linearBlend(bottomColor.R, topColor.R, alpha)
	g := linearBlend(bottomColor.G, topColor.G, alpha)
	b := linearBlend(bottomColor.B, topColor.B, alpha)
	return pixel.NewColor[T](r, g, b)
}

// Blend the top value into the bottom value, with the given alpha value.
func linearBlend(bottom, top, topAlpha uint8) uint8 {
	if is16bit {
		// Version optimized for AVR.
		bottomPart := uint16(bottom) * uint16(255-topAlpha)
		topPart := uint16(top) * uint16(topAlpha)
		return uint8((bottomPart + topPart + 255) / 256)
	}
	// Version optimized for 32-bit and higher.
	bottomPart := int(bottom) * (255 - int(topAlpha))
	topPart := int(top) * int(topAlpha)
	return uint8((bottomPart + topPart + 255) / 256)
}

// Draw a horizontal straight line without antialiasing.
func drawLine[T pixel.Color](img pixel.Image[T], x0, x1, y int, color T) {
	imgWidth, imgHeight := img.Size()
	if y < 0 || y >= imgHeight {
		return
	}
	if x0 < 0 {
		x0 = 0
	}
	if x1 >= imgWidth {
		x1 = imgWidth
	}
	if x0 >= x1 {
		return
	}
	// TODO: this can be optimized a lot when implemented directly inside
	// pixel.Image. Especially with RGB444.
	for x := x0; x < x1; x++ {
		img.Set(x, y, color)
	}
}

// Draw our color to a given pixel, given a particular alpha channel.
// This is used for antialiasing.
func blendPixel[T pixel.Color](img pixel.Image[T], x, y int, color T, alpha uint8) {
	w, h := img.Size()
	if uint(x) < uint(w) && uint(y) < uint(h) {
		color := naiveBlend(img.Get(x, y), color, alpha)
		img.Set(x, y, color)
	}
}
