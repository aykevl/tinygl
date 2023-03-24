package tinygl

import (
	"github.com/aykevl/tinygl/pixel"
	"github.com/aykevl/tinygl/style"
)

type VBox[T pixel.Color] struct {
	Rect[T]
	children []Object[T]
}

func NewVBox[T pixel.Color](base style.Style[T], children ...Object[T]) *VBox[T] {
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

	box.Rect.init(base, maxWidth, heightSum)
	box.flags |= flagNeedsUpdate
	return box
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
	// TODO: growable objects (like VBox, which grows in height but not in width)
	var heightSum int
	for _, child := range b.children {
		_, childHeight := child.minSize()
		child.Layout(x, y+heightSum, width, childHeight)
		heightSum += childHeight
	}
	if heightSum < int(b.minHeight) {
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

	if b.flags&flagNeedsUpdate != 0 {
		paintSolidColor(screen, b.background, int(b.displayX), int(b.displayY)+int(b.minHeight), int(b.displayWidth), int(b.displayHeight)-int(b.minHeight))
	}

	b.flags &^= flagNeedsUpdate | flagNeedsChildUpdate // updated, so no need to redraw next time
}

func (b *VBox[T]) Draw(x, y int, img pixel.Image[T]) {
	panic("todo: VBox.Draw")
}
