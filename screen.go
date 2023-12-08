package tinygl

import (
	"time"
	"unsafe"

	"tinygo.org/x/drivers"
	"tinygo.org/x/drivers/pixel"
)

const showStats = false

// The Displayer that is drawn to.
type Displayer[T pixel.Color] interface {
	Size() (int16, int16)
	DrawBitmap(x, y int16, bitmap pixel.Image[T]) error
	Display() error
	Rotation() drivers.Rotation
}

// Asynchronous display. Some displays support asynchronously sending data in
// the background (usually using DMA). We will transparently use this feature if
// available to speed up display updates.
type AsyncDisplay[T pixel.Color] interface {
	Displayer[T]
	IsAsync() bool
	StartDrawBitmap(x, y int16, bitmap pixel.Image[T]) error
	Wait() error
}

// ScrollableDisplay is a display that supports hardware scrolling.
// A display doesn't have to support it, but if it does, scrolling can become a
// lot more smooth.
type ScrollableDisplay[T pixel.Color] interface {
	Displayer[T]
	SetScrollArea(topFixedArea, bottomFixedArea int16)
	SetScroll(line int16)
	StopScroll()
}

type Screen[T pixel.Color] struct {
	display           Displayer[T]
	scrollableDisplay ScrollableDisplay[T]
	asyncDisplay      AsyncDisplay[T]
	child             Object[T]
	buffer1           pixel.Image[T]
	buffer2           pixel.Image[T]
	statPixels        int
	statBuffers       uint16
	ppi               int16
	touchX            int16
	touchY            int16
	touchEvent        Event
	scrolling         bool
	bufferState       uint8 // see Buffer() for possible values
}

// NewScreen creates a new screen to fill the whole display.
// The buffer needs to be big enough to fill at least horizontal row of pixels,
// but should preferably be bigger (10% of the screen for example).
// The ppi parameter is the number of pixels per inch, which is important
// for touch events.
func NewScreen[T pixel.Color](display Displayer[T], buffer pixel.Image[T], ppi int) *Screen[T] {
	width, height := display.Size()
	maxSize := width
	if height > width {
		maxSize = height
	}
	if buffer.Len() < int(maxSize) {
		panic("buffer too small")
	}
	hwscroll, _ := display.(ScrollableDisplay[T])
	buffer = buffer.Rescale(1, buffer.Len())
	screen := &Screen[T]{
		display:           display,
		scrollableDisplay: hwscroll,
		buffer1:           buffer,
		ppi:               int16(ppi),
	}

	// Try to use async operations for the display if possible. The display
	// needs to support it and we need to have a buffer that's large enough that
	// we can split it in two.
	if asyncDisplay, ok := display.(AsyncDisplay[T]); ok && asyncDisplay.IsAsync() {
		top, bottom := buffer.Split(buffer.Len() / 2)
		if top.Len() >= int(maxSize) && bottom.Len() >= int(maxSize) {
			screen.buffer1 = top
			screen.buffer2 = bottom
			screen.asyncDisplay = asyncDisplay
		}
	}

	return screen
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
	if s.scrolling {
		// TODO: use StopScroll only after updating the child.
		// This makes it possible to update the screen with a monochrome
		// animation that hides the StopScroll flash.
		s.scrollableDisplay.StopScroll()
		s.scrolling = false
	}
	s.child = child
	child.RequestUpdate()
}

// Layout determines the size for all objects in the screen.
// This is called from Update() so it normally doesn't need to be called
// manually, but it can sometimes be helpful to know the size of objects before
// doing further initialization, for example when drawing on a canvas.
func (s *Screen[T]) Layout() {
	width, height := s.display.Size()
	s.child.Layout(int(width), int(height))
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
	width, height := s.display.Size()
	s.child.Update(s, 0, 0, int(width), int(height), 0, 0)
	if showStats {
		duration := time.Since(start)
		println("sent", s.statPixels, "pixels using", s.statBuffers, "buffers in", duration.String())
	}
	if s.asyncDisplay != nil {
		s.asyncDisplay.Wait()
	}
	err := s.display.Display()
	if err != nil {
		return err
	}
	return nil
}

