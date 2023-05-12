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
	box.flags |= flagNeedsLayout | flagNeedsUpdate | flagNeedsChildUpdate
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
	if b.flags&flagNeedsLayout == 0 {
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
	b.flags &^= flagNeedsLayout
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

// ScrollBox is a scrollable wrapper of an object.
// It is growable and reports a minimal size of (0, 0) to the parent, so that it
// will take up the remaining space in the parent box.
type ScrollBox[T pixel.Color] struct {
	child      Object[T]
	offsetX    int
	offsetY    int
	maxOffsetX int
	maxOffsetY int
	lastTouchX int16
	lastTouchY int16
}

// NewScrollBox returns an initialized scroll box with the given child.
// If you want to scroll multiple children (vertically, for example), you can
// use a VBox instead.
func NewScrollBox[T pixel.Color](child Object[T]) *ScrollBox[T] {
	return &ScrollBox[T]{
		child: child,
	}
}

func (b *ScrollBox[T]) MinSize() (width, height int) {
	// A scroll area should be set to expand.
	return 0, 0
}

func (b *ScrollBox[T]) growable() (horizontal, vertical int) {
	// A scroll box is always growable.
	return 1, 1
}

func (b *ScrollBox[T]) SetGrowable(horizontal, vertical int) {
	// No-op, because the scroll area is always growable.
}

func (b *ScrollBox[T]) Layout(width, height int) {
	// Allow the child to use its entire MinSize if it wants to (stretching it
	// to the parent object if needed).
	minWidth, minHeight := b.child.MinSize()
	if width < minWidth {
		b.maxOffsetX = minWidth - width
		width = minWidth
	} else {
		b.maxOffsetX = 0
	}
	if height < minHeight {
		b.maxOffsetY = minHeight - height
		height = minHeight
	} else {
		b.maxOffsetY = 0
	}
	b.child.Layout(width, height)
}

func (b *ScrollBox[T]) Update(screen *Screen[T], displayX, displayY, displayWidth, displayHeight, x, y int) {
	// Redraw the child (if needed) with an offset.
	b.child.Update(screen, displayX, displayY, displayWidth, displayHeight, x+b.offsetX, y+b.offsetY)
}

func (b *ScrollBox[T]) HandleEvent(event Event, x, y int) {
	switch event {
	case TouchStart:
		b.lastTouchX = int16(x)
		b.lastTouchY = int16(y)
	case TouchMove:
		// Add the last distance moved to the offset.
		offsetX := b.offsetX + (int(b.lastTouchX) - x)
		if offsetX < 0 {
			offsetX = 0
		}
		if offsetX > b.maxOffsetX {
			offsetX = b.maxOffsetX
		}
		offsetY := b.offsetY + (int(b.lastTouchY) - y)
		if offsetY < 0 {
			offsetY = 0
		}
		if offsetY > b.maxOffsetY {
			offsetY = b.maxOffsetY
		}
		if offsetX != b.offsetX || offsetY != b.offsetY {
			// The offset changed, so redraw the child.
			// One example where this does not happen, is when scrolling further
			// than is possible in the scroll area. In that case, there is
			// nothing to update as the child position didn't change even if the
			// touch point changed.
			// TODO: use hardware scrolling if available.
			b.offsetX = offsetX
			b.offsetY = offsetY
			b.child.RequestUpdate()
		}
		b.lastTouchX = int16(x)
		b.lastTouchY = int16(y)
	case TouchTap:
		// Pass tap events to the child (with the scroll offset).
		b.child.HandleEvent(event, x+b.offsetX, y+b.offsetY)
	}
}

func (b *ScrollBox[T]) RequestLayout() {
	// Dummy forwarding call.
	b.child.RequestLayout()
}

func (b *ScrollBox[T]) RequestUpdate() {
	// Dummy forwarding call.
	b.child.RequestUpdate()
}

func (b *ScrollBox[T]) SetParent(parent *Rect[T]) {
	// Dummy forwarding call.
	b.child.SetParent(parent)
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
