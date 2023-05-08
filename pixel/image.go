package pixel

import (
	"image/color"
	"unsafe"
)

type Image[T Color] struct {
	// note: no stride because otherwise Buffer() won't work
	width  int16
	height int16
	data   unsafe.Pointer
}

func NewImage[T Color](width, height int) Image[T] {
	buf := make([]T, width*height)
	return Image[T]{
		width:  int16(width),
		height: int16(height),
		data:   unsafe.Pointer(&buf[0]),
	}
}

func NewImageFromBuffer[T Color](buffer []T, width int) Image[T] {
	height := len(buffer) / width
	if len(buffer) != width*height {
		panic("buffer of unexpected size")
	}
	return Image[T]{
		width:  int16(width),
		height: int16(height),
		data:   unsafe.Pointer(&buffer[0]),
	}
}

// Rescale returns a new Image buffer based on the img buffer.
// The contents is undefined after the Rescale operation, and any modification
// to the returned image will overwrite the underlying image buffer in undefined
// ways. It will panic if width*height is larger than img.Len().
func (img Image[T]) Rescale(width, height int) Image[T] {
	if width*height > img.Len() {
		panic("Image.Rescale size out of bounds")
	}
	return Image[T]{
		width:  int16(width),
		height: int16(height),
		data:   img.data,
	}
}

// LimitHeight returns a subimage with the bottom part cut off, as specified by
// height.
func (img Image[T]) LimitHeight(height int) Image[T] {
	if height < 0 || height > int(img.height) {
		panic("Image.LimitHeight: out of bounds")
	}
	return Image[T]{
		width:  img.width,
		height: int16(height),
		data:   img.data,
	}
}

// Len returns the number of pixels in this image buffer.
func (img Image[T]) Len() int {
	return int(img.width) * int(img.height)
}

func (img Image[T]) Buffer() []T {
	return unsafe.Slice((*T)(img.data), int(img.width)*int(img.height))
}

func (img Image[T]) RawBuffer() []uint8 {
	var zeroColor T
	numBytes := int(unsafe.Sizeof(zeroColor)) * int(img.width) * int(img.height)
	return unsafe.Slice((*byte)(img.data), numBytes)
}

func (img Image[T]) Size() (int, int) {
	return int(img.width), int(img.height)
}

func (img Image[T]) Set(x, y int, c T) {
	var zeroColor T
	offset := (y*int(img.width) + x) * int(unsafe.Sizeof(zeroColor))
	ptr := unsafe.Add(img.data, offset)
	*((*T)(ptr)) = c
}

func (img Image[T]) Get(x, y int) T {
	var zeroColor T
	offset := (y*int(img.width) + x) * int(unsafe.Sizeof(zeroColor))
	ptr := unsafe.Add(img.data, offset)
	return *((*T)(ptr))
}

// FillSolidColor fills the entire image with the given color.
// This may be faster than setting individual pixels.
func (img Image[T]) FillSolidColor(color T) {
	ptr := img.data
	var zeroColor T
	for i := 0; i < img.Len(); i++ {
		// TODO: this can be optimized a lot.
		// - The store can be done as a 32-bit integer, after checking for
		//   alignment.
		// - Perhaps the loop can be unrolled to improve copy performance.
		*(*T)(ptr) = color
		ptr = unsafe.Add(ptr, unsafe.Sizeof(zeroColor))
	}
}

// Color is a helper to easily get a color T from R/G/B.
func (img Image[T]) Color(r, g, b uint8) T {
	return NewColor[T](r, g, b)
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
