package gfx

import (
	"github.com/aykevl/tinygl"
	"github.com/aykevl/tinygl/pixel"
)

const blockSize = 16 // Note: this value may need some tuning.

// A canvas based on small tiles that supports modifying these objects and only
// redraws the affected tiles. This results in very fast incremental updates if
// only small parts of the screen changed.
type Canvas[T pixel.Color] struct {
	tinygl.Rect[T]
	minWidth     int16
	minHeight    int16
	canvasWidth  int16
	canvasHeight int16
	dirty        []byte
	objects      []Object[T]
	eventHandler func(tinygl.Event, int, int)
}

// NewCanvas creates a new tile canvas.
func NewCanvas[T pixel.Color](background T, width, height int) *Canvas[T] {
	c := &Canvas[T]{
		Rect:      tinygl.MakeRect(background),
		minWidth:  int16(width),
		minHeight: int16(height),
	}
	return c
}

// MinSize returns the minimal size as set when the canvas was created.
func (c *Canvas[T]) MinSize() (width, height int) {
	return int(c.minWidth), int(c.minHeight)
}

// Size returns the current canvas size, as updated after the first layout.
func (c *Canvas[T]) Size() (width, height int) {
	return int(c.canvasWidth), int(c.canvasHeight)
}

// Layout implements tinygl.Object.
func (c *Canvas[T]) Layout(width, height int) {
	numBlocksWidth := (width + blockSize - 1) / blockSize
	numBlocksHeight := (height + blockSize - 1) / blockSize
	if width != int(c.canvasWidth) || height != int(c.canvasHeight) {
		c.canvasWidth = int16(width)
		c.canvasHeight = int16(height)
		numBlocks := numBlocksWidth * numBlocksHeight
		dirtyBytes := (numBlocks + 8 - 1) / 8
		// Note: for canvases that frequently change size, it might be
		// worthwhile to grow/shrink the slice as needed.
		c.dirty = make([]byte, dirtyBytes)
		// Mark all blocks on the canvas as needing an update.
		for i := range c.dirty {
			c.dirty[i] = 0xff
		}
	}
}

// Update implements tinygl.Object.
func (c *Canvas[T]) Update(screen *tinygl.Screen[T], displayX, displayY, displayWidth, displayHeight, x, y int) {
	if this, _ := c.NeedsUpdate(); !this { // check the needsUpdate flag
		return
	}
	// needsUpdate flag is cleared

	// Go through all the tiles and update those that changed.
	drawX0 := x
	drawY0 := y
	drawX1 := x + displayWidth
	drawY1 := y + displayHeight
	blockX0 := (drawX0 + blockSize - 1) / blockSize
	blockY0 := (drawY0 + blockSize - 1) / blockSize
	blockX1 := (drawX1 + blockSize - 1) / blockSize
	blockY1 := (drawY1 + blockSize - 1) / blockSize
	for blockY := blockY0; blockY < blockY1; blockY++ {
		// TODO: don't just limit the block height, also respect the other 3
		// sides.
		blockHeight := blockSize
		blockMaxHeight := drawY1 - blockY*blockSize
		if blockHeight > blockMaxHeight {
			blockHeight = blockMaxHeight
		}
		if blockHeight <= 0 {
			panic("unreachable?")
		}
		buf := screen.Buffer()
		maxBlocksPerRow := buf.Len() / (blockSize * blockHeight)
		for blockX := blockX0; blockX < blockX1; blockX++ {
			// Note: could be sped up by checking whole bytes at a time.
			if !c.isDirty(blockX, blockY) {
				continue
			}
			// TODO: combine blocks into a larger rectangle to be drawn at a
			// single time.
			c.clearDirty(blockX, blockY)

			// Find other blocks on the same row that could be painted at the
			// same time.
			numBlocks := 1
			for {
				if numBlocks >= maxBlocksPerRow || blockX+numBlocks >= blockX1 {
					break
				}
				if !c.isDirty(blockX+numBlocks, blockY) {
					break
				}
				c.clearDirty(blockX+numBlocks, blockY)
				numBlocks++
			}

			// Paint block and send.
			img := buf.Rescale(blockSize*numBlocks, blockHeight)
			img.FillSolidColor(c.Background())
			for _, obj := range c.objects {
				obj.Draw(blockX*blockSize, blockY*blockSize, img)
			}
			screen.Send(displayX-x+blockX*blockSize, displayY-y+blockY*blockSize, img)
		}
	}
}

// SetEventHandler sets the callback when a (touch) event occurs on this canvas.
func (c *Canvas[T]) SetEventHandler(eventHandler func(event tinygl.Event, x, y int)) {
	c.eventHandler = eventHandler
}

// HandleEvent handles events such as touch events and calls the event handler
// with the x/y coordinate as parameters.
func (c *Canvas[T]) HandleEvent(event tinygl.Event, x, y int) {
	if c.eventHandler == nil {
		return
	}
	c.eventHandler(event, x, y)
}

func (c *Canvas[T]) isDirty(blockX, blockY int) bool {
	blockNum := blockY*((int(c.canvasWidth)+blockSize-1)/blockSize) + blockX
	return c.dirty[blockNum/8]&(1<<(blockNum%8)) != 0
}

func (c *Canvas[T]) clearDirty(blockX, blockY int) {
	blockNum := blockY*((int(c.canvasWidth)+blockSize-1)/blockSize) + blockX
	c.dirty[blockNum/8] &^= (1 << (blockNum % 8))
}

func (c *Canvas[T]) setDirty(blockX, blockY int) {
	blockNum := blockY*((int(c.canvasWidth)+blockSize-1)/blockSize) + blockX
	c.dirty[blockNum/8] |= (1 << (blockNum % 8))
}

// markDirty marks all tiles in the given rectangle as dirty.
func (c *Canvas[T]) markDirty(x, y, width, height int) {
	if x < 0 {
		width -= x
		x = 0
	}
	if y < 0 {
		height -= y
		y = 0
	}
	if x+width > int(c.canvasWidth) {
		width = int(c.canvasWidth) - x
	}
	if y+height > int(c.canvasHeight) {
		height = int(c.canvasHeight) - y
	}
	if width == 0 || height == 0 {
		return
	}
	startX := x / blockSize
	startY := y / blockSize
	stopX := (x + width + blockSize - 1) / blockSize
	stopY := (y + height + blockSize - 1) / blockSize
	if startX < 0 {
		startX = 0
	}
	if startY < 0 {
		startY = 0
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
func (c *Canvas[T]) CreateRect(x, y, width, height int, color T) *Rect[T] {
	obj := &Rect[T]{
		canvas: c,
		x:      int16(x),
		y:      int16(y),
		width:  int16(width),
		height: int16(height),
		color:  color,
	}
	obj.markDirty()
	c.objects = append(c.objects, obj)
	return obj
}
