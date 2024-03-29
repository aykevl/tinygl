package tinygl

import (
	"tinygo.org/x/drivers/pixel"
)

type Object[T pixel.Color] interface {
	// Request an update to this object (that doesn't change the layout).
	RequestUpdate()

	// Request a re-layout of this object.
	RequestLayout()

	// Layout the object in the provided area. The object will take up all the
	// given area (if needed, by filling in the rest with its background color).
	Layout(width, height int)

	// Update the screen if needed. It recurses into children, if the object has
	// any. It does not change any state: this is done separately in
	// MarkUpdated.
	// The displayX, displayY, displayWidth, and displayHeight parameters are
	// the area of the display that needs to be updated. The width and height
	// must never be zero.
	// The x and y parameter are the offset from the start of the object, which
	// can be zero or positive. They can be positive when the parent is a scroll
	// container, for example.
	Update(screen *Screen[T], displayX, displayY, displayWidth, displayHeight, x, y int)

	// MarkUpdated clears any 'needs update' flags that may be present on this
	// object, including its children if any.
	MarkUpdated()

	// Handle some event (usually, touch events).
	HandleEvent(event Event, x, y int)

	// Called when adding a child to a parent. Should only ever be called during
	// the construction of a container object.
	SetParent(object Object[T])

	// Minimal width and height of an object.
	MinSize() (width, height int)

	// Grow factor in horizontal and vertical direction. Space that is left over
	// in a container, is spread over the children according to this factor.
	growable() (horizontal, vertical int)

	// SetGrowable sets the grow factor in the horizontal and vertical
	// direction.
	SetGrowable(horizontal, vertical int)

	// If possible, scroll the parent (or parent-of-parent, recursively) scroll
	// container so that the given area moves into view. The top line is the
	// topmost pixel that should become visible, while the bottom offset is one
	// pixel past the area that should become visible (so that bottom-top is the
	// height of the area to make visible).
	ScrollIntoViewVertical(top, bottom int, object Object[T])

	requestChildUpdate()
	requestChildLayout()
}

type objectFlag uint8

const (
	flagNeedsUpdate      objectFlag = 1 << iota // whether the element needs to have a full redraw
	flagNeedsChildUpdate                        // one of the children needs an update (but not necessarily the object itself)
	flagNeedsLayout                             // this object needs a re-layout
)

type Rect[T pixel.Color] struct {
	parent     Object[T]
	background T
	flags      objectFlag
	grow       uint8 // two 4-bit values (0..15), X is the lower 4 bits and Y the upper 4 bits
}

// MakeRect returns a new initialized Rect object. This is mostly useful to
// initialize an embedded Rect struct in a custom object.
func MakeRect[T pixel.Color](background T) Rect[T] {
	return Rect[T]{
		background: background,
		flags:      flagNeedsUpdate | flagNeedsLayout,
	}
}

func (r *Rect[T]) SetParent(parent Object[T]) {
	if r.parent != nil {
		panic("SetParent: already set")
	}
	r.parent = parent
}

// Parent returns the current parent element, if available.
func (r *Rect[T]) Parent() Object[T] {
	return r.parent
}

// Background returns the current background for this object.
func (r *Rect[T]) Background() T {
	return r.background
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
	r.flags |= flagNeedsLayout
	if r.parent != nil {
		r.parent.requestChildLayout()
	}
}

// NeedsLayout returns whether this object needs to be re-layout, and clears the
// layout flag at the same time.
func (r *Rect[T]) NeedsLayout() (needsLayout bool) {
	needsLayout = r.flags&flagNeedsLayout != 0
	r.flags &^= flagNeedsLayout
	return
}

// NeedsUpdate returns whether this object and/or a child object needs an
// update, and clears the update flag at the same time.
func (r *Rect[T]) NeedsUpdate() (this, child bool) {
	this = r.flags&flagNeedsUpdate != 0
	child = r.flags&flagNeedsChildUpdate != 0
	return this, child
}

// MarkUpdated clears the 'update' flags for this object.
func (r *Rect[T]) MarkUpdated() {
	r.flags &^= (flagNeedsUpdate | flagNeedsChildUpdate)
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

// HandleEvent handles input events like touch taps.
func (r *Rect[T]) HandleEvent(event Event, x, y int) {
	// The base object doesn't handle any events.
	// This is a no-op.
}

func (r *Rect[T]) ScrollIntoViewVertical(top, bottom int, child Object[T]) {
	// The default implementation does nothing.
	// - Regular containers (e.g. VBox) should adjust to the child offset and
	//   forward this call to the parent.
	// - Scroll containers should scroll the element into view.
}
