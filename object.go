package tinygl

import (
	"github.com/aykevl/tinygl/pixel"
)

type Object[T pixel.Color] interface {
	// Request an update to this object (that doesn't change the layout).
	RequestUpdate()
	requestChildUpdate()

	// Request a re-layout of this object.
	RequestLayout()
	requestChildLayout()

	// Layout the object in the provided area. The object will take up all the
	// given area (if needed, by filling in the rest with its background color).
	Layout(x, y, width, height int)

	// Update the screen if needed. It recurses into children, if the object has
	// any. The object will typically mark itself as having finished updating so
	// another call to Update() won't send any new data to the display.
	Update(screen *Screen[T])

	// Handle some event (usually, touch events).
	HandleEvent(event Event, x, y int)

	// Called when adding a child to a parent. Should only ever be called during
	// the construction of a container object.
	SetParent(object Object[T])

	// Minimal width and height of an object.
	minSize() (width, height int)

	// Display rect used by this object.
	bounds() (x, y, width, height int)

	// Grow factor in horizontal and vertical direction. Space that is left over
	// in a container, is spread over the children according to this factor.
	growable() (horizontal, vertical int)

	// SetGrowable sets the grow factor in the horizontal and vertical
	// direction.
	SetGrowable(horizontal, vertical int)
}

type objectFlag uint8

const (
	flagNeedsUpdate      objectFlag = 1 << iota // whether the element needs to have a full redraw
	flagNeedsChildUpdate                        // one of the children needs an update (but not necessarily the object itself)
	flagNeedsLayout                             // this object needs a re-layout
	flagNeedsChildLayout                        // one of the children of this object needs a re-layout
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
	grow          uint8 // two 4-bit values (0..15), X is the lower 4 bits and Y the upper 4 bits
}

func NewRect[T pixel.Color](background T, width, height int) *Rect[T] {
	rect := &Rect[T]{}
	rect.init(background, width, height)
	return rect
}

// MakeRect returns a new initialized Rect object. This is mostly useful to
// initialize an embedded Rect struct in a custom object, use NewRect otherwise.
func MakeRect[T pixel.Color](background T, width, height int) Rect[T] {
	rect := Rect[T]{}
	rect.init(background, width, height)
	return rect
}

func (r *Rect[T]) init(background T, width, height int) {
	r.minWidth = int16(width)
	r.minHeight = int16(height)
	r.background = background
}

func (r *Rect[T]) SetParent(parent Object[T]) {
	if r.parent != nil {
		panic("SetParent: already set")
	}
	r.parent = parent
}

// Background returns the current background for this object.
func (r *Rect[T]) Background() T {
	return r.background
}

// Bounds returns the raw (unscaled) display coordinates.
func (r *Rect[T]) Bounds() (int, int, int, int) {
	return int(r.displayX), int(r.displayY), int(r.displayWidth), int(r.displayHeight)
}

func (r *Rect[T]) RequestUpdate() {
	r.flags |= flagNeedsUpdate
	if r.parent != nil {
		r.parent.requestChildUpdate()
	}
}

func (r *Rect[T]) requestChildUpdate() {
	r.flags |= flagNeedsChildUpdate
	if r.parent != nil {
		r.parent.requestChildUpdate()
	}
}

func (r *Rect[T]) RequestLayout() {
	r.flags |= flagNeedsLayout
	if r.parent != nil {
		r.parent.requestChildLayout()
	}
}

func (r *Rect[T]) requestChildLayout() {
	r.flags |= flagNeedsChildLayout
	if r.parent != nil {
		r.parent.requestChildLayout()
	}
}

// NeedsUpdate returns whether this object and/or a child object needs an
// update, and clears the update flag at the same time.
func (r *Rect[T]) NeedsUpdate() (this, child bool) {
	this = r.flags&flagNeedsUpdate != 0
	child = r.flags&flagNeedsChildUpdate != 0
	r.flags &^= (flagNeedsUpdate | flagNeedsChildUpdate)
	return this, child
}

func (r *Rect[T]) minSize() (int, int) {
	return int(r.minWidth), int(r.minHeight)
}

// bounds returns the area on the display that is being used by this rectangle.
func (r *Rect[T]) bounds() (x, y, width, height int) {
	return int(r.displayX), int(r.displayY), int(r.displayWidth), int(r.displayHeight)
}

func (r *Rect[T]) Layout(x, y, width, height int) {
	if int(r.displayX) != x || int(r.displayY) != y || int(r.displayWidth) != width || int(r.displayHeight) != height || r.flags&flagNeedsLayout != 0 {
		r.displayX = int16(x)
		r.displayY = int16(y)
		r.displayWidth = int16(width)
		r.displayHeight = int16(height)
		r.RequestUpdate()
		r.flags &^= flagNeedsLayout
	}
}

// SetGrowable sets the grow factor in the horizontal and vertical direction.
// This means, for example, if one object has a factor of 1 and another of 2,
// the first object will get ⅓ and the second object will get ⅔ of the remaining
// space in the container after the minimal space has been given to all
// children.
func (r *Rect[T]) SetGrowable(horizontal, vertical int) {
	if horizontal&0x0f != horizontal || vertical&0x0f != vertical {
		panic("SetGrowable: out of range")
	}
	oldGrow := r.grow
	r.grow = uint8(horizontal&0x0f) | uint8(vertical&0x0f)<<4
	if r.grow != oldGrow {
		r.RequestLayout()
	}
}

func (r *Rect[T]) growable() (x, y int) {
	return int(r.grow & 0x0f), int(r.grow&0xf0) >> 4
}

func (r *Rect[T]) Update(screen *Screen[T]) {
	if r.flags&flagNeedsUpdate == 0 || r.displayWidth == 0 || r.displayHeight == 0 {
		return // nothing to do
	}

	PaintSolidColor(screen, r.background, int(r.displayX), int(r.displayY), int(r.displayWidth), int(r.displayHeight))

	r.flags &^= flagNeedsUpdate
}

// HandleEvent handles input events like touch taps.
func (r *Rect[T]) HandleEvent(event Event, x, y int) {
	// The base object doesn't handle any events.
	// This is a no-op.
}
