package gfx_test

import (
	"flag"
	"fmt"
	"image"
	"image/png"
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
			continue
		}

		// Load golden file (the file to check against).
		f, err := os.Open(filename)
		if err != nil {
			t.Error("could not open golden file:", err)
			continue
		}
		defer f.Close()
		golden, err := png.Decode(f)
		if err != nil {
			t.Error("could not read golden file:", err)
			continue
		}

		// Compare the output.
		if img.Bounds().Size() != golden.Bounds().Size() {
			size := img.Bounds().Size()
			t.Errorf("unexpected size for r=%d: %dx%d", r, size.X, size.Y)
		}
		mismatch := 0
		for y := 0; y < w; y++ {
			for x := 0; x < h; x++ {
				r1, g1, b1, _ := img.At(x, y).RGBA()
				r2, g2, b2, _ := golden.At(x, y).RGBA()
				if r1 != r2 || g1 != g2 || b1 != b2 {
					mismatch++
				}
			}
		}
		if mismatch != 0 {
			t.Errorf("mismatch found for r=%d: %d pixels are different", r, mismatch)
		}
	}
}
