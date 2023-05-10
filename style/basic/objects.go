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
	slack      int16
	selected   int16
	foreground T
	tint       T
	textHeight int8
	numCols    uint8 // number of columns, to make it a grid
}

// Create a new listbox with the given elements. The elements (and number of
// elements) cannot be changed after creation.
func (theme *Basic[T]) NewListBox(elements []string) *ListBox[T] {
	// Avoid some heap allocations by allocating all children at once.
	children := make([]tinygl.Text[T], len(elements))
	textHeight := theme.Font.BBox[1]
	box := &ListBox[T]{
		Rect:       tinygl.MakeRect(theme.Background, 0, int(textHeight)),
		children:   children,
		textHeight: textHeight,
		selected:   -1,
		foreground: theme.Foreground,
		tint:       theme.Tint,
		numCols:    1,
	}
	for i, text := range elements {
		child := &children[i]
		*child = tinygl.MakeText(theme.Font, theme.Foreground, theme.Background, text)
		child.SetParent(box)
		child.SetAlign(tinygl.AlignLeft)
	}

	return box
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

	box.layoutChildren(x, y, width, height)

	numRows := (len(box.children) + int(box.numCols) - 1) / int(box.numCols)
	box.slack = int16(height - int(box.textHeight)*numRows)
	if box.slack > 0 {
		// More of the extra space at the end is exposed, so redraw that area.
		// TODO: only redraw newly exposed area.
		box.RequestUpdate()
	}
}

func (box *ListBox[T]) layoutChildren(x, y, width, height int) {
	colIndex := 0
	rowIndex := 0
	remainingWidth := width
	for i := range box.children {
		child := &box.children[i]
		childHeight := int(box.textHeight)
		if (rowIndex+1)*int(box.textHeight) > height {
			childHeight = height - rowIndex*int(box.textHeight)
		}
		if childHeight < 0 {
			childHeight = 0
		}
		childWidth := remainingWidth / (int(box.numCols) - colIndex)
		child.Layout(x+width-remainingWidth, y+rowIndex*int(box.textHeight), childWidth, childHeight)
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

	if needsUpdate && box.slack > 0 {
		displayX, displayY, displayWidth, displayHeight := box.Rect.Bounds()
		tinygl.PaintSolidColor(screen, box.Background(), displayX, displayY+displayHeight-int(box.slack), displayWidth, int(box.slack))
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
