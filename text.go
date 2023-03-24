package tinygl

import (
	"github.com/aykevl/tinygl/pixel"
	"github.com/aykevl/tinygl/style"
	"tinygo.org/x/tinyfont"
)

type Text[T pixel.Color] struct {
	Rect[T]
	text  string
	font  *tinyfont.Font
	color T
}

func NewText[T pixel.Color](base style.Style[T], text string) *Text[T] {
	t := &Text[T]{text: text}
	t.font = base.Font
	t.color = base.Foreground

	// Calculate bounding box for the text.
	_, outerWidth := tinyfont.LineWidth(t.font, text)
	height := t.font.BBox[1]
	t.init(base, int(outerWidth), int(height))
	return t
}

func (t *Text[T]) Update(screen *Screen[T]) {
	if t.flags&flagNeedsUpdate == 0 || t.displayWidth == 0 || t.displayHeight == 0 {
		return // nothing to do
	}

	// Draw text in the center of the provided area.
	textX := int(t.displayX) + (int(t.displayWidth)-int(t.minWidth))/2
	textY := int(t.displayY) + (int(t.displayHeight)-int(t.minHeight))/2 - int(t.font.BBox[3])
	paintText(screen, int(t.displayX), int(t.displayY), int(t.displayWidth), int(t.displayHeight), textX, textY, t.text, t.font, t.background, t.color)

	t.flags &^= flagNeedsUpdate
}

func paintText[T pixel.Color](screen *Screen[T], x, y, width, height, textX, textY int, text string, font *tinyfont.Font, background, foreground T) {
	linesPerChunk := len(screen.buffer) / width
	if linesPerChunk > height {
		linesPerChunk = height
	}
	buffer := screen.buffer[:width*linesPerChunk]
	img := pixel.NewImageFromBuffer(buffer, width)
	lineStart := 0
	// Note: it may be more efficient to draw text left to right rather than
	// downwards, drawing only the glyphs that are part of that area.
	for lineStart < height {
		lines := linesPerChunk
		if lineStart+lines > height {
			lines = height - lineStart
		}
		subimg := img.SubImage(0, 0, width, lines)
		fillSolidColor(subimg, background)
		tinyfont.WriteLine(pixel.DisplayerImage[T]{Image: subimg}, font, int16(textX), int16(textY-lineStart), text, foreground.RGBA())
		screen.Send(subimg.Buffer(), x, y+lineStart, width, lines)
		lineStart += linesPerChunk
	}
}
