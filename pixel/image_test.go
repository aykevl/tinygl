package pixel_test

import (
	"image/color"
	"testing"

	"github.com/aykevl/tinygl/pixel"
)

func TestImageRGB565BE(t *testing.T) {
	image := pixel.NewImage[pixel.RGB565BE](5, 3)
	if width, height := image.Size(); width != 5 && height != 3 {
		t.Errorf("image.Size(): expected 5, 3 but got %d, %d", width, height)
	}
	for _, c := range []color.RGBA{
		{R: 0xff, A: 0xff},
		{G: 0xff, A: 0xff},
		{B: 0xff, A: 0xff},
		{R: 0x10, A: 0xff},
		{G: 0x10, A: 0xff},
		{B: 0x10, A: 0xff},
	} {
		image.Set(4, 2, image.Color(c.R, c.G, c.B))
		c2 := image.Get(4, 2).RGBA()
		if c2 != c {
			t.Errorf("failed to roundtrip color: expected %v but got %v", c, c2)
		}
	}
}
