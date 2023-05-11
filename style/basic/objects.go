package basic

import (
	"github.com/aykevl/tinygl"
	"github.com/aykevl/tinygl/pixel"
)

// Create a new text label with the default style.
func (theme *Basic[T]) NewText(text string) *tinygl.Text[T] {
	return tinygl.NewText(theme.Font, theme.Foreground, theme.Background, text)
}

// NewVBox returns a new VBox with the default style.
func (theme *Basic[T]) NewVBox(children ...tinygl.Object[T]) *tinygl.VBox[T] {
	return tinygl.NewVBox(theme.Background, children...)
}

// A scrollable list of strings, of which one is currently selected.
type ListBox[T pixel.Color] struct {
	tinygl.Rect[T]
	children   []tinygl.Text[T]
	handler    func(tinygl.Event, int) // event handler
	selected   int16
	foreground T
	tint       T
	numCols    uint8 // number of columns, to make it a grid
}

// Create a new listbox with the given elements. The elements (and number of
// elements) cannot be changed after creation.
func (theme *Basic[T]) NewListBox(elements []string) *ListBox[T] {
	// Avoid some heap allocations by allocating all children at once.
	children := make([]tinygl.Text[T], len(elements))
	box := &ListBox[T]{
		Rect:       tinygl.MakeRect(theme.Background),
		children:   children,
		selected:   -1,
		foreground: theme.Foreground,
		tint:       theme.Tint,
		numCols:    1,
	}
	for i, text := range elements {
		child := &children[i]
		*child = tinygl.MakeText(theme.Font, theme.Foreground, theme.Background, text)
		child.SetParent(&box.Rect)
		child.SetAlign(tinygl.AlignLeft)
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
func (box *ListBox[T]) SetAlign(align tinygl.TextAlign) {
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
func (box *ListBox[T]) Layout(x, y, width, height int) {
	displayX, displayY, displayWidth, displayHeight := box.Rect.Bounds()
	if x != displayX || y != displayY || width != displayWidth || height != displayHeight {
		box.Rect.Layout(x, y, width, height)
	}

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
		childHeight := commonChildHeight
		if (rowIndex+1)*commonChildHeight > height {
			childHeight = height - rowIndex*commonChildHeight
		}
		if childHeight < 0 {
			childHeight = 0
		}
		childWidth := remainingWidth / (int(box.numCols) - colIndex)
		child.Layout(x+width-remainingWidth, y+rowIndex*commonChildHeight, childWidth, childHeight)
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
func (box *ListBox[T]) Update(screen *tinygl.Screen[T]) {
	needsUpdate, childNeedsUpdate := box.NeedsUpdate()
	if !needsUpdate && !childNeedsUpdate {
		return
	}

	// TODO: combine multiple children in a single buffer if possible
	if childNeedsUpdate {
		for i := range box.children {
			child := &box.children[i]
			child.Update(screen)
		}
	}

	// Draw the area below the children.
	if needsUpdate {
		displayX, displayY, displayWidth, displayHeight := box.Rect.Bounds()
		if len(box.children) == 0 {
			tinygl.PaintSolidColor(screen, box.Background(), displayX, displayY, displayWidth, displayHeight)
		} else {
			childX, childY, childWidth, childHeight := box.children[len(box.children)-1].Bounds()
			if childX+childWidth < displayX+displayWidth {
				// This is a grid, with some extra space at the end.
				// Like the area below indicated with "*":
				//
				//    #  #  #
				//    #  #  #
				//    #  *  *
				//    -  -  -
				paintX := childX + childWidth
				paintY := childY
				paintWidth := (displayX + displayWidth) - (childX + childWidth)
				paintHeight := childHeight
				tinygl.PaintSolidColor(screen, box.Background(), paintX, paintY, paintWidth, paintHeight)
			}
			if childY+childHeight < displayY+displayHeight {
				// This is the remaining area below the children.
				// Like the area indicated with "-":
				//
				//    #  #  #
				//    #  #  #
				//    #  *  *
				//    -  -  -
				paintX := displayX
				paintY := childY + childHeight
				paintWidth := displayWidth
				paintHeight := (displayY + displayHeight) - (childY + childHeight)
				tinygl.PaintSolidColor(screen, box.Background(), paintX, paintY, paintWidth, paintHeight)
			}
		}
	}
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
	if index != -1 {
		child := &box.children[box.selected]
		child.SetBackground(box.tint)
		child.SetColor(box.Background())
	}
}

// SetEventHandler sets the callback when one of the elements in the list gets
// selected.
func (box *ListBox[T]) SetEventHandler(eventHandler func(event tinygl.Event, index int)) {
	box.handler = eventHandler
}

// HandleEvent handles events such as touch events and calls the event handler
// with the given child as a parameter.
func (box *ListBox[T]) HandleEvent(event tinygl.Event, x, y int) {
	if box.handler == nil {
		return
	}
	for i, child := range box.children {
		childX, childY, childW, childH := child.Bounds()
		if childX <= x && childY <= y && childX+childW > x && childY+childH > y {
			box.handler(event, i)
			break
		}
	}
}
