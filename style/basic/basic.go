package basic

// The basic theme is a plain default style that should work fine in most cases.
// It prioritizes being lightweight and practical over looking nice.

import (
	"github.com/aykevl/tinygl"
	"github.com/aykevl/tinygl/style"
	"tinygo.org/x/drivers/pixel"
	font "tinygo.org/x/tinygl-font"
	"tinygo.org/x/tinygl-font/roboto"
)

type Theme[T pixel.Color] struct {
	style.Theme[T]
}

// New returns a new Basic Theme with the given properties.
//
// The default style is black text on a white background with a blue tint color.
func New[T pixel.Color](scale style.Scale, screen *tinygl.Screen[T]) *Theme[T] {
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

	return &Theme[T]{
		Theme: style.Theme[T]{
			Screen:     screen,
			Font:       font,
			Foreground: pixel.NewColor[T](0, 0, 0),       // black
			Background: pixel.NewColor[T](255, 255, 255), // white
			Tint:       pixel.NewColor[T](64, 64, 255),   // light blue
			Scale:      scale,
		},
	}
}
