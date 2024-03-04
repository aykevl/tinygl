package image

import (
	"errors"

	"tinygo.org/x/drivers/pixel"
)

// Quite OK Image format, a format that's similar to PNG but is much simpler to
// encode/decode.
//
// For more information, see: https://qoiformat.org/
type QOI[T pixel.Color] struct {
	data          []byte
	width, height int
	state         qoiDecodeState[T]
}

const (
	qoiColorSpaceSRGB   = iota // sRGB with linear alpha
	qoiColorSpaceLinear        // all channels linear
)

var (
	errQOIHeaderTooSmall        = errors.New("image: QOI header too small")
	errQOIInvalidMagic          = errors.New("image: QOI magic missing")
	errQOIUnsupportedChannelNum = errors.New("image: unknown QOI channel number")
	errQOIUnsupportedColorSpace = errors.New("image: unknown QOI color space")
)

// NewQOI loads a new QOI image, with the data in the binary string. It parses
// the header and returns any error it finds, but it doesn't parse the file
// itself.
//
// The best way to embed the data is using the Go embed package.
func NewQOI[T pixel.Color](data string) (*QOI[T], error) {
	if len(data) < 14 {
		return nil, errQOIHeaderTooSmall
	}
	if data[:4] != "qoif" {
		return nil, errQOIInvalidMagic
	}
	if data[12] != 3 {
		// We only support RGB, we don't support the alpha channel (yet!).
		return nil, errQOIUnsupportedChannelNum
	}
	if data[13] != qoiColorSpaceSRGB {
		// The colorspace affects how a pixel.Color is encoded.
		return nil, errQOIUnsupportedColorSpace
	}
	img := &QOI[T]{
		data: []byte(data),
	}
	img.width = int(uint32(data[4])<<24 | uint32(data[5])<<16 | uint32(data[6])<<8 | uint32(data[7])<<0)
	img.height = int(uint32(data[8])<<24 | uint32(data[9])<<16 | uint32(data[10])<<8 | uint32(data[11])<<0)
	img.state.reset()
	return img, nil
}

func NewQOIFromBytes[T pixel.Color](data []byte) (*QOI[T], error) {
	if len(data) < 14 {
		return nil, errQOIHeaderTooSmall
	}
	if string(data[:4]) != "qoif" {
		return nil, errQOIInvalidMagic
	}
	if data[12] != 3 {
		// We only support RGB, we don't support the alpha channel (yet!).
		return nil, errQOIUnsupportedChannelNum
	}
	if data[13] != qoiColorSpaceSRGB {
		// The colorspace affects how a pixel.Color is encoded.
		return nil, errQOIUnsupportedColorSpace
	}
	img := &QOI[T]{
		data: data,
	}
	img.width = int(uint32(data[4])<<24 | uint32(data[5])<<16 | uint32(data[6])<<8 | uint32(data[7])<<0)
	img.height = int(uint32(data[8])<<24 | uint32(data[9])<<16 | uint32(data[10])<<8 | uint32(data[11])<<0)
	img.state.reset()
	return img, nil
}

func (img *QOI[T]) Size() (width, height int) {
	return int(img.width), int(img.height)
}

// Special pixel format. It is faster to use a full 32-bit value than to use 4
// individual RGB values as with color.RGBA.
type qoiPixel uint32

func makeQOIPixel(r, g, b, a uint8) qoiPixel {
	return qoiPixel(a)<<24 | qoiPixel(b)<<16 | qoiPixel(g)<<8 | qoiPixel(r)<<0
}

func (p qoiPixel) RGBA() (r, g, b, a uint8) {
	return uint8(p >> 0), uint8(p >> 8), uint8(p >> 16), uint8(p >> 24)
}

func (p qoiPixel) add(r, g, b uint8) qoiPixel {
	// TODO: convince the compiler that it can use SIMD (UADD8) on Cortex-M4.
	r0, g0, b0, _ := p.RGBA()
	return p&0xff00_0000 | qoiPixel(b0+b)<<16 | qoiPixel(g0+g)<<8 | qoiPixel(r0+r)<<0
}

