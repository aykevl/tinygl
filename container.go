package tinygl

import (
	"tinygo.org/x/drivers/pixel"
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
		child.SetParent(box)
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
		childDisplayHeight := int(b.childHeights[i])
		// Cut off the top part if it is outside the update area.
		if childDisplayY < displayY {
			childDisplayHeight -= displayY - childDisplayY
			childY = displayY - childDisplayY
			childDisplayY = displayY
		}
		// Cut off the bottom part if it is outside the update area.
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
	Rect[T]
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
	box := &ScrollBox[T]{
		child: child,
	}
	child.SetParent(box)
	return box
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
			b.RequestUpdate()
		}
		b.lastTouchX = int16(x)
		b.lastTouchY = int16(y)
	case TouchTap:
		// Pass tap events to the child (with the scroll offset).
		b.child.HandleEvent(event, x+b.offsetX, y+b.offsetY)
	}
}

func (b *ScrollBox[T]) RequestUpdate() {
	b.child.RequestUpdate()
}

// VerticalScrollBox is a scrollable wrapper of an object that will use hardware
// acceleration when available.
// It is growable and reports a minimal size of (0, 0) to the parent, so that it
// will take up the remaining space in the parent box.
type VerticalScrollBox[T pixel.Color] struct {
	Rect[T]
	child            Object[T]
	top              Object[T] // top fixed object, or nil
	bottom           Object[T] // bottom fixed object, or nil
	scrollOffset     int       // current distance from the top
	lastScrollOffset int       // scroll offset in previous Update call
	maxScrollOffset  int       // max value for scrollOffset
	lastTouchY       int16
	childHeight      int16
	bottomHeight     int16
	topHeight        int16
}

// NewVerticalScrollBox returns an initialized scroll box with the given child.
// If you want to scroll multiple children (vertically, for example), you can
// use a VBox to wrap these children.
func NewVerticalScrollBox[T pixel.Color](top, child, bottom Object[T]) *VerticalScrollBox[T] {
	b := &VerticalScrollBox[T]{
		top:    top,
		child:  child,
		bottom: bottom,
	}
	if top != nil {
		top.SetParent(b)
	}
	if bottom != nil {
		bottom.SetParent(b)
	}
	child.SetParent(b)
	return b
}

func (b *VerticalScrollBox[T]) MinSize() (width, height int) {
	// The scroll area takes up whatever space the parent gives it.
	var topH, bottomH int
	if b.top != nil {
		_, topH = b.top.MinSize()
	}
	if b.bottom != nil {
		_, bottomH = b.bottom.MinSize()
	}
	return 0, topH + bottomH
}

func (b *VerticalScrollBox[T]) growable() (horizontal, vertical int) {
	// A scroll box is always growable.
	return 1, 1
}

func (b *VerticalScrollBox[T]) SetGrowable(horizontal, vertical int) {
	// No-op, because the scroll area is always growable.
}

func (b *VerticalScrollBox[T]) Layout(width, height int) {
	// Layout the top area.
	if b.top != nil {
		_, topH := b.top.MinSize()
		if topH > height {
			topH = height
		}
		if topH != int(b.topHeight) {
			b.topHeight = int16(topH)
			b.top.Layout(width, topH)
			b.top.RequestUpdate()
			b.child.RequestUpdate()
		}
		height -= topH
	}

	// Layout the bottom area.
	if b.bottom != nil {
		_, bottomH := b.bottom.MinSize()
		if bottomH > height {
			bottomH = height
		}
		if bottomH != int(b.bottomHeight) {
			b.bottomHeight = int16(bottomH)
			b.bottom.Layout(width, bottomH)
			b.bottom.RequestUpdate()
			b.child.RequestUpdate()
		}
		height -= bottomH
	}

	// Allow the child to use its entire MinSize if it wants to (stretching it
	// to the parent object if needed).
	_, childH := b.child.MinSize()
	if childH > height {
		b.maxScrollOffset = childH - height
		childH = height
	} else {
		b.maxScrollOffset = 0
	}
	if childH != int(b.childHeight) {
		b.childHeight = int16(childH)
		b.child.Layout(width, height)
		b.child.RequestUpdate()
	}
}

