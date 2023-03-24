package style

import (
	"image/color"

	"github.com/aykevl/tinygl/pixel"
	"tinygo.org/x/tinyfont"
)

type Scale uint8

func NewScale(percent int) Scale {
	return Scale(percent / 25)
}

type Style[T pixel.Color] struct {
	Scale      Scale
	Foreground T // text color
	Background T // background color
	Font       *tinyfont.Font
}

func New[T pixel.Color](scale int, foreground, background T, font *tinyfont.Font) Style[T] {
	return Style[T]{
		Scale:      NewScale(scale),
		Foreground: foreground,
		Background: background,
		Font:       font,
	}
}

func (s Style[T]) WithBackground(bg color.RGBA) Style[T] {
	newStyle := s
	newStyle.Background = pixel.NewColor[T](bg.R, bg.G, bg.B)
	return newStyle
}
