// Package font provides simple fonts for use on embedded devices.
//
// For more information about the format, see the README.
package font

import (
	"tinygo.org/x/drivers/pixel"
)

// Bitmap font with two bits per pixel.
type Font struct {
	data string
}

// Make a new font from a raw data string.
// This will panic if the font is of a different version than is expected.
func Make(data string) Font {
	if data[0] != 0 {
		panic("font: unknown version")
	}
	return Font{data: data}
}

// Size returns the original font size in pixels. This is usually smaller than
// the height.
func (f Font) Size() int {
	return int(f.data[1])
}

// Height returns the recommended amount of vertical space between lines. This
// is usually a bit more than the font size.
func (f Font) Height() int {
	return int(f.data[2])
}

// Ascent returns the distance in pixels from the top of the line height to the
// baseline.
func (f Font) Ascent() int {
	return int(f.data[3])
}

// Calculate the width of a single of text in pixels.
func (f Font) LineWidth(text string) int {
	width := 0
	for _, r := range text {
		glyph, ok := f.glyph(r)
		if !ok {
			continue
		}
		width += glyph.advance()
	}
	return width
}

// Look up the glyph for this rune.
func (f Font) glyph(r rune) (fontGlyph, bool) {
	offset := 4 // header size
	for {
		num := int(f.data[offset])<<0 + int(f.data[offset+1])<<8
		if num == 0 {
			break
		}
		start := rune(f.data[offset+2])<<0 + rune(f.data[offset+3])<<8 + rune(f.data[offset+4])<<16
		offset += 2 + 3
		if r >= start && r < start+rune(num) {
			glyph := fontGlyph{
				font: f,
			}
			glyph.offset = int(f.data[offset+int(r-start)*2+0]) + int(f.data[offset+int(r-start)*2+1])<<8
			return glyph, true
		}
		offset += num * 2
	}
	return fontGlyph{}, false
}

// Single glyph loaded from the font data.
type fontGlyph struct {
	font   Font
	offset int // index into the font data
}

// Return the advance (that is, how much space this character takes up by
// default). This does not include kerning.
func (g fontGlyph) advance() int {
	return int(g.font.data[g.offset+0])
}

// Return the bounding box of the bitmap.
func (g fontGlyph) boundingBox() (top, bottom, left, right int) {
	top = int(int8(g.font.data[g.offset+1]))
	bottom = int(int8(g.font.data[g.offset+2]))
	left = int(int8(g.font.data[g.offset+3]))
	right = int(int8(g.font.data[g.offset+4]))
	return
}

// Draw the given text at the given coordinates.
// The font is drawn from the origin, so if you want to draw a line of size
// font.Height(), you'd use font.Ascent() for the y value.
// Some characters actually stick out slightly to the left, so be aware of that
// (this is one reason to have a bit of margin).
func Draw[T pixel.Color](font Font, text string, x, y int, color T, buf pixel.Image[T]) {
	const bits = 2
	w, h := buf.Size()
	for _, r := range text {
		glyph, ok := font.glyph(r)
		if !ok {
			continue
		}
		top, bottom, left, right := glyph.boundingBox()
		maskWidth := right - left
		maskHeight := bottom - top
		index := (glyph.offset + 5) * 8
		for gy := 0; gy < maskHeight; gy++ {
			py := y + gy + top
			if uint(py) >= uint(h) {
				// Skip this line.
				index += maskWidth * bits
				continue
			}
			for gx := 0; gx < maskWidth; gx++ {
				c := (font.data[index/8] >> (index % 8)) & (1<<bits - 1)
				if c != 0 {
					px := x + gx + left
					if uint(px) < uint(w) {
						const cmax = 1<<bits - 1
						if c == cmax {
							buf.Set(px, py, color)
						} else {
							bottom := buf.Get(px, py)
							blended := naiveBlend(bottom, color, c*(255/cmax))
							buf.Set(px, py, blended)
						}
					}
				}
				index += bits
			}
		}
		x += glyph.advance()
	}
}