func (b *VerticalScrollBox[T]) Update(screen *Screen[T], displayX, displayY, displayWidth, displayHeight, x, y int) {
	// The code below assumes x and y are both 0.
	if x != 0 || y != 0 {
		panic("todo: x and y inside VerticalScrollBox")
	}

	// Draw fixed top area.
	if b.top != nil {
		b.top.Update(screen, displayX, displayY, displayWidth, int(b.topHeight), 0, 0)
		displayY += int(b.topHeight)
		displayHeight -= int(b.topHeight)
	}

	// Draw fixed bottom area.
	if b.bottom != nil {
		bottomH := int(b.bottomHeight)
		b.bottom.Update(screen, displayX, displayY+displayHeight-bottomH, displayWidth, bottomH, 0, 0)
		displayHeight -= bottomH
	}

	// Now follows updating the complicated part: updating the child.
	// We return after it has been updated.

	// Try to modify the screen scroll mode.
	hwscroll := false
	if b.Rect.parent == nil {
		line := int(b.topHeight) + b.scrollOffset%int(b.childHeight)
		hwscroll = screen.setScroll(b.topHeight, b.bottomHeight, int16(line))
	}

	// Draw the scrollable child.
	if !hwscroll {
		if b.scrollOffset != b.lastScrollOffset {
			// Fallback: update entire child on scroll.
			b.child.RequestUpdate()
		}
		b.lastScrollOffset = b.scrollOffset
		b.child.Update(screen, displayX, displayY, displayWidth, displayHeight, 0, b.scrollOffset)
		return
	}

	// Check whether the entire screen was changed, in which case we simply
	// update everything.
	diff := b.scrollOffset - b.lastScrollOffset
	if diff >= int(b.childHeight) || diff <= -int(b.childHeight) {
		b.child.RequestUpdate()
		b.child.Update(screen, displayX, displayY, displayWidth, displayHeight, 0, b.scrollOffset)
		b.lastScrollOffset = b.scrollOffset
		return
	}

	// The screen may have been moved, but at least some part is still visible.
	if b.scrollOffset > b.lastScrollOffset {
		// Moved up.
		newArea := b.scrollOffset - b.lastScrollOffset // size of newly exposed area
		existingArea := int(b.childHeight) - newArea   // size of moved area

		// Update the moved part at the top as usual.
		b.child.Update(screen, displayX, displayY, displayWidth, existingArea, 0, b.scrollOffset)

		// Update the newly exposed area at the bottom, using a bit of a hack:
		// by requesting an update and then drawing only the newly exposed part.
		b.child.RequestUpdate()
		scrollLine := (b.scrollOffset + existingArea) % int(b.childHeight)
		if scrollLine+newArea >= displayHeight {
			// The newly exposed area straddles the scroll line.
			// Therefore, we have to do the update in two parts, one before the
			// scroll line and one after.
			updateHeight1 := displayHeight - scrollLine
			updateHeight2 := newArea - updateHeight1
			b.child.Update(screen, displayX, displayY+scrollLine, displayWidth, updateHeight1, 0, b.scrollOffset+existingArea)
			b.child.RequestUpdate()
			b.child.Update(screen, displayX, displayY, displayWidth, updateHeight2, 0, b.scrollOffset+existingArea+updateHeight1)
		} else {
			// The newly exposed area doesn't cross a scroll line, so we can do
			// this update in one go.
			b.child.Update(screen, displayX, displayY+scrollLine, displayWidth, newArea, 0, b.scrollOffset+existingArea)
		}
	} else if b.scrollOffset < b.lastScrollOffset {
		// Moved down.
		newArea := b.lastScrollOffset - b.scrollOffset // size of newly exposed area
		existingArea := int(b.childHeight) - newArea   // size of moved area

		// Update the moved part at the bottom as usual.
		b.child.Update(screen, displayX, displayY+newArea, displayWidth, existingArea, 0, b.scrollOffset+newArea)

		// Update the newly exposed area at the top, using a bit of a hack:
		// by requesting an update and then drawing only the newly exposed part.
		b.child.RequestUpdate()
		scrollLine := b.scrollOffset % int(b.childHeight)
		if scrollLine+newArea >= displayHeight {
			// The newly exposed area straddles the scroll line.
			// Therefore, we have to do the update in two parts, one before the
			// scroll line and one after.
			updateHeight1 := displayHeight - scrollLine
			updateHeight2 := newArea - updateHeight1
			b.child.Update(screen, displayX, displayY+scrollLine, displayWidth, updateHeight1, 0, b.scrollOffset)
			b.child.RequestUpdate()
			b.child.Update(screen, displayX, displayY, displayWidth, updateHeight2, 0, b.scrollOffset+updateHeight1)
		} else {
			// The newly exposed area doesn't cross a scroll line, so we can do
			// this update in one go.
			b.child.Update(screen, displayX, displayY+scrollLine, displayWidth, newArea, 0, b.scrollOffset)
		}
	} else {
		b.child.Update(screen, displayX, displayY, displayWidth, displayHeight, 0, b.scrollOffset)
	}
	b.lastScrollOffset = b.scrollOffset

	b.NeedsUpdate() // clear update flag
}

func (b *VerticalScrollBox[T]) HandleEvent(event Event, x, y int) {
	switch event {
	case TouchStart:
		b.lastTouchY = int16(y)
	case TouchMove:
		// Add the last distance moved to the scrollOffset.
		scrollOffset := b.scrollOffset + (int(b.lastTouchY) - y)
		if scrollOffset < 0 {
			scrollOffset = 0
		}
		if scrollOffset > b.maxScrollOffset {
			scrollOffset = b.maxScrollOffset
		}
		b.scrollOffset = scrollOffset
		b.lastTouchY = int16(y)
	case TouchTap:
		if y < int(b.topHeight) {
			b.top.HandleEvent(event, x, y)
		} else if y < int(b.topHeight)+int(b.childHeight) {
			b.child.HandleEvent(event, x, y-int(b.topHeight)+b.scrollOffset)
		} else {
			b.bottom.HandleEvent(event, x, y-int(b.topHeight)-int(b.childHeight))
		}
	}
}

func (b *VerticalScrollBox[T]) RequestUpdate() {
	b.Rect.RequestUpdate()
	if b.top != nil {
		b.top.RequestUpdate()
	}
	if b.bottom != nil {
		b.bottom.RequestUpdate()
	}
	b.child.RequestUpdate()
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
