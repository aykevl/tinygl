package tinygl

import (
	"github.com/aykevl/tinygl/pixel"
	"github.com/aykevl/tinygl/style"
)

type Object[T pixel.Color] interface {
	// Called when adding a child to a parent. Should only ever be called during
	// the construction of a container object.
	SetParent(object Object[T])

	// Layout the object in the provided area. The object will take up all the
	// given area (if needed, by filling in the rest with its background color).
	Layout(x, y, width, height int)

	// Update the screen if needed. It recurses into children, if the object has
	// any. The object will typically mark itself as having finished updating so
	// another call to Update() won't send any new data to the display.
	Update(screen *Screen[T])

	// Draw the object unconditionally, without changing any flags.
	Draw(x, y int, buf pixel.Image[T])

	// Minimal width and height of an object.
	minSize() (width, height int)
}

type objectFlag uint8

const (
	flagNeedsUpdate objectFlag = iota + 1 // whether the element needs to have a full redraw
)

type Rect[T pixel.Color] struct {
	parent Object[T]

	minWidth      int16 // device independent pixels
	minHeight     int16
	displayX      int16 // "display" pixels are physical pixels
	displayY      int16
	displayWidth  int16
	displayHeight int16
	background    T
	flags         objectFlag
}

func NewRect[T pixel.Color](base style.Style[T], width, height int) *Rect[T] {
	rect := &Rect[T]{}
	rect.init(base, width, height)
	return rect
}

func (r *Rect[T]) init(base style.Style[T], width, height int) {
	r.minWidth = int16(width)
	r.minHeight = int16(height)
	r.background = base.Background
}

func (r *Rect[T]) SetParent(parent Object[T]) {
	if r.parent != nil {
		panic("SetParent: already set")
	}
	r.parent = parent
}

func (r *Rect[T]) minSize() (int, int) {
	return int(r.minWidth), int(r.minHeight)
}

func (r *Rect[T]) Layout(x, y, width, height int) {
	if int(r.displayX) != x || int(r.displayY) != y || int(r.displayWidth) != width || int(r.displayHeight) != height {
		r.displayX = int16(x)
		r.displayY = int16(y)
		r.displayWidth = int16(width)
		r.displayHeight = int16(height)
		r.flags |= flagNeedsUpdate
	}
}

func (r *Rect[T]) Update(screen *Screen[T]) {
	if r.flags&flagNeedsUpdate == 0 || r.displayWidth == 0 || r.displayHeight == 0 {
		return // nothing to do
	}

	paintSolidColor(screen, r.background, int(r.displayX), int(r.displayY), int(r.displayWidth), int(r.displayHeight))

	r.flags &^= flagNeedsUpdate
}

func (r *Rect[T]) Draw(x, y int, img pixel.Image[T]) {
	fillSolidColor(img, r.background)
}

func fillSolidColor[T pixel.Color](img pixel.Image[T], color T) {
	buf := img.Buffer()
	for i := range buf {
		// TODO: this can be optimized a lot.
		// - The store can be done as a 32-bit integer, after checking for
		//   alignment.
		// - Perhaps multiple stores can be done in a row, to reduce the loop
		//   overhead.
		buf[i] = color
	}
}

func paintSolidColor[T pixel.Color](screen *Screen[T], color T, x, y, width, height int) {
	linesPerChunk := len(screen.buffer) / width
	if linesPerChunk > height {
		linesPerChunk = height
	}
	buffer := screen.buffer[:width*linesPerChunk]
	img := pixel.NewImageFromBuffer(buffer, width)
	fillSolidColor(img, color)
	lineStart := 0
	for lineStart < height {
		lines := linesPerChunk
		if lineStart+lines > height {
			lines = height - lineStart
		}
		screen.Send(buffer, x, y+lineStart, width, lines)
		lineStart += linesPerChunk
	}
}
