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

	var heightSum int
	var maxWidth int
	for _, child := range children {
		child.SetParent(box)
		childWidth, childHeight := child.minSize()
		if childWidth > maxWidth {
			maxWidth = childWidth
		}
		heightSum += childHeight
	}

	box.Rect.init(background, maxWidth, heightSum)
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
	// TODO: update b.minHeight?
	var minHeightSum int
	var growableSum int
	for _, child := range b.children {
		_, childGrowable := child.growable()
		_, childHeight := child.minSize()
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
		_, childHeight := child.minSize()
		if childGrowable != 0 {
			growPixels := leftoverHeight * childGrowable / leftoverGrowable
			childHeight += growPixels
			leftoverHeight -= growPixels
			leftoverGrowable -= childGrowable
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
