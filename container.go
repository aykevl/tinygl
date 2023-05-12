package tinygl

import (
	"github.com/aykevl/tinygl/pixel"
)

type VBox[T pixel.Color] struct {
	Rect[T]
	children     []Object[T]
	childHeights []int16
}

func NewVBox[T pixel.Color](background T, children ...Object[T]) *VBox[T] {
	box := &VBox[T]{}
	box.children = children
	box.childHeights = make([]int16, len(children))

	for _, child := range children {
		child.SetParent(&box.Rect)
	}

	box.Rect.background = background
	box.flags |= flagNeedsLayout | flagNeedsChildLayout | flagNeedsUpdate | flagNeedsChildUpdate
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

func (b *VBox[T]) Layout(width, height int) {
	if b.flags&(flagNeedsLayout|flagNeedsChildLayout) == 0 {
		return // nothing to do
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

	// Go through each child and determine its size.
	// The 'leftover' parts are the extra pixels at the bottom of the VBox.
	leftoverHeight := height - minHeightSum
	leftoverGrowable := growableSum
	if leftoverHeight < 0 {
		leftoverHeight = 0 // don't shrink children when the VBox is full
	}
	childY := 0 // Y relative to the VBox start
	previousChildY := 0
	for i, child := range b.children {
		// Determine child size.
		_, childGrowable := child.growable()
		_, childHeight := child.MinSize()
		if childGrowable != 0 {
			growPixels := leftoverHeight * childGrowable / leftoverGrowable
			childHeight += growPixels
			leftoverHeight -= growPixels
			leftoverGrowable -= childGrowable
		}
		child.Layout(width, childHeight)

		// Check whether the child changed position or size.
		if childHeight != int(b.childHeights[i]) || childY != previousChildY {
			child.RequestUpdate()
		}

		// Keep track of the child height.
		previousChildY += int(b.childHeights[i])
		childY += childHeight
		b.childHeights[i] = int16(childHeight)
	}
	if childY != previousChildY && leftoverHeight != 0 {
		// Redraw the area at the bottom of the VBox (the "slack").
		b.Rect.RequestUpdate()
	}
	b.flags &^= flagNeedsLayout | flagNeedsChildLayout
}

func (b *VBox[T]) Update(screen *Screen[T], displayX, displayY, displayWidth, displayHeight, x, y int) {
	if b.flags&(flagNeedsUpdate|flagNeedsChildUpdate) == 0 {
		return
	}

	childOffset := 0
	for i, child := range b.children {
		childHeight := int(b.childHeights[i])
		childDisplayY := displayY - y + childOffset
		childY := 0
		if childDisplayY < displayY {
			childDisplayY = displayY
			childY = displayY - childDisplayY
		}
		childDisplayHeight := int(b.childHeights[i])
		if childDisplayY+childDisplayHeight > displayY+displayHeight {
			childDisplayHeight = (displayY + displayHeight) - childDisplayY
		}
		if childDisplayHeight > 0 {
			child.Update(screen, displayX, childDisplayY, displayWidth, childDisplayHeight, x, childY)
		}
		childOffset += childHeight
	}

	slackTop := displayY - y + childOffset
	slackBottom := displayY + displayHeight
	slackHeight := slackBottom - slackTop
	if b.flags&flagNeedsUpdate != 0 && slackHeight > 0 {
		PaintSolidColor(screen, b.background, displayX, slackTop, displayWidth, slackHeight)
	}

	b.flags &^= flagNeedsUpdate | flagNeedsChildUpdate // updated, so no need to redraw next time
}

// HandleEvent propagates an event to its children.
func (b *VBox[T]) HandleEvent(event Event, x, y int) {
	childY := 0
	for i, child := range b.children {
		childHeight := b.childHeights[i]
		if y >= childY && y < childY+int(childHeight) {
			child.HandleEvent(event, x, y-childY)
		}
		childY += int(childHeight)
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
