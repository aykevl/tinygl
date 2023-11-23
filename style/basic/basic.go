package basic

// The basic theme is a plain default style that should work fine in most cases.
// It prioritizes being lightweight and practical over looking nice.

import (
	"github.com/aykevl/tinygl"
	"github.com/aykevl/tinygl/style"
	"tinygo.org/x/drivers/pixel"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freesans"
)

type Basic[T pixel.Color] struct {
	screen *tinygl.Screen[T]

	// Theme related properties.
	Font       *tinyfont.Font
	Foreground T // text, borders, etc
	Background T // background
	Tint       T // highlights (checked boxes, active list element, etc)
	Scale      style.Scale
}

// NewTheme returns a new basic theme with the given properties.
//
// The default style is black text on a white background with a blue tint color.
func NewTheme[T pixel.Color](scale style.Scale, screen *tinygl.Screen[T]) *Basic[T] {
	// Pick a font that is suitable for the given scale.
	// We can't just pick any size, so we have to use some heuristics.
	percent := scale.Percent()
	var font *tinyfont.Font
	switch {
	case percent <= 112: // around 100%
		font = &freesans.Regular9pt7b
	case percent <= 150: // around 133%
		font = &freesans.Regular12pt7b
	case percent <= 225: // around 200%
		font = &freesans.Regular18pt7b
	default: // around 266%
		font = &freesans.Regular24pt7b
	}

	return &Basic[T]{
		screen:     screen,
		Font:       font,
		Foreground: pixel.NewColor[T](0, 0, 0),       // black
		Background: pixel.NewColor[T](255, 255, 255), // white
		Tint:       pixel.NewColor[T](64, 64, 255),   // light blue
		Scale:      scale,
	}
}
