package font_test

import (
	"flag"
	"image"
	"image/color"
	"image/png"
	"os"
	"testing"

	"github.com/aykevl/tinygl/font"
	"github.com/aykevl/tinygl/font/roboto"
	"tinygo.org/x/drivers/pixel"
)

var flagUpdate = flag.Bool("update", false, "update tests")

func TestRoboto(t *testing.T) {
	t.Run("Roboto Regular 16", func(t *testing.T) {
		testFont(t, roboto.Regular16, "testdata/regular16.png")
	})
	t.Run("Roboto Regular 24", func(t *testing.T) {
		testFont(t, roboto.Regular24, "testdata/regular24.png")
	})
	t.Run("Roboto Regular 36", func(t *testing.T) {
		testFont(t, roboto.Regular36, "testdata/regular36.png")
	})
	t.Run("Roboto Regular 48", func(t *testing.T) {
		testFont(t, roboto.Regular48, "testdata/regular48.png")
	})
}

func testFont(t *testing.T, face font.Font, filename string) {
	var (
		white = pixel.NewRGB888(255, 255, 255)
		black = pixel.NewRGB888(0, 0, 0)
	)
	var lines = []string{
		"The quick brown fox jumps over the lazy dog.",
		"abcdefghijklmnopqrstuvwxyz",
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		" !\"#$%&\\'()*+,-./0123456789:;<=>?@",
		"[\\]^_`{|}~",
	}

	// Draw text onto buffer.
	var width = face.Size() * 24
	var height = len(lines) * face.Height()
	buf := pixel.NewImage[pixel.RGB888](width, height)
	buf.FillSolidColor(white)
	for i, line := range lines {
		font.Draw(face, line, 8, i*face.Height()+face.Ascent(), black, buf)
	}

	// Convert buffer to Go image.
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			c := buf.Get(x, y)
			img.SetRGBA(x, y, color.RGBA{c.R, c.G, c.B, 255})
		}
	}

	// Compare output.
	golden := loadImage(t, filename)
	if x, y, equal := isSameImage(img, golden); !equal {
		if *flagUpdate {
			// -update flag was passed, so update the PNG file.
			f, err := os.Create(filename)
			if err != nil {
				t.Fatal("could not open test output:", err)
			}
			defer f.Close()
			png.Encode(f, img)
		} else {
			t.Errorf("image %s is not the same (mismatch at x=%d, y=%d)", filename, x, y)
		}
	}
}

func loadImage(t *testing.T, path string) image.Image {
	f, err := os.Open(path)
	if err != nil {
		if *flagUpdate {
			return nil
		}
		t.Fatal("could not open test output:", err)
	}
	defer f.Close()
	golden, err := png.Decode(f)
	if err != nil {
		t.Fatal("could not decode test output:", err)
	}
	return golden
}

func isSameImage(image1, image2 image.Image) (x, y int, equal bool) {
	if image1 == nil || image2 == nil {
		return 0, 0, false
	}
	if image1.Bounds() != image2.Bounds() {
		return 0, 0, false
	}
	rect := image1.Bounds().Size()
	for y := 0; y < rect.Y; y++ {
		for x := 0; x < rect.X; x++ {
			r1, g1, b1, a1 := image1.At(x, y).RGBA()
			r2, g2, b2, a2 := image2.At(x, y).RGBA()
			if r1 != r2 || g1 != g2 || b1 != b2 || a1 != a2 {
				return x, y, false
			}
		}
	}

	return 0, 0, true
}
