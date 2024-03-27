package tinygl

import (
	"github.com/aykevl/tinygl/font"
	"tinygo.org/x/drivers/pixel"
)

// A scrollable list of strings, of which one is currently selected.
type ListBox[T pixel.Color] struct {
	Rect[T]
	children   []Text[T]
	handler    func(Event, int) // event handler
	selected   int16
	width      int16
	foreground T
	tint       T
	numCols    uint8 // number of columns, to make it a grid
}

// Create a new listbox with the given elements. The elements (and number of
// elements) cannot be changed after creation.
func NewListBox[T pixel.Color](font font.Font, foreground, background, tint T, elements []string) *ListBox[T] {
	// Avoid some heap allocations by allocating all children at once.
	children := make([]Text[T], len(elements))

	box := &ListBox[T]{
		Rect:       MakeRect(background),
		children:   children,
		selected:   -1,
		foreground: foreground,
		tint:       tint,
		numCols:    1,
	}
	for i, text := range elements {
		child := &children[i]
		*child = MakeText(font, foreground, background, text)
		child.SetParent(box)
		child.SetAlign(AlignLeft)
	}

	return box
}

// MinSize returns the height of all the list items combined.
// The minimal width is always zero (it is expected to set to expand).
func (box *ListBox[T]) MinSize() (width, height int) {
	if len(box.children) == 0 {
		return 0, 0
	}
	firstChild := &box.children[0]
	_, childHeight := firstChild.MinSize()
	numRows := (len(box.children) + int(box.numCols) - 1) / int(box.numCols)
	return 0, childHeight * numRows
}

// Len returns the number of elements in the listbox.
func (box *ListBox[T]) Len() int {
	return len(box.children)
}

// SetColumns splits the view into multiple columns, so that elements can also
// display horizontally.
func (box *ListBox[T]) SetColumns(columns int) {
	if columns < 1 {
		columns = 1
	}
	box.numCols = uint8(columns)
}

// SetAlign sets the alignment of all text children.
func (box *ListBox[T]) SetAlign(align TextAlign) {
	for i := range box.children {
		box.children[i].SetAlign(align)
	}
}

// Set padding for each text child.
func (box *ListBox[T]) SetPadding(horizontal, vertical int) {
	for i := range box.children {
		box.children[i].SetPadding(horizontal, vertical)
	}
	box.Rect.RequestUpdate()
}

// RequestUpdate will request an update for this object and all of its children.
func (box *ListBox[T]) RequestUpdate() {
	box.Rect.RequestUpdate()
	for i := range box.children {
		box.children[i].RequestUpdate()
	}
}

// Layout implements tinygl.Object.
func (box *ListBox[T]) Layout(width, height int) {
	if !box.NeedsLayout() {
		return
	}

	box.width = int16(width)

	commonChildHeight := 0
	if len(box.children) != 0 {
		firstChild := &box.children[0]
		_, commonChildHeight = firstChild.MinSize()
	}

	colIndex := 0
	rowIndex := 0
	remainingWidth := width
	for i := range box.children {
		child := &box.children[i]
		childWidth := remainingWidth / (int(box.numCols) - colIndex)
		child.Layout(childWidth, commonChildHeight)
		remainingWidth -= childWidth

		colIndex++
		if colIndex >= int(box.numCols) {
			// Wrap again back to start.
			colIndex = 0
			rowIndex++
			remainingWidth = width
		}
	}
}

