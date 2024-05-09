package mono

// The mono theme is a plain default style for monochrome displays.

import (
	"github.com/aykevl/tinygl"
	"github.com/aykevl/tinygl/style"
	"tinygo.org/x/drivers/pixel"
	font "tinygo.org/x/tinygl-font"
	"tinygo.org/x/tinygl-font/roboto"
)

type Mono[T pixel.Color] struct {
	style.Theme[T]
}

// NewTheme returns a new basic theme with the given properties.
//
// The default style is black text on a white background.
func NewTheme[T pixel.Color](scale style.Scale, screen *tinygl.Screen[T]) *Mono[T] {
	// Pick a font that is suitable for the given scale.
	// We can't just pick any size, so we have to use some heuristics.
	percent := scale.Percent()
	var font font.Font
	switch {
	case percent <= 100:
		font = roboto.Regular16
	case percent <= 125:
		font = roboto.Regular20
	case percent <= 150:
		font = roboto.Regular24
	case percent <= 175:
		font = roboto.Regular28
	case percent <= 200:
		font = roboto.Regular32
	case percent <= 225:
		font = roboto.Regular36
	case percent <= 250:
		font = roboto.Regular40
	case percent <= 275:
		font = roboto.Regular44
	default: // 300% and larger
		font = roboto.Regular48
	}

	return &Mono[T]{
		Theme: style.Theme[T]{
			Screen:     screen,
			Font:       font,
			Foreground: pixel.NewColor[T](255, 255, 255), // black
			Background: pixel.NewColor[T](0, 0, 0),       // white
			Tint:       pixel.NewColor[T](0, 0, 0),       // white
			Scale:      scale,
		},
	}
}