// Buffer returns a buffer that can be painted and then sent to the screen.
// After it has been painted, it must be passed to Send.
func (s *Screen[T]) Buffer() pixel.Image[T] {
	if s.asyncDisplay == nil {
		// No async display, so only buffer1 is used.
		return s.buffer1
	}

	// We can be in one of four possible states:
	//  0. buffer 1 is available for use, buffer 2 may or may not be currently
	//     being sent to the display
	//  1. buffer 1 is borrowed for drawing, buffer 2 might still be sending
	//  2. buffer 1 is being sent, buffer 2 is ready for use
	//  3. buffer 2 is borrowed for drawing, buffer 1 might still be sending

	if s.bufferState%2 != 0 {
		// Currently in state 1 or 3, so no buffers are available.
		panic("tinygl: called Buffer() twice in a row")
	}
	s.bufferState++ // move to the next state (1 or 3)
	if s.bufferState == 1 {
		return s.buffer1
	} else {
		return s.buffer2
	}
}

// Send sends a buffer to the given coordinates. After calling Send, the buffer
// must not be reused. Instead, a new buffer must be obtained from Buffer.
//
// This function is used internally, and should only be used to implement custom
// widgets.
func (s *Screen[T]) Send(x, y int, buffer pixel.Image[T]) {
	var startWait, startSend time.Time
	if showStats {
		var zeroColor T
		s.statBuffers++
		s.statPixels += buffer.Len() * int(unsafe.Sizeof(zeroColor))
		startWait = time.Now()
	}

	if s.asyncDisplay == nil {
		// Asynchronous updates are not supported, so send the buffer
		// synchronously.
		s.display.DrawBitmap(int16(x), int16(y), buffer)
		if showStats {
			duration := time.Since(startWait)
			println("buffer send:", len(buffer.RawBuffer()), duration.String())
		}
		return
	}

	// Async updates are supported, so use them.
	// First make sure we track buffers correctly.
	if s.bufferState%2 == 0 {
		// We have two buffers and switch between them.
		// Sending more than one at a time is not supported.
		// For more information, see s.Buffer().
		panic("tinygl: submitted two buffers in a row")
	}
	s.bufferState = (s.bufferState + 1) % 4

	// Wait until the previous buffer has been fully sent.
	s.asyncDisplay.Wait()

	// Start sending the new buffer.
	if showStats {
		startSend = time.Now()
	}
	s.asyncDisplay.StartDrawBitmap(int16(x), int16(y), buffer)

	if showStats {
		sendDuration := time.Since(startSend)
		waitDuration := startSend.Sub(startWait)
		println("buffer send:", len(buffer.RawBuffer()), waitDuration.String(), sendDuration.String())
	}
}

func (s *Screen[T]) setScroll(topFixed, bottomFixed, line int16) bool {
	if !s.scrolling {
		if s.scrollableDisplay == nil {
			return false // not scrollable
		}
		switch s.display.Rotation() {
		case drivers.Rotation0, drivers.Rotation180:
			// good
		default:
			return false // scrolls in the wrong direction
		}
		s.scrolling = true
		s.scrollableDisplay.SetScrollArea(topFixed, bottomFixed)
	}
	s.scrollableDisplay.SetScroll(line)
	return true
}

// Internal function. Do not use directly except in custom widgets.
//
// It paints the given area on screen with the given color.
func PaintSolidColor[T pixel.Color](s *Screen[T], color T, x, y, width, height int) {
	lineStart := 0
	for lineStart < height {
		buf := s.Buffer()
		lines := buf.Len() / width
		if lineStart+lines > height {
			lines = height - lineStart
		}
		buf = buf.Rescale(width, lines)
		buf.FillSolidColor(color)
		s.Send(x, y+lineStart, buf)
		lineStart += lines
	}
}