func getQOIPixel[T pixel.Color](p qoiPixel) T {
	r, g, b, _ := p.RGBA()
	return pixel.NewColor[T](r, g, b)
}

// QOI decode state, that can be saved and restored as needed.
type qoiDecodeState[T pixel.Color] struct {
	index          uint
	y              int
	pixel          qoiPixel
	previouslySeen [64]qoiPixel
	run            int
	runColor       T
}

// Initialize the decode state to start decoding at the first pixel in the
// image.
func (state *qoiDecodeState[T]) reset() {
	*state = qoiDecodeState[T]{
		index: 14,
		pixel: makeQOIPixel(0, 0, 0, 255),
	}
}

func (img *QOI[T]) Draw(buf pixel.Image[T], bufX, bufY, scale int) {
	// Here is the specification for this image format:
	// https://qoiformat.org/qoi-specification.pdf
	// This decoder has been written directy from specification.

	if scale != 1 {
		panic("todo: scale")
	}

	if img.state.y > bufY {
		// Can't reuse decode state, so restart from the beginning.
		// If bufX/bufY aren't zero, it means the first part of the image is
		// decoded but the result is discarded.
		img.state.reset()
	}

	_, bufHeight := buf.Size()
	state := &img.state
	run := state.run
	for y := state.y; y < img.height && y-bufY < bufHeight; y++ {
		state.y = y + 1
		for x := 0; x < img.width; x++ {
			if run > 0 {
				run--
				setPixel(buf, x-bufX, y-bufY, state.runColor)
				continue
			}
			opcode := img.data[state.index]
			switch opcode >> 6 {
			case 0b00: // QOI_OP_INDEX
				state.index++
				index := opcode % 64
				state.pixel = state.previouslySeen[index]
				setPixel(buf, x-bufX, y-bufY, getQOIPixel[T](state.pixel))
				continue // no need to store this pixel again
			case 0b01: // QOI_OP_DIFF
				state.index++
				state.pixel = state.pixel.add((opcode>>4)%4-2, (opcode>>2)%4-2, (opcode>>0)%4-2)
				setPixel(buf, x-bufX, y-bufY, getQOIPixel[T](state.pixel))
			case 0b10: // QOI_OP_LUMA
				byte1 := img.data[state.index+1]
				state.index += 2
				diffGreen := (opcode % 64) - 32
				diffRed := byte1>>4 - 8 + diffGreen
				diffBlue := byte1%16 - 8 + diffGreen
				state.pixel = state.pixel.add(diffRed, diffGreen, diffBlue)
				setPixel(buf, x-bufX, y-bufY, getQOIPixel[T](state.pixel))
			case 0b11:
				if opcode == 0b11111110 { // QOI_OP_RGB
					r := img.data[state.index+1]
					g := img.data[state.index+2]
					b := img.data[state.index+3]
					_, _, _, a := state.pixel.RGBA()
					state.pixel = makeQOIPixel(r, g, b, a)
					state.index += 4
					setPixel(buf, x-bufX, y-bufY, getQOIPixel[T](state.pixel))
				} else if opcode == 0b11111111 { // QOI_OP_RGBA
					r := img.data[state.index+1]
					g := img.data[state.index+2]
					b := img.data[state.index+3]
					a := img.data[state.index+4]
					state.pixel = makeQOIPixel(r, g, b, a)
					state.index += 5
					setPixel(buf, x-bufX, y-bufY, getQOIPixel[T](state.pixel))
				} else { // QOI_OP_RUN
					state.index++
					length := int(opcode%64) + 1
					state.runColor = getQOIPixel[T](state.pixel)
					setPixel(buf, x-bufX, y-bufY, state.runColor)
					run = length - 1
				}
			default:
				// unreachable, all possible cases have been covered
			}

			// Save the pixel in the index, for QOI_OP_INDEX.
			r, g, b, a := state.pixel.RGBA()
			indexPosition := (r*3 + g*5 + b*7 + a*11) % 64
			state.previouslySeen[indexPosition] = state.pixel
		}
	}
	state.run = run
}
