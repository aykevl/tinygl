package pixel

import (
	"image/color"
	"math/bits"
)

// Pixel with a particular color, matching the underlying hardware of a
// particular display. Each pixel is at least 1 byte in size.
// The color format is sRGB (or close to it) in all cases.
type Color interface {
	RGB565BE | RGB555 | RGB888
	RGBA() color.RGBA
}

func NewColor[T Color](r, g, b uint8) T {
	// Ugly cast from color.RGBA to T. The type switch and interface casts are
	// trivially optimized away after instantiation.
	var value T
	switch any(value).(type) {
	case RGB565BE:
		return any(NewRGB565BE(r, g, b)).(T)
	case RGB555:
		return any(NewRGB555(r, g, b)).(T)
	case RGB888:
		return any(NewRGB888(r, g, b)).(T)
	default:
		panic("unknown color format")
	}
}

// RGB565 as used in many SPI displays. Stored as a big endian value.
type RGB565BE uint16

func NewRGB565BE(r, g, b uint8) RGB565BE {
	val := uint16(r&0xF8)<<8 +
		uint16(g&0xFC)<<3 +
		uint16(b&0xF8)>>3
	// Swap endianness (make big endian).
	// This is done using a single instruction on ARM (rev16).
	val = bits.ReverseBytes16(val)
	return RGB565BE(val)
}

func (c RGB565BE) RGBA() color.RGBA {
	// Note: on ARM, the compiler uses a rev instruction instead of a rev16
	// instruction. I wonder whether this can be optimized further to use rev16
	// instead?
	c = c<<8 | c>>8
	color := color.RGBA{
		R: uint8(c>>11) << 3,
		G: uint8(c>>5) << 2,
		B: uint8(c) << 3,
		A: 255,
	}
	// Correct color rounding, so that 0xff roundtrips back to 0xff.
	color.R |= color.R >> 5
	color.G |= color.G >> 6
	color.B |= color.B >> 5
	return color
}

// Color format used on the GameBoy Advance among others.
type RGB555 uint16

func NewRGB555(r, g, b uint8) RGB555 {
	return RGB555(r)>>3 | (RGB555(g)>>3)<<5 | (RGB555(b)>>3)<<10
}

func (c RGB555) RGBA() color.RGBA {
	color := color.RGBA{
		R: uint8(c>>10) << 3,
		G: uint8(c>>5) << 3,
		B: uint8(c) << 3,
		A: 255,
	}
	// Correct color rounding, so that 0xff roundtrips back to 0xff.
	color.R |= color.R >> 5
	color.G |= color.G >> 5
	color.B |= color.B >> 5
	return color
}

// RGB888 format, more commonly used in other places. Mainly here to prove the
// abstractions here work for other formats too.
type RGB888 struct {
	R, G, B uint8
}

func NewRGB888(r, g, b uint8) RGB888 {
	return RGB888{r, g, b}
}

func (c RGB888) RGBA() color.RGBA {
	return color.RGBA{
		R: c.R,
		G: c.G,
		B: c.B,
		A: 255,
	}
}
