package tinygl

type Event uint8

const (
	NoEvent Event = iota
	TouchStart
	TouchMove
	TouchEnd
	TouchTap
)

// Set the position of the current touch point, or (-1, -1) if nothing currently
// touching the screen.
//
// TODO: handle multitouch. This will require an API change. In other words,
// this API is going to change at some point in the future.
func (s *Screen[T]) SetTouchState(x, y int16) {
	if x >= 0 || y >= 0 {
		// Currently touching.
		if s.touchEvent == NoEvent {
			// This is the start of a touch event.
			s.touchEvent = TouchTap
			s.touchX = x
			s.touchY = y
			s.child.HandleEvent(TouchStart, int(x), int(y))
		} else if s.touchEvent == TouchTap {
			// Continuing touch event.
			if int(s.touchX) != int(x) || int(s.touchY) != int(y) {
				// The touch point moved. Treat it as a move event.
				// The caller is responsible for implementing some form of
				// hysteresis to detect actual movement instead of just ADC
				// noise.
				s.touchEvent = TouchMove
				s.child.HandleEvent(TouchMove, int(x), int(y))
			}
			// TODO: long press, double tap, etc.
		} else if s.touchEvent == TouchMove {
			s.child.HandleEvent(TouchMove, int(x), int(y))
		}
	} else if s.touchEvent != NoEvent {
		// Not touching anymore: the touch event ended.
		s.child.HandleEvent(TouchEnd, -1, -1)
		if s.touchEvent == TouchTap {
			// Notify the tap event.
			s.child.HandleEvent(TouchTap, int(s.touchX), int(s.touchY))
		}
		// ...and reset the touch state to the initial state.
		s.touchEvent = NoEvent
		s.touchX = x
		s.touchY = y
	}
}

// Abs returns the absolute value of x.
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
