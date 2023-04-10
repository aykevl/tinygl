package pixel

import (
	"image/color"
	"unsafe"
)

type Image[T Color] struct {
	// note: no stride because otherwise Buffer() won't work
	width int // height can be inferred from width and len(data)
	data  []T
}

func NewImage[T Color](width, height int) Image[T] {
	return Image[T]{
		width: width,
		data:  make([]T, width*height),
	}
}

func NewImageFromBuffer[T Color](buffer []T, width int) Image[T] {
	if len(buffer) != len(buffer)/width*width {
		panic("buffer of unexpected size")
	}
	return Image[T]{
		width: width,
		data:  buffer,
	}
}

func (img Image[T]) Buffer() []T {
	return img.data
}

func (img Image[T]) RawBuffer() []uint8 {
	return BufferFromSlice(img.data)
}

func (img Image[T]) Size() (int, int) {
	return img.width, len(img.data) / img.width
}

func (img Image[T]) Set(x, y int, c T) {
	img.data[y*img.width+x] = c
}

func (img Image[T]) Get(x, y int) T {
	return img.data[y*img.width+x]
}

// Color is a helper to easily get a color T from R/G/B.
func (img Image[T]) Color(r, g, b uint8) T {
	return NewColor[T](r, g, b)
}

func (img Image[T]) SubImage(x, y, width, height int) Image[T] {
	if x != 0 || width != img.width {
		panic("Image: todo: SubImage with stride")
	}
	sub := img
	sub.data = sub.data[y*img.width : height*img.width]
	return sub
}

func BufferFromSlice[T Color](data []T) []byte {
	var zeroColor T // used for size calculation

	if len(data) == 0 {
		return nil
	}

	// Cast data (which is a []T) to a []byte slice.
	// This should be a safe operation, at least in TinyGo.
	ptr := (*uint8)(unsafe.Pointer(unsafe.SliceData(data)))
	return unsafe.Slice(ptr, len(data)*int(unsafe.Sizeof(zeroColor)))
}

// Wrapper for Image that implements the drivers.Displayer interface.
type DisplayerImage[T Color] struct {
	Image[T]
}

// SetPixel implements the Displayer interface.
func (img DisplayerImage[T]) SetPixel(x, y int16, color color.RGBA) {
	if x < 0 || y < 0 {
		return
	}
	width, height := img.Image.Size()
	if int(x) >= width || int(y) >= height {
		return
	}
	img.Set(int(x), int(y), img.Color(color.R, color.G, color.B))
}

// Size implements the Displayer interface.
func (img DisplayerImage[T]) Size() (int16, int16) {
	width, height := img.Image.Size()
	return int16(width), int16(height)
}

// Display implements the Displayer interface. It is a no-op.
func (img DisplayerImage[T]) Display() error {
	return nil
}
