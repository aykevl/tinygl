package basic

import (
	"github.com/aykevl/tinygl"
)

// Create a new text label with the default style.
func (theme *Basic[T]) NewText(text string) *tinygl.Text[T] {
	return tinygl.NewText(theme.Font, theme.Foreground, theme.Background, text)
}

// NewVBox returns a new VBox with the default style.
func (theme *Basic[T]) NewVBox(children ...tinygl.Object[T]) *tinygl.VBox[T] {
	return tinygl.NewVBox(theme.Background, children...)
}

// Create a new listbox with the given elements. The elements (and number of
// elements) cannot be changed after creation.
func (theme *Basic[T]) NewListBox(elements []string) *tinygl.ListBox[T] {
	return tinygl.NewListBox(theme.Font, theme.Foreground, theme.Background, theme.Tint, elements)
}
