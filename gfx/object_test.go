package gfx_test

import (
	"flag"
	"fmt"
	"image"
	"image/png"
	"math"
	"os"
	"testing"

	"github.com/aykevl/tinygl/gfx"
	"tinygo.org/x/drivers/pixel"
)

var flagUpdate = flag.Bool("update", false, "Update golden test files")

func TestCircle(t *testing.T) {
	for r := -1; r <= 20; r++ {
		w, h := 48, 48
		middleX := w / 2
		middleY := h / 2
		filename := fmt.Sprintf("testdata/circle-%d.png", r)

		// Create pixel buffer with circle rectangle black and a gray border.
		black := pixel.NewColor[pixel.RGB888](0, 0, 0)
		gray := pixel.NewColor[pixel.RGB888](64, 64, 64)
		white := pixel.NewColor[pixel.RGB888](255, 255, 255)
		buf := pixel.NewImage[pixel.RGB888](w, h)
		buf.FillSolidColor(gray)
		for y := middleY - r; y < middleY+r; y++ {
			for x := middleX - r; x < middleX+r; x++ {
				buf.Set(x, y, black)
			}
		}

		// Draw a circle in the middle.
		circle := gfx.NewCircle[pixel.RGB888](white, middleX, middleY, r)
		circle.Draw(0, 0, buf)

		compareTestOutput(t, buf, filename)
	}
}

func TestLine(t *testing.T) {
	t.Run("width=1", func(t *testing.T) {
		testLineWidth(t, 1, 61, 50)
	})
	t.Run("width=3", func(t *testing.T) {
		testLineWidth(t, 3, 61, 50)
	})
	t.Run("width=10", func(t *testing.T) {
		testLineWidth(t, 10, 17, 50)
	})
	t.Run("width=40", func(t *testing.T) {
		testLineWidth(t, 40, 17, 130)
	})
}

func testLineWidth(t *testing.T, lineWidth, numLines, centerRadius int) {
	const width, height = 400, 400
	const middleX = width / 2
	const middleY = height / 2

	black := pixel.NewColor[pixel.RGB888](0, 0, 0)
	white := pixel.NewColor[pixel.RGB888](255, 255, 255)
	startEndDot := pixel.NewColor[pixel.RGB888](255, 0, 0)

	buf := pixel.NewImage[pixel.RGB888](width, height)
	buf.FillSolidColor(black)

	// Draw lots of lines in a star pattern.
	// Use a prime number so we can see lines in lots of different angles (not
	// repeated 4 times).
	for i := 0; i < numLines; i++ {
		radians := float64(i) / float64(numLines) * 2 * math.Pi
		x1 := middleX + int(math.Cos(radians)*float64(centerRadius))
		y1 := middleY + int(math.Sin(radians)*float64(centerRadius))
		x2 := middleX + int(math.Cos(radians)*180)
		y2 := middleY + int(math.Sin(radians)*180)
		line := gfx.NewLine(white, x1, y1, x2, y2, lineWidth)
		line.Draw(0, 0, buf)

		// Add two dots at the start and end, to help with checking for accuracy.
		gfx.NewRect(startEndDot, x1, y1, 1, 1).Draw(0, 0, buf)
		gfx.NewRect(startEndDot, x2, y2, 1, 1).Draw(0, 0, buf)
	}
	compareTestOutput(t, buf, fmt.Sprintf("testdata/line-%d.png", lineWidth))
}

func compareTestOutput(t *testing.T, buf pixel.Image[pixel.RGB888], filename string) {
	w, h := buf.Size()

	// Convert the pixel buffer to a Go pixel buffer.
	img := image.NewRGBA(image.Rectangle{image.Point{}, image.Point{w, h}})
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			img.Set(x, y, buf.Get(x, y).RGBA())
		}
	}

	if *flagUpdate {
		// Write the image out to the testdata directory.
		f, err := os.Create(filename)
		if err != nil {
			t.Fatal("could not create golden output file:", err)
		}
		png.Encode(f, img)
		f.Close()
		return
	}

	// Load golden file (the file to check against).
	f, err := os.Open(filename)
	if err != nil {
		t.Error("could not open golden file:", err)
		return
	}
	defer f.Close()
	golden, err := png.Decode(f)
	if err != nil {
		t.Error("could not read golden file:", err)
		return
	}

	// Compare the output.
	if img.Bounds().Size() != golden.Bounds().Size() {
		size := img.Bounds().Size()
		t.Errorf("unexpected size for %s: %dx%d", filename, size.X, size.Y)
	}
	mismatch := 0
	firstX := 0
	firstY := 0
	for y := 0; y < w; y++ {
		for x := 0; x < h; x++ {
			r1, g1, b1, _ := img.At(x, y).RGBA()
			r2, g2, b2, _ := golden.At(x, y).RGBA()
			if r1 != r2 || g1 != g2 || b1 != b2 {
				mismatch++
				if mismatch == 1 {
					firstX = x
					firstY = y
				}
			}
		}
	}
	if mismatch != 0 {
		t.Errorf("mismatch found for %s: %d pixels are different, first diff at (%d, %d)", filename, mismatch, firstX, firstY)
	}
}
