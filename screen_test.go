package tinygl_test

import (
	"flag"
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
	"testing"

	"github.com/aykevl/tinygl"
	"github.com/aykevl/tinygl/pixel"
	"github.com/aykevl/tinygl/style"
	"tinygo.org/x/tinyfont/freesans"
)

var flagUpdate = flag.Bool("update", false, "update test outputs")

type testCase struct {
	name string
	draw func(*tinygl.Screen[pixel.RGB888], style.Style[pixel.RGB888])
}

func TestScreen(t *testing.T) {
	for _, tc := range []testCase{
		{"hello-banner", testHelloBanner},
	} {
		t.Run(tc.name, func(t *testing.T) {
			img := NewImageDisplay(160, 128) // same size as ST7735

			// Update the screen.
			font := &freesans.Regular12pt7b
			foreground := pixel.NewRGB888(0xff, 0xff, 0xff)
			background := pixel.NewRGB888(64, 64, 64)
			base := style.New(100, foreground, background, font)
			buf := make([]pixel.RGB888, 160*8)
			screen := tinygl.NewScreen(img, base, buf)
			tc.draw(screen, base)

			path := filepath.Join("testdata", tc.name+".png")

			if *flagUpdate {
				f, err := os.Create(path)
				if err != nil {
					t.Fatal("could not open test output:", err)
				}
				defer f.Close()
				png.Encode(f, img.image)
				return
			}

			f, err := os.Open(path)
			if err != nil {
				t.Fatal("could not open test output:", err)
			}
			defer f.Close()
			golden, err := png.Decode(f)
			if err != nil {
				t.Fatal("could not decode test output:", err)
			}
			if x, y, equal := isSameImage(img.image, golden); !equal {
				t.Errorf("image %s is not the same (mismatch at x=%d, y=%d)", path, x, y)
			}
		})
	}
}

func testHelloBanner(screen *tinygl.Screen[pixel.RGB888], base style.Style[pixel.RGB888]) {
	topbar := tinygl.NewText(base.WithBackground(color.RGBA{R: 255, A: 255}), "Hello world!")
	timelabel := tinygl.NewText(base.WithBackground(color.RGBA{A: 255}), "00:00")
	all := tinygl.NewVBox[pixel.RGB888](base, topbar, timelabel)
	screen.SetChild(all)
	screen.Update()
	timelabel.SetText("00:00:00")
	screen.Update()
}

type ImageDisplay struct {
	image *image.NRGBA
}

func NewImageDisplay(width, height int) ImageDisplay {
	return ImageDisplay{
		image: image.NewNRGBA(image.Rect(0, 0, width, height)),
	}
}

func (img ImageDisplay) Size() (int16, int16) {
	rect := img.image.Rect.Size()
	return int16(rect.X), int16(rect.Y)
}

func (img ImageDisplay) DrawRGBBitmap8(startX, startY int16, data []byte, width, height int16) error {
	for y := 0; y < int(height); y++ {
		for x := 0; x < int(width); x++ {
			offset := (y*int(width) + x) * 3
			c := color.NRGBA{
				R: data[offset+0],
				G: data[offset+1],
				B: data[offset+2],
				A: 0xff,
			}
			img.image.SetNRGBA(x+int(startX), y+int(startY), c)
		}
	}
	return nil
}

func (img ImageDisplay) Display() error {
	return nil // no-op
}

func isSameImage(image1, image2 image.Image) (x, y int, equal bool) {
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
