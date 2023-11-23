package tinygl

import (
	"image/color"

	"tinygo.org/x/drivers/pixel"
	"tinygo.org/x/tinyfont"
)

type TextAlign uint8

const (
	AlignCenter TextAlign = iota
	AlignLeft
)

type Text[T pixel.Color] struct {
	Rect[T]
	text      string
	font      *tinyfont.Font
	textX     int16
	textY     int16
	textWidth int16
	extra     uint16 // alignment, padding
	color     T
}

func NewText[T pixel.Color](font *tinyfont.Font, foreground, background T, text string) *Text[T] {
	t := MakeText(font, foreground, background, text)
	return &t
}

// MakeText returns a new initialized Rect object. This is mostly useful to
// initialize an embedded Text struct in a custom object. If you want a
// standalone text object, use NewText.
func MakeText[T pixel.Color](font *tinyfont.Font, foreground, background T, text string) Text[T] {
	t := Text[T]{
		text: text,
		font: font,
		Rect: MakeRect(background),
	}
	t.color = foreground

	// Calculate bounding box for the text.
	_, outerWidth := tinyfont.LineWidth(font, text)
	t.textWidth = int16(outerWidth)

	return t
}

// MinSize returns the minimal size of the text label.
func (t *Text[T]) MinSize() (width, height int) {
	padHorizontal, padVertical := t.padding()
	width = int(t.textWidth) + padHorizontal*2
	height = int(t.font.BBox[1]) + padVertical*2
	return
}

// SetText changes the text for this text label.
func (t *Text[T]) SetText(text string) {
	if t.text != text {
		t.text = text
		_, outerWidth := tinyfont.LineWidth(t.font, text)
		if uint32(t.textWidth) != outerWidth {
			t.textWidth = int16(outerWidth)
			t.RequestLayout()
		}
		t.RequestUpdate()
	}
}

func (t *Text[T]) SetAlign(align TextAlign) {
	t.extra = (t.extra &^ 0x0003) | uint16(align)
	t.RequestUpdate()
}

func (t *Text[T]) align() TextAlign {
	return TextAlign(t.extra & 0x0003)
}

// Set horizontal and vertical padding in screen pixels. The padding must be a
// positive integer that is less than 128.
func (t *Text[T]) SetPadding(horizontal, vertical int) {
	t.extra = (t.extra & 0x0003) | uint16(horizontal&0x7f)<<2 | uint16(vertical&0x7f)<<9

	t.RequestLayout()
	t.RequestUpdate()
}

func (t *Text[T]) padding() (horizontal, vertical int) {
	horizontal = (int(t.extra) >> 2) & 0x7f
	vertical = int(t.extra) >> 9
	return
}

// SetBackground changes the background color of the text.
func (t *Text[T]) SetBackground(background T) {
	t.background = background
	t.RequestUpdate()
}

// SetColor sets the text color.
func (t *Text[T]) SetColor(color T) {
	t.color = color
	t.RequestUpdate()
}

func (t *Text[T]) Layout(width, height int) {
	switch t.align() {
	case AlignLeft:
		padHorizontal, _ := t.padding()
		t.textX = int16(padHorizontal)
	default: // AlignCenter
		t.textX = int16((width - int(t.textWidth)) / 2)
	}
	t.textY = int16((height-int(t.font.BBox[1]))/2 - int(t.font.BBox[3]))
	t.flags &^= flagNeedsLayout
}

func (t *Text[T]) Update(screen *Screen[T], displayX, displayY, displayWidth, displayHeight, x, y int) {
	if t.flags&flagNeedsUpdate == 0 {
		return // nothing to do
	}

	// Draw text in the center of the provided area.
	paintText(screen, displayX, displayY, displayWidth, displayHeight, displayX+int(t.textX)-x, displayY+int(t.textY)-y, t.text, t.font, t.background, t.color)

	t.flags &^= flagNeedsUpdate
}

func paintText[T pixel.Color](screen *Screen[T], x, y, width, height, textX, textY int, text string, font *tinyfont.Font, background, foreground T) {
	linesPerChunk := screen.buffer.Len() / width
	if linesPerChunk > height {
		linesPerChunk = height
	}
	lineStart := 0
	// Note: it may be more efficient to draw text left to right rather than
	// downwards, drawing only the glyphs that are part of that area.
	for lineStart < height {
		lines := linesPerChunk
		if lineStart+lines > height {
			lines = height - lineStart
		}
		subimg := screen.buffer.Rescale(width, lines)
		subimg.FillSolidColor(background)
		tinyfont.WriteLine(displayerImage[T]{Image: subimg}, font, int16(textX-x), int16(textY-y-lineStart), text, foreground.RGBA())
		screen.Send(x, y+lineStart, subimg)
		lineStart += linesPerChunk
	}
}

// Wrapper for Image that implements the drivers.Displayer interface.
type displayerImage[T pixel.Color] struct {
	pixel.Image[T]
}

// SetPixel implements the Displayer interface.
func (img displayerImage[T]) SetPixel(x, y int16, color color.RGBA) {
	if x < 0 || y < 0 {
		return
	}
	width, height := img.Image.Size()
	if int(x) >= width || int(y) >= height {
		return
	}
	img.Set(int(x), int(y), pixel.NewColor[T](color.R, color.G, color.B))
}

// Size implements the Displayer interface.
func (img displayerImage[T]) Size() (int16, int16) {
	width, height := img.Image.Size()
	return int16(width), int16(height)
}

// Display implements the Displayer interface. It is a no-op.
func (img displayerImage[T]) Display() error {
	return nil
}
