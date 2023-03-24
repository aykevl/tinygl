package pixel

import "image/color"

// Pixel with a particular color, matching the underlying hardware of a
// particular display. Each pixel is at least 1 byte in size.
// The color format is sRGB (or close to it) in all cases.
type Color interface {
	RGB565BE | RGB888
	RGBA() color.RGBA
}

func NewColor[T Color](r, g, b uint8) T {
	// Ugly cast from color.RGBA to T. The type switch and interface casts are
	// trivially optimized away after instantiation.
	var value T
	switch any(value).(type) {
	case RGB565BE:
		return any(NewRGB565BE(r, g, b)).(T)
	case RGB888:
		return any(NewRGB888(r, g, b)).(T)
	default:
		panic("unknown color format")
	}
}

// RGB565 as used in many SPI displays. Stored as a big endian value.
type RGB565BE uint16

func NewRGB565BE(r, g, b uint8) RGB565BE {
	val := RGB565BE(r&0xF8)<<8 +
		RGB565BE(g&0xFC)<<3 +
		RGB565BE(b&0xF8)>>3
	// Swap endianness (make big endian).
	// TODO: this can be done in a single instruction on ARM (rev16), but the
	// compiler doesn't seem to realize this.
	val = val<<8 | val>>8
	return val
}

func (c RGB565BE) RGBA() color.RGBA {
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
