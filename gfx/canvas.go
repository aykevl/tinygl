package gfx

import (
	"image/color"

	"github.com/aykevl/tinygl"
	"github.com/aykevl/tinygl/pixel"
	"github.com/aykevl/tinygl/style"
)

const blockSize = 16 // Note: this value may need some tuning.

// A canvas based on small tiles that supports modifying these objects and only
// redraws the affected tiles. This results in very fast incremental updates if
// only small parts of the screen changed.
type Canvas[T pixel.Color] struct {
	tinygl.Rect[T]
	blocksWidth  int16
	blocksHeight int16
	dirty        []byte
	objects      []Object[T]
}

// NewCanvas creates a new tile canvas.
func NewCanvas[T pixel.Color](base style.Style[T], width, height int) *Canvas[T] {
	c := &Canvas[T]{
		Rect: tinygl.MakeRect(base, width, height),
	}
	return c
}

// Layout implements tinygl.Object.
func (c *Canvas[T]) Layout(x, y, width, height int) {
	displayX, displayY, displayWidth, displayHeight := c.Rect.Bounds()
	if x != displayX || y != displayY || width != displayWidth || height != displayHeight {
		c.Rect.Layout(x, y, width, height)
		numBlocksWidth := (width + blockSize - 1) / blockSize
		numBlocksHeight := (height + blockSize - 1) / blockSize
		if numBlocksWidth != int(c.blocksWidth) || numBlocksHeight != int(c.blocksHeight) {
			c.blocksWidth = int16(numBlocksWidth)
			c.blocksHeight = int16(numBlocksHeight)
			numBlocks := numBlocksWidth * numBlocksHeight
			dirtyBytes := (numBlocks + 8 - 1) / 8
			// Note: for canvases that frequently change size, it might be
			// worthwhile to grow/shrink the slice as needed.
			c.dirty = make([]byte, dirtyBytes)
		}
		// Mark all blocks on the canvas as needing an update.
		for i := range c.dirty {
			c.dirty[i] = 0xff
		}
	}
}

// Update implements tinygl.Object.
func (c *Canvas[T]) Update(screen *tinygl.Screen[T]) {
	if !c.NeedsUpdate() { // check the needsUpdate flag
		return
	}
	// needsUpdate flag is cleared

	// Go through all the tiles and update those that changed.
	x, y, _, _ := c.Bounds()
	buf := screen.Buffer()[:blockSize*blockSize]
	img := pixel.NewImageFromBuffer(buf, blockSize)
	for blockY := 0; blockY < int(c.blocksHeight); blockY++ {
		for blockX := 0; blockX < int(c.blocksWidth); blockX++ {
			// Note: could be sped up by checking whole bytes at a time.
			if !c.isDirty(blockX, blockY) {
				continue
			}
			// TODO: combine blocks into a larger rectangle to be drawn at a
			// single time.
			c.clearDirty(blockX, blockY)

			// Paint block and send.
			background := c.Background()
			for i := range buf {
				buf[i] = background
			}
			for _, obj := range c.objects {
				obj.Draw(blockX*blockSize, blockY*blockSize, img)
			}
			screen.Send(buf, x+blockX*blockSize, y+blockY*blockSize, blockSize, blockSize)
		}
	}
}

func (c *Canvas[T]) Draw(x, y int, img pixel.Image[T]) {
	panic("todo: Canvas.Draw")
}

func (c *Canvas[T]) isDirty(blockX, blockY int) bool {
	blockNum := blockY*int(c.blocksWidth) + blockX
	return c.dirty[blockNum/8]&(1<<(blockNum%8)) != 0
}

func (c *Canvas[T]) clearDirty(blockX, blockY int) {
	blockNum := blockY*int(c.blocksWidth) + blockX
	c.dirty[blockNum/8] &^= (1 << (blockNum % 8))
}

func (c *Canvas[T]) setDirty(blockX, blockY int) {
	blockNum := blockY*int(c.blocksWidth) + blockX
	c.dirty[blockNum/8] |= (1 << (blockNum % 8))
}

// markDirty marks all tiles in the given rectangle as dirty.
func (c *Canvas[T]) markDirty(x, y, width, height int) {
	startX := x / blockSize
	startY := y / blockSize
	stopX := (x + width + blockSize - 1) / blockSize
	stopY := (y + height + blockSize - 1) / blockSize
	if stopX < 0 || stopY < 0 || startX >= int(c.blocksWidth) || startY >= int(c.blocksHeight) {
		return // out of range
	}
	if startX < 0 {
		startX = 0
	}
	if startY < 0 {
		startY = 0
	}
	if stopX >= int(c.blocksWidth) {
		stopX = int(c.blocksWidth)
	}
	if stopY >= int(c.blocksHeight) {
		stopY = int(c.blocksHeight)
	}

	for blockY := startY; blockY < stopY; blockY++ {
		for blockX := startX; blockX < stopX; blockX++ {
			c.setDirty(blockX, blockY)
		}
	}
	c.RequestUpdate()
}

// Clear removes all objects from the canvas.
func (c *Canvas[T]) Clear() {
	for _, obj := range c.objects {
		// Hide (if not already hidden), so that the area will be repainted.
		obj.SetHidden(true)
	}
	c.objects = c.objects[:0]
}

// CreateRect creates a new rectangle at the given coordinates with the given
// color.
func (c *Canvas[T]) CreateRect(x, y, width, height int, color color.RGBA) *Rect[T] {
	obj := &Rect[T]{
		canvas: c,
		x:      int16(x),
		y:      int16(y),
		width:  int16(width),
		height: int16(height),
		color:  pixel.NewColor[T](color.R, color.G, color.B),
	}
	obj.markDirty()
	c.objects = append(c.objects, obj)
	return obj
}
