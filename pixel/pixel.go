package pixel

import (
	"image/color"
	"math/bits"
)

// Pixel with a particular color, matching the underlying hardware of a
// particular display. Each pixel is at least 1 byte in size.
// The color format is sRGB (or close to it) in all cases.
type Color interface {
	RGB888 | RGB565BE | RGB555 | LinearGRB888 | RGB444BE

	// The number of bits when stored.
	// This means for example that RGB555 (which is still stored as a 16-bit
	// integer) returns 16.
	BitsPerPixel() int

	RGBA() color.RGBA
}

// NewColor returns the given color based on the RGB values passed in the
// parameters. The input value is assumed to be in sRGB color space.
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
	case LinearGRB888:
		return any(NewLinearGRB888FromSRGB(r, g, b)).(T)
	case RGB444BE:
		return any(NewRGB444BE(r, g, b)).(T)
	default:
		panic("unknown color format")
	}
}

// NewLinearColor returns the given color based on the linear RGB values passed
// in the parameters. Use this if the RGB values are actually linear colors
// (like those that are used in most RGB LEDs) and not when it is in the usual
// sRGB color space (which is not linear).
//
// The input is assumed to be in the linear sRGB color space.
func NewLinearColor[T Color](r, g, b uint8) T {
	var zeroColor T
	switch any(zeroColor).(type) {
	case LinearGRB888:
		return any(NewLinearGRB888(r, g, b)).(T)
	default:
		r = gammaEncodeTable[r]
		g = gammaEncodeTable[g]
		b = gammaEncodeTable[b]
		return NewColor[T](r, g, b)
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

// Color used by the WS2812 and many similar LEDs.
// While most color formats are gamma-encoded sRGB, this is just three LED
// colors (red, green, blue). For ease of use, we'll assume they are in a linear
// sRGB color space (sRGB without gamma encoding).
type LinearGRB888 struct {
	G uint8
	R uint8
	B uint8
}

// Create LinearRGB888 value from linear RGB values.
func NewLinearGRB888(r, g, b uint8) LinearGRB888 {
	return LinearGRB888{
		R: r,
		G: g,
		B: b,
	}
}

// Create LinearRGB888 value from colors in the sRGB color space.
// These colors need to be looked up in a table, so this is not a very efficient
// operation and it looses some color.
func NewLinearGRB888FromSRGB(r, g, b uint8) LinearGRB888 {
	return LinearGRB888{
		R: gammaDecodeTable[r],
		G: gammaDecodeTable[g],
		B: gammaDecodeTable[b],
	}
}

func (c LinearGRB888) BitsPerPixel() int {
	return 24
}

// Return RGBA value in gamma-encoded sRGB color space (that is, the usual color
// format used almost everywhere).
func (c LinearGRB888) RGBA() color.RGBA {
	// We have a linear (light) color RGB, but need to output sRGB.
	// We'll assume the source color is linear sRGB and convert it to regular
	// (non-linear) sRGB.
	return color.RGBA{
		R: gammaEncodeTable[c.R],
		G: gammaEncodeTable[c.G],
		B: gammaEncodeTable[c.B],
		A: 255,
	}
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

// Gamma brightness lookup table:
// https://victornpb.github.io/gamma-table-generator
// gamma = 2.20 steps = 256 range = 0-255
var gammaDecodeTable = [256]uint8{
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2,
	3, 3, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 5, 6, 6, 6,
	6, 7, 7, 7, 8, 8, 8, 9, 9, 9, 10, 10, 11, 11, 11, 12,
	12, 13, 13, 13, 14, 14, 15, 15, 16, 16, 17, 17, 18, 18, 19, 19,
	20, 20, 21, 22, 22, 23, 23, 24, 25, 25, 26, 26, 27, 28, 28, 29,
	30, 30, 31, 32, 33, 33, 34, 35, 35, 36, 37, 38, 39, 39, 40, 41,
	42, 43, 43, 44, 45, 46, 47, 48, 49, 49, 50, 51, 52, 53, 54, 55,
	56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71,
	73, 74, 75, 76, 77, 78, 79, 81, 82, 83, 84, 85, 87, 88, 89, 90,
	91, 93, 94, 95, 97, 98, 99, 100, 102, 103, 105, 106, 107, 109, 110, 111,
	113, 114, 116, 117, 119, 120, 121, 123, 124, 126, 127, 129, 130, 132, 133, 135,
	137, 138, 140, 141, 143, 145, 146, 148, 149, 151, 153, 154, 156, 158, 159, 161,
	163, 165, 166, 168, 170, 172, 173, 175, 177, 179, 181, 182, 184, 186, 188, 190,
	192, 194, 196, 197, 199, 201, 203, 205, 207, 209, 211, 213, 215, 217, 219, 221,
	223, 225, 227, 229, 231, 234, 236, 238, 240, 242, 244, 246, 248, 251, 253, 255,
}

// Gamma brightness lookup table:
// https://victornpb.github.io/gamma-table-generator
// gamma = 0.45 steps = 256 range = 0-255
var gammaEncodeTable = [256]uint8{
	0, 21, 28, 34, 39, 43, 46, 50, 53, 56, 59, 61, 64, 66, 68, 70,
	72, 74, 76, 78, 80, 82, 84, 85, 87, 89, 90, 92, 93, 95, 96, 98,
	99, 101, 102, 103, 105, 106, 107, 109, 110, 111, 112, 114, 115, 116, 117, 118,
	119, 120, 122, 123, 124, 125, 126, 127, 128, 129, 130, 131, 132, 133, 134, 135,
	136, 137, 138, 139, 140, 141, 142, 143, 144, 144, 145, 146, 147, 148, 149, 150,
	151, 151, 152, 153, 154, 155, 156, 156, 157, 158, 159, 160, 160, 161, 162, 163,
	164, 164, 165, 166, 167, 167, 168, 169, 170, 170, 171, 172, 173, 173, 174, 175,
	175, 176, 177, 178, 178, 179, 180, 180, 181, 182, 182, 183, 184, 184, 185, 186,
	186, 187, 188, 188, 189, 190, 190, 191, 192, 192, 193, 194, 194, 195, 195, 196,
	197, 197, 198, 199, 199, 200, 200, 201, 202, 202, 203, 203, 204, 205, 205, 206,
	206, 207, 207, 208, 209, 209, 210, 210, 211, 212, 212, 213, 213, 214, 214, 215,
	215, 216, 217, 217, 218, 218, 219, 219, 220, 220, 221, 221, 222, 223, 223, 224,
	224, 225, 225, 226, 226, 227, 227, 228, 228, 229, 229, 230, 230, 231, 231, 232,
	232, 233, 233, 234, 234, 235, 235, 236, 236, 237, 237, 238, 238, 239, 239, 240,
	240, 241, 241, 242, 242, 243, 243, 244, 244, 245, 245, 246, 246, 247, 247, 248,
	248, 249, 249, 249, 250, 250, 251, 251, 252, 252, 253, 253, 254, 254, 255, 255,
}
