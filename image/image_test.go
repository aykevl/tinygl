package image_test

import (
	"image/color"
	"image/png"
	"os"
	"testing"

	"github.com/aykevl/tinygl/image"
	"tinygo.org/x/drivers/pixel"
)

func TestQOI(t *testing.T) {
	for _, name := range []string{"tinygo-logo-noalpha"} {
		t.Run(name, func(t *testing.T) {
			data, err := os.ReadFile("testdata/" + name + ".qoi")
			if err != nil {
				t.Fatal("could not read input file:", err)
			}
			img, err := image.NewQOI[pixel.RGB888](string(data))
			if err != nil {
				t.Fatal("could not decode input file:", err)
			}
			testImage(t, img, "testdata/"+name+".png")
		})
	}
}

func testImage(t *testing.T, img image.Image[pixel.RGB888], reference string) {
	// Load the reference image.
	f, err := os.Open(reference)
	if err != nil {
		t.Fatal("could not open reference image:", err)
	}
	defer f.Close()
	golden, err := png.Decode(f)
	if err != nil {
		t.Fatal("could not decode reference image:", err)
	}

	// Check the image size.
	width, height := img.Size()
	if width != golden.Bounds().Dx() || height != golden.Bounds().Dy() {
		// Size didn't match, so don't compare contents to avoid out-of-bounds
		// errors.
		t.Fatalf("unexpected size: %dx%d", width, height)
	}

	// Decode the image.
	buf := pixel.NewImage[pixel.RGB888](width, height)
	img.Draw(buf, 0, 0, 1)

	// Check for differences in the decoded image.
	mismatch := 0
	firstX := 0
	firstY := 0
	var firstExpected, firstActual color.RGBA
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			c := buf.Get(x, y).RGBA()
			r2, g2, b2, _ := golden.At(x, y).RGBA()
			c2 := color.RGBA{R: uint8(r2 >> 8), G: uint8(g2 >> 8), B: uint8(b2 >> 8), A: 255}
			if c != c2 {
				mismatch++
				if mismatch == 1 {
					firstX = x
					firstY = y
					firstExpected = c
					firstActual = c2
				}
			}
		}
	}
	if mismatch != 0 {
		t.Errorf("mismatch found: %d pixels are different (first diff at (%d, %d), expected %v, actual %v)", mismatch, firstX, firstY, firstExpected, firstActual)
	}
}

func TestMono(t *testing.T) {
	for _, name := range []string{"tinygo-logo-noalpha"} {
		t.Run(name, func(t *testing.T) {
			data, err := os.ReadFile("testdata/" + name + ".raw")
			if err != nil {
				t.Fatal("could not read input file:", err)
			}
			img, e := image.NewMono[pixel.Monochrome](true, false, 299, 255, string(data))
			if e != nil {
				t.Fatal("could not decode input file:", e)
			}
			if x, y := img.Size(); x != 299 || y != 255 {
				t.Fatalf("unexpected size: %d, %d", x, y)
			}

			// Decode the image.
			buf := pixel.NewImage[pixel.Monochrome](299, 255)
			img.Draw(buf, 0, 0, 1)

			// TODO: check for data validity
		})
	}
}
