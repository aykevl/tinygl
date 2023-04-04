package tinygl

import (
	"time"
	"unsafe"

	"github.com/aykevl/tinygl/pixel"
)

const showStats = false

// The Displayer that is drawn to.
type Displayer interface {
	Size() (int16, int16)
	DrawRGBBitmap8(x, y int16, data []byte, w, h int16) error
	Display() error
}

type Screen[T pixel.Color] struct {
	display     Displayer
	child       Object[T]
	buffer      []T
	statBuffers uint16
	statPixels  int
}

func NewScreen[T pixel.Color](display Displayer, buffer []T) *Screen[T] {
	width, height := display.Size()
	maxSize := width
	if height > width {
		maxSize = height
	}
	if len(buffer) < int(maxSize) {
		panic("buffer too small")
	}
	return &Screen[T]{
		display: display,
		buffer:  buffer,
	}
}

// SetChild sets the root child. This will typically be a container of some
// sort.
func (s *Screen[T]) SetChild(child Object[T]) {
	if s.child != nil {
		panic("child already set")
	}
	s.child = child
}

// Layout determines the size for all objects in the screen.
// This is called from Update() so it normally doesn't need to be called
// manually, but it can sometimes be helpful to know the size of objects before
// doing further initialization, for example when drawing on a canvas.
func (s *Screen[T]) Layout() {
	width, height := s.display.Size()
	s.child.Layout(0, 0, int(width), int(height))
}

// Update sends all changes in the screen to the (hardware) display.
func (s *Screen[T]) Update() error {
	var start time.Time
	if showStats {
		s.statBuffers = 0
		s.statPixels = 0
		start = time.Now()
	}
	s.Layout()
	s.child.Update(s)
	if showStats {
		duration := time.Since(start)
		println("sent", s.statPixels, "bytes using", s.statBuffers, "buffers in", duration.String())
	}
	return s.display.Display()
}

// Buffer returns the pixel buffer used for sending data to the screen. It can
// be used inside an Update call.
func (s *Screen[T]) Buffer() []T {
	return s.buffer
}

// Internal function: send an image buffer to the given coordinates.
func (s *Screen[T]) Send(buffer []T, x, y, width, height int) {
	rawBuffer := pixel.BufferFromSlice(buffer)
	var start time.Time
	if showStats {
		var zeroColor T
		s.statBuffers++
		s.statPixels += len(buffer) * int(unsafe.Sizeof(zeroColor))
		start = time.Now()
	}
	s.display.DrawRGBBitmap8(int16(x), int16(y), rawBuffer, int16(width), int16(height))
	if showStats {
		duration := time.Since(start)
		println("buffer send:", len(rawBuffer), duration.String())
	}
}
