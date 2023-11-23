package gfx

import (
	"github.com/aykevl/tinygl"
	"tinygo.org/x/drivers/pixel"
)

// CustomCanvas is a canvas widget to draw custom data.
type CustomCanvas[T pixel.Color] struct {
	tinygl.Rect[T]
	minWidth     int16
	minHeight    int16
	canvasWidth  int16
	canvasHeight int16
	update       func(screen *tinygl.Screen[T], displayX, displayY, displayWidth, displayHeight, x, y int)
}

// NewCustomCanvas returns a ready made custom canvas.
// The update callback should always redraw the entire canvas, even if not the
// entire canvas changed.
func NewCustomCanvas[T pixel.Color](background T, minWidth, minHeight int, update func(screen *tinygl.Screen[T], displayX, displayY, displayWidth, displayHeight, x, y int)) *CustomCanvas[T] {
	c := &CustomCanvas[T]{
		Rect:      tinygl.MakeRect(background),
		minWidth:  int16(minWidth),
		minHeight: int16(minHeight),
		update:    update,
	}
	return c
}

// MinSize returns the minimal size as set when the canvas was created.
func (c *CustomCanvas[T]) MinSize() (width, height int) {
	return int(c.minWidth), int(c.minHeight)
}

// Size returns the current canvas size, as updated after the first layout.
func (c *CustomCanvas[T]) Size() (width, height int) {
	return int(c.canvasWidth), int(c.canvasHeight)
}

// Layout implements tinygl.Object.
func (c *CustomCanvas[T]) Layout(width, height int) {
	if int(c.canvasWidth) != width || int(c.canvasHeight) != height {
		c.canvasWidth = int16(width)
		c.canvasHeight = int16(height)
		c.Rect.RequestUpdate()
	}
}

func (c *CustomCanvas[T]) Update(screen *tinygl.Screen[T], displayX, displayY, displayWidth, displayHeight, x, y int) {
	if this, _ := c.NeedsUpdate(); !this { // check the needsUpdate flag
		// No update is needed.
		return
	}
	// needsUpdate flag is cleared

	// Update the contents of the canvas.
	c.update(screen, displayX, displayY, displayWidth, displayHeight, x, y)
}
