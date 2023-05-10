package tinygl

import (
	"github.com/aykevl/tinygl/pixel"
)

type VBox[T pixel.Color] struct {
	Rect[T]
	children []Object[T]
	slack    int16
}

func NewVBox[T pixel.Color](background T, children ...Object[T]) *VBox[T] {
	box := &VBox[T]{}
	box.children = children

	for _, child := range children {
		child.SetParent(box)
	}

	box.Rect.background = background
	box.flags |= flagNeedsUpdate
	return box
}

// RequestUpdate will request an update for this object and all of its children.
func (b *VBox[T]) RequestUpdate() {
	b.Rect.RequestUpdate()
	for _, child := range b.children {
		child.RequestUpdate()
	}
}

// MinSize returns the minimal size to fit around all the elements in this
// container.
func (b *VBox[T]) MinSize() (width, height int) {
	for _, child := range b.children {
		childWidth, childHeight := child.MinSize()
		if childWidth > width {
			width = childWidth
		}
		height += childHeight
	}
	return
}

func (b *VBox[T]) Layout(x, y, width, height int) {
	hasRectChange := x != int(b.displayX) || y != int(b.displayY) || width != int(b.displayWidth) || height != int(b.displayHeight)
	if !hasRectChange && b.flags&(flagNeedsLayout|flagNeedsChildLayout) == 0 {
		return // nothing to do
	}
	if hasRectChange {
		b.displayX = int16(x)
		b.displayY = int16(y)
		b.displayWidth = int16(width)
		b.displayHeight = int16(height)
	}

	// Determine minimal size of all children.
	var minHeightSum int
	var growableSum int
	for _, child := range b.children {
		_, childGrowable := child.growable()
		_, childHeight := child.MinSize()
		minHeightSum += childHeight
		growableSum += childGrowable
	}

	// Go through each child and determine its position and size.
	// The 'leftover' parts are the extra pixels at the bottom of the VBox.
	leftoverHeight := height - minHeightSum
	leftoverGrowable := growableSum
	if leftoverHeight < 0 {
		leftoverHeight = 0 // don't shrink children when the VBox is full
	}
	childY := y
	for _, child := range b.children {
		_, childGrowable := child.growable()
		childHeight := 0
		if childY < y+height {
			_, childHeight = child.MinSize()
			if childGrowable != 0 {
				growPixels := leftoverHeight * childGrowable / leftoverGrowable
				childHeight += growPixels
				leftoverHeight -= growPixels
				leftoverGrowable -= childGrowable
			}
		}
		child.Layout(x, childY, width, childHeight)
		childY += childHeight
	}
	b.slack = int16(leftoverHeight)
	if leftoverHeight > 0 {
		// More of the extra space at the end is exposed, so redraw that area.
		// TODO: only redraw newly exposed area.
		b.flags |= flagNeedsUpdate
	}
	b.flags &^= flagNeedsLayout | flagNeedsChildLayout
}

func (b *VBox[T]) Update(screen *Screen[T]) {
	if b.flags&(flagNeedsUpdate|flagNeedsChildUpdate) == 0 {
		return
	}

	// TODO: combine multiple children in a single buffer if possible
	for _, child := range b.children {
		child.Update(screen)
	}

	if b.flags&flagNeedsUpdate != 0 && b.slack != 0 {
		PaintSolidColor(screen, b.background, int(b.displayX), int(b.displayY)+int(b.displayHeight)-int(b.slack), int(b.displayWidth), int(b.slack))
	}

	b.flags &^= flagNeedsUpdate | flagNeedsChildUpdate // updated, so no need to redraw next time
}

// HandleEvent propagates an event to its children.
func (b *VBox[T]) HandleEvent(event Event, x, y int) {
	for _, child := range b.children {
		childX, childY, childWidth, childHeight := child.Bounds()
		if x < childX || y < childY || x >= childX+childWidth || y >= childY+childHeight {
			continue
		}
		child.HandleEvent(event, x, y)
	}
}

// EventBox wraps an object and handles events for it.
type EventBox[T pixel.Color] struct {
	Object[T]
	handleEvent func(event Event, x, y int)
}

// Create a new wrapper container that handles events.
func NewEventBox[T pixel.Color](child Object[T]) *EventBox[T] {
	return &EventBox[T]{
		Object: child,
	}
}

// SetEventHandler sets the event callback for this object.
func (b *EventBox[T]) SetEventHandler(handler func(event Event, x, y int)) {
	b.handleEvent = handler
}

func (b *EventBox[T]) HandleEvent(event Event, x, y int) {
	if b.handleEvent != nil {
		b.handleEvent(event, x, y)
	}
}
