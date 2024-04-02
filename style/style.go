package style

import (
	"github.com/aykevl/tinygl"
	"tinygo.org/x/drivers/pixel"
	font "tinygo.org/x/tinygl-font"
)

type Scale uint8

func NewScale(percent int) Scale {
	return Scale((percent + 12) / 25)
}

func (s Scale) Percent() int {
	return int(s) * 25
}

type Theme[T pixel.Color] struct {
	Screen *tinygl.Screen[T]

	// Theme related properties.
	Font       font.Font
	Foreground T // text, borders, etc
	Background T // background
	Tint       T // highlights (checked boxes, active list element, etc)
	Scale      Scale
}
