package tinygl_test

import (
	"flag"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
	"path/filepath"
	"testing"

	"github.com/aykevl/tinygl"
	"tinygo.org/x/drivers"
	"tinygo.org/x/drivers/pixel"
	"tinygo.org/x/tinyfont/freesans"
)

var flagUpdate = flag.Bool("update", false, "update test outputs")

type testCase struct {
	name string
	test func(t *testing.T, name string, img ImageDisplay)
}

func TestScreen(t *testing.T) {
	for _, tc := range []testCase{
		{"hello-banner", testHelloBanner},
		{"flex-grow", testGrowable},
	} {
		t.Run(tc.name, func(t *testing.T) {
			img := NewImageDisplay(160, 128) // same size as ST7735

			// Run the test.
			tc.test(t, tc.name, img)
		})
	}
}

func makeScreen(img ImageDisplay) *tinygl.Screen[pixel.RGB888] {
	buf := pixel.NewImage[pixel.RGB888](160, 8)
	screen := tinygl.NewScreen[pixel.RGB888](img, buf, 120) // 120PPI
	return screen
}

func testImage(t *testing.T, name string, img ImageDisplay) {
	path := filepath.Join("testdata", name+".png")

	golden := loadImage(t, path)
	if x, y, equal := isSameImage(img.image, golden); !equal {
		if *flagUpdate {
			// -update flag was passed, so update the PNG file.
			f, err := os.Create(path)
			if err != nil {
				t.Fatal("could not open test output:", err)
			}
			defer f.Close()
			png.Encode(f, img.image)
		} else {
			t.Errorf("image %s is not the same (mismatch at x=%d, y=%d)", path, x, y)
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

func testHelloBanner(t *testing.T, name string, img ImageDisplay) {
	screen := makeScreen(img)
	font := &freesans.Regular12pt7b
	foreground := pixel.NewRGB888(0xff, 0xff, 0xff)
	background := pixel.NewRGB888(64, 64, 64)
	topbar := tinygl.NewText(font, foreground, pixel.NewRGB888(255, 0, 0), "Hello world!")
	timelabel := tinygl.NewText(font, foreground, pixel.NewRGB888(0, 0, 0), "00:00")
	all := tinygl.NewVBox[pixel.RGB888](background, topbar, timelabel)
	screen.SetChild(all)
	screen.Update()
	testImage(t, name+"-0", img)

	timelabel.SetText("00:00:00")
	screen.Update()
	testImage(t, name+"-1", img)
}

func testGrowable(t *testing.T, name string, img ImageDisplay) {
	screen := makeScreen(img)
	font := &freesans.Regular12pt7b
	foreground := pixel.NewRGB888(0xff, 0xff, 0xff)
	background := pixel.NewRGB888(64, 64, 64)
	topbar := tinygl.NewText(font, foreground, pixel.NewRGB888(255, 0, 0), "Grow objects")
	line1 := tinygl.NewText(font, foreground, pixel.NewRGB888(0, 0, 0), "line 1")
	line2 := tinygl.NewText(font, foreground, pixel.NewRGB888(0, 0, 255), "line 2")
	all := tinygl.NewVBox[pixel.RGB888](background, topbar, line1, line2)
	screen.SetChild(all)

	// Test with a single element that is growable (it takes up all remaining
	// space).
	line1.SetGrowable(0, 1)
	screen.Update()
	testImage(t, name+"-0", img)

	// Add a second element that's growable. This time, it has a factor of 2, so
	// it will get two thirds of the slack pixels.
	line2.SetGrowable(0, 2)
	screen.Update()
	testImage(t, name+"-1", img)
}

type ImageDisplay struct {
	image *image.NRGBA
}

func NewImageDisplay(width, height int) ImageDisplay {
	// Create a new image, initialized with random data (to make non-drawn areas
	// very obvious).
	img := ImageDisplay{
		image: image.NewNRGBA(image.Rect(0, 0, width, height)),
	}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r := rand.Uint32()
			c := color.RGBA{
				R: uint8(r >> 0),
				G: uint8(r >> 8),
				B: uint8(r >> 16),
			}
			img.image.Set(x, y, c)
		}
	}
	return img
}

func (img ImageDisplay) Size() (int16, int16) {
	rect := img.image.Rect.Size()
	return int16(rect.X), int16(rect.Y)
}

func (img ImageDisplay) DrawBitmap(startX, startY int16, image pixel.Image[pixel.RGB888]) error {
	width, height := image.Size()
	for y := 0; y < int(height); y++ {
		for x := 0; x < int(width); x++ {
			inCol := image.Get(x, y)
			c := color.NRGBA{
				R: inCol.R,
				G: inCol.G,
				B: inCol.B,
				A: 0xff,
			}
			img.image.SetNRGBA(x+int(startX), y+int(startY), c)
		}
	}
	return nil
}

func (img ImageDisplay) Rotation() drivers.Rotation {
	return drivers.Rotation0
}

func (img ImageDisplay) Display() error {
	return nil // no-op
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
