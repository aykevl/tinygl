package pixel

import (
	"image/color"
	"math/bits"
)

// Pixel with a particular color, matching the underlying hardware of a
// particular display. Each pixel is at least 1 byte in size.
// The color format is sRGB (or close to it) in all cases.
type Color interface {
	RGB888 | RGB565BE | RGB555 | RGB444BE

	// The number of bits when stored.
	// This means for example that RGB555 (which is still stored as a 16-bit
	// integer) returns 16.
	BitsPerPixel() int

	RGBA() color.RGBA
}

func NewColor[T Color](r, g, b uint8) T {
	// Ugly cast from color.RGBA to T. The type switch and interface casts are
	// trivially optimized away after instantiation.
	var value T
	switch any(value).(type) {
	case RGB888:
		return any(NewRGB888(r, g, b)).(T)
	case RGB565BE:
		return any(NewRGB565BE(r, g, b)).(T)
	case RGB555:
		return any(NewRGB555(r, g, b)).(T)
	case RGB444BE:
		return any(NewRGB444BE(r, g, b)).(T)
	default:
		panic("unknown color format")
	}
}

// RGB888 format, more commonly used in other places. Mainly here to prove the
// abstractions here work for other formats too.
type RGB888 struct {
	R, G, B uint8
}

func NewRGB888(r, g, b uint8) RGB888 {
	return RGB888{r, g, b}
}

func (c RGB888) BitsPerPixel() int {
	return 24
}

func (c RGB888) RGBA() color.RGBA {
	return color.RGBA{
		R: c.R,
		G: c.G,
		B: c.B,
		A: 255,
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

func (c RGB565BE) BitsPerPixel() int {
	return 16
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

func (c RGB555) BitsPerPixel() int {
	// 15 bits per pixel, but there are 16 bits when stored
	return 16
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

// Color format that is supported by the ST7789 for example.
// It may be a bit faster to use than RGB565BE on very slow SPI buses.
type RGB444BE uint16

func NewRGB444BE(r, g, b uint8) RGB444BE {
	return RGB444BE(r>>4)<<8 | RGB444BE(g>>4)<<4 | RGB444BE(b>>4)
}

func (c RGB444BE) BitsPerPixel() int {
	return 12
}

func (c RGB444BE) RGBA() color.RGBA {
	color := color.RGBA{
		R: uint8(c>>8) << 4,
		G: uint8(c>>4) << 4,
		B: uint8(c>>0) << 4,
		A: 255,
	}
	// Correct color rounding, so that 0xff roundtrips back to 0xff.
	color.R |= color.R >> 4
	color.G |= color.G >> 4
	color.B |= color.B >> 4
	return color
}