// Update implements tinygl.Object.
func (box *ListBox[T]) Update(screen *Screen[T], displayX, displayY, displayWidth, displayHeight, x, y int) {
	needsUpdate, childNeedsUpdate := box.NeedsUpdate()
	if !needsUpdate && !childNeedsUpdate {
		return
	}

	commonChildHeight := 0
	if len(box.children) != 0 {
		firstChild := &box.children[0]
		_, commonChildHeight = firstChild.MinSize()
	}

	remainingWidth := int(box.width)
	if childNeedsUpdate {
		colIndex := 0
		rowIndex := 0
		for i := range box.children {
			child := &box.children[i]

			// Calculate child offset from the (0, 0) coordinate of the ListBox.
			childOffsetX := int(box.width) - int(remainingWidth)
			childOffsetY := commonChildHeight * rowIndex

			// Calculate idealized x/y/width/height for the child, ignoring the
			// draw boundaries.
			childDisplayX := displayX - x + childOffsetX
			childDisplayY := displayY - y + childOffsetY
			childDisplayWidth := remainingWidth / (int(box.numCols) - colIndex)
			childDisplayHeight := commonChildHeight
			remainingWidth -= childDisplayWidth

			// Limit the x/y/width/height of the child to the draw area.
			childX := 0
			childY := 0
			if childDisplayX < displayX {
				childX = displayX - childDisplayX
				childDisplayX = displayX
				childDisplayWidth -= childX
			}
			if childDisplayY < displayY {
				childY = displayY - childDisplayY
				childDisplayY = displayY
				childDisplayHeight -= childY
			}
			maxChildDisplayWidth := displayX + displayWidth - childDisplayX
			if childDisplayWidth > maxChildDisplayWidth {
				childDisplayWidth = maxChildDisplayWidth
			}
			maxChildDisplayHeight := displayY + displayHeight - childDisplayY
			if childDisplayHeight > maxChildDisplayHeight {
				childDisplayHeight = maxChildDisplayHeight
			}

			if childDisplayWidth > 0 && childDisplayHeight > 0 {
				child.Update(screen, childDisplayX, childDisplayY, childDisplayWidth, childDisplayHeight, childX, childY)
			}

			colIndex++
			if colIndex >= int(box.numCols) {
				// Wrap again back to start.
				colIndex = 0
				rowIndex++
				remainingWidth = int(box.width)
			}
		}
	}

	// Draw the area below the children.
	if needsUpdate {
		if len(box.children) == 0 {
			PaintSolidColor(screen, box.Background(), displayX, displayY, displayWidth, displayHeight)
		} else {
			if len(box.children)%int(box.numCols) != 0 {
				// The grid may not fill every row entirely, like the area
				// indicated with "*" here:
				//
				//    #  #  #
				//    #  #  #
				//    #  *  *
				//    -  -  -
				//
				// Calculate the idealized paint area to draw the background
				// color.
				paintX := displayX - x + (int(box.width) - remainingWidth)
				paintY := displayY - y + len(box.children)/int(box.numCols)*commonChildHeight
				paintWidth := int(box.width) - remainingWidth
				paintHeight := commonChildHeight

				// Limit the paint area to the allowed display area.
				if paintX < displayX {
					paintWidth -= displayX - paintX
					paintX = displayX
				}
				if paintY < displayY {
					paintHeight -= displayY - paintY
					paintY = displayY
				}
				if paintWidth > displayWidth {
					paintWidth = displayWidth
				}
				if paintHeight > displayHeight {
					paintHeight = displayHeight
				}

				// Draw this area.
				if paintWidth > 0 && paintHeight > 0 {
					PaintSolidColor(screen, box.Background(), paintX, paintY, paintWidth, paintHeight)
				}
			}

			// Now do the same for the area below the children.
			// Like the area indicated with "-":
			//
			//    #  #  #
			//    #  #  #
			//    #  *  *
			//    -  -  -
			paintX := displayX - x
			paintY := displayY - y + (len(box.children)+int(box.numCols)-1)/int(box.numCols)*commonChildHeight

			// Limit the paint area.
			if paintX < displayX {
				paintX = displayX
			}
			if paintY < displayY {
				paintY = displayY
			}
			paintWidth := displayX + displayWidth - paintX
			paintHeight := displayY + displayHeight - paintY

			// Now draw this area, if there is something to draw.
			if paintWidth > 0 && paintHeight > 0 {
				PaintSolidColor(screen, box.Background(), paintX, paintY, paintWidth, paintHeight)
			}
		}
	}
}

func (box *ListBox[T]) MarkUpdated() {
	if _, anyChild := box.NeedsUpdate(); anyChild {
		for i := range box.children {
			box.children[i].MarkUpdated()
		}
	}
	box.Rect.MarkUpdated()
}

// Selected returns the index of the currently selected element, or -1 if no
// element is selected.
func (box *ListBox[T]) Selected() int {
	return int(box.selected)
}

// Select selects the given index. The index -1 means "no child selected".
func (box *ListBox[T]) Select(index int) {
	if box.selected != -1 {
		child := &box.children[box.selected]
		child.SetBackground(box.Background())
		child.SetColor(box.foreground)
	}
	if index >= len(box.children) {
		panic("ListBox.Select: out of range")
	}
	box.selected = int16(index)
	if index >= 0 {
		child := &box.children[box.selected]
		child.SetBackground(box.tint)
		child.SetColor(box.Background())
		firstChild := &box.children[0]
		_, commonChildHeight := firstChild.MinSize()
		childY := (index / int(box.numCols)) * commonChildHeight
		if box.Parent() != nil {
			box.Parent().ScrollIntoViewVertical(childY, childY+int(commonChildHeight), box)
		}
	}
}

// SetEventHandler sets the callback when one of the elements in the list gets
// selected.
func (box *ListBox[T]) SetEventHandler(eventHandler func(event Event, index int)) {
	box.handler = eventHandler
}

// HandleEvent handles events such as touch events and calls the event handler
// with the given child as a parameter.
func (box *ListBox[T]) HandleEvent(event Event, x, y int) {
	if box.handler == nil {
		return
	}
	if len(box.children) == 0 {
		return
	}

	childHeight := 0
	if len(box.children) != 0 {
		firstChild := &box.children[0]
		_, childHeight = firstChild.MinSize()
	}

	row := y / childHeight
	remainingWidth := int(box.width)
	childX := 0
	for col := 0; col < int(box.numCols); col++ {
		childWidth := remainingWidth / (int(box.numCols) - col)
		if x >= childX && x < childX+childWidth {
			index := row*int(box.numCols) + col
			if index < len(box.children) {
				box.handler(event, index)
			}
		}
		childX += childWidth
		remainingWidth -= childWidth
	}
}
