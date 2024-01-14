package font

// This file contains some utility functions, copied from elsewhere. They should
// really be merged into the tinygo.org/x/drivers/pixel package.

import "tinygo.org/x/drivers/pixel"

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
