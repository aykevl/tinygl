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
	display        Displayer
	child          Object[T]
	updateCallback func(screen *Screen[T])
	buffer         pixel.Image[T]
	statPixels     int
	statBuffers    uint16
	ppi            int16
	touchX         int16
	touchY         int16
	touchEvent     Event
}

// NewScreen creates a new screen to fill the whole display.
// The buffer needs to be big enough to fill at least horizontal row of pixels,
// but should preferably be bigger (10% of the screen for example).
// The ppi parameter is the number of pixels per inch, which is important
// for touch events.
func NewScreen[T pixel.Color](display Displayer, buffer pixel.Image[T], ppi int) *Screen[T] {
	width, height := display.Size()
	maxSize := width
	if height > width {
		maxSize = height
	}
	if buffer.Len() < int(maxSize) {
		panic("buffer too small")
	}
	return &Screen[T]{
		display: display,
		buffer:  buffer,
		ppi:     int16(ppi),
	}
}

// Size returns the size of the screen in pixels.
func (s *Screen[T]) Size() (width, height int) {
	w, h := s.display.Size()
	return int(w), int(h)
}

// SetChild sets the root child. This will typically be a container of some
// sort.
func (s *Screen[T]) SetChild(child Object[T]) {
	if child == s.child {
		return // nothing to do
	}
	s.child = child
	child.RequestUpdate()
}

// SetUpdateCallback sets the callback that is called after every screen update.
// It is mostly useful to read events an update the screen with those events
// (for example, to call SetTouchState).
func (s *Screen[T]) SetUpdateCallback(inputRead func(screen *Screen[T])) {
	s.updateCallback = inputRead
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
	err := s.display.Display()
	if err != nil {
		return err
	}
	if s.updateCallback != nil {
		s.updateCallback(s) // read keyboard/touch events
	}
	return nil
}

// Buffer returns the pixel buffer used for sending data to the screen. It can
// be used inside an Update call.
func (s *Screen[T]) Buffer() pixel.Image[T] {
	return s.buffer
}

// Internal function: send an image buffer to the given coordinates.
func (s *Screen[T]) Send(x, y int, buffer pixel.Image[T]) {
	rawBuffer := buffer.RawBuffer()
	var start time.Time
	if showStats {
		var zeroColor T
		s.statBuffers++
		s.statPixels += buffer.Len() * int(unsafe.Sizeof(zeroColor))
		start = time.Now()
	}
	width, height := buffer.Size()
	s.display.DrawRGBBitmap8(int16(x), int16(y), rawBuffer, int16(width), int16(height))
	if showStats && len(rawBuffer) >= 4096 {
		duration := time.Since(start)
		println("buffer send:", len(rawBuffer), duration.String())
	}
}

// Internal function. Do not use directly except in custom widgets.
//
// It paints the given area on screen with the given color.
func PaintSolidColor[T pixel.Color](s *Screen[T], color T, x, y, width, height int) {
	linesPerChunk := s.buffer.Len() / width
	if linesPerChunk > height {
		linesPerChunk = height
	}
	img := s.buffer.Rescale(width, linesPerChunk)
	img.FillSolidColor(color)
	lineStart := 0
	for lineStart < height {
		lines := linesPerChunk
		if lineStart+lines > height {
			lines = height - lineStart
		}
		s.Send(x, y+lineStart, img.LimitHeight(lines))
		lineStart += linesPerChunk
	}
}
