package main

// This is the code that generates font files for the font renderer.
// It outputs a heavily documented Go file that contain the binary data of the
// font file as a string (because strings are read-only in Go and can be stored
// in flash).
//
// Right now the format only supports a simple two bits per pixel bitmap, which
// provides just enough antialiasing to be a whole lot better than no
// antialiasing but without the flash storage demands of something like 4 bits
// per pixels (which does look better).
//
// We use 72dpi, because that's the DPI at which a given size in px maps to the
// same size in pt. So a size 16px (as used on the web) renders at the same size
// on a monitor set to 100% scaling.
//
// For more information about the format, see the README.
//
// Terminology:
// - rune: Unicode code point
// - glyph: actual shape of a character (rendered as bitmap)

import (
	"bytes"
	"flag"
	"fmt"
	"image"
	"image/color"
	"math"
	"os"
	"text/template"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"
)

var (
	flagFont    = flag.String("font", "", "Font file (.ttf or .otf)")
	flagSize    = flag.Float64("size", 16, "render at given size")
	flagDPI     = flag.Float64("dpi", 72, "render at given DPI")
	flagOut     = flag.String("o", "", "output filename")
	flagPackage = flag.String("package", "font", "package name for output Go file")
	flagVerbose = flag.Bool("verbose", false, "print debugging information")
)

// Store calculated font data before writing it out as a .go file.
type FontData struct {
	size         uint8
	height       uint8
	ascent       uint8
	runeTables   []*runeTable
	glyphs       []*glyph
	glyphIndices map[string]int

	// Keep track of the glyph bounding boxes, for verbose mode.
	glyphAdvance rangeTrack
	glyphTop     rangeTrack
	glyphBottom  rangeTrack
	glyphLeft    rangeTrack
	glyphRight   rangeTrack

	// Compression statistics.
	totalGlyphRows        int
	glyphRowsRepeats      int
	glyphRowsRepeatPixels int
}

// Write font data out as a Go file, don't provide any font mathematics.
func (d *FontData) writeFile(filename string) error {
	// Calculate glyph offsets.
	offset := 4 // font header
	for _, table := range d.runeTables {
		offset += 3 + 2                 // number of runes + first rune
		offset += len(table.glyphs) * 2 // glyphs
	}
	offset += 2 // rune table end marker
	for _, glyph := range d.glyphs {
		glyph.offset = uint32(offset)
		offset += len(glyph.data)
	}

	// Write beginning of the file.
	w := &bytes.Buffer{}
	template.Must(template.New("").Parse(
		`// File generated using:
// 	go run ./generate -font={{.font}} -size={{.flagSize}} -dpi={{.flagDPI}} -package={{.package}}

package {{.package}}

import (
	"github.com/aykevl/tinygl/font"
)

// This font takes up {{.binarySize}} bytes.

var Regular{{.size}} = font.Make("" +
	"\x00" + // version: 0
	"\x{{printf "%02x" .size  }}" + // size:   {{.size}}
	"\x{{printf "%02x" .height}}" + // height: {{.height}}
	"\x{{printf "%02x" .ascent}}" + // ascent: {{.ascent}}
`)).Execute(w, map[string]any{
		"font":       *flagFont,
		"flagSize":   *flagSize,
		"flagDPI":    *flagDPI,
		"package":    *flagPackage,
		"size":       d.size,
		"height":     d.height,
		"ascent":     d.ascent,
		"binarySize": offset,
	})

	// Add rune tables.
	for _, table := range d.runeTables {
		fmt.Fprintf(w, "\n\t// Runes %d..%d (%d runes)\n", table.start, int(table.start)+len(table.glyphs)-1, len(table.glyphs))
		fmt.Fprintf(w, "\t%s + // number of runes (%d)\n", encode16(uint32(len(table.glyphs))), len(table.glyphs))
		fmt.Fprintf(w, "\t%s + // start rune (%d)\n", encode24(uint32(table.start)), table.start)
		for i, index := range table.glyphs {
			glyph := d.glyphs[index]
			r := table.start + rune(i)
			fmt.Fprintf(w, "\t%s + // %#v at index %d\n", encode16(glyph.offset), string(r), glyph.offset)
		}
	}
	fmt.Fprintf(w, "\n\t// mark the end of the rune tables\n")
	fmt.Fprintf(w, "\t%#v +\n", "\x00\x00")

	// Add glyphs.
	for _, glyph := range d.glyphs {
		if len(glyph.runes) == 1 {
			fmt.Fprintf(w, "\n\t// glyph %d for rune %q\n", glyph.offset, glyph.runes[0])
		} else {
			fmt.Fprintf(w, "\n\t// glyph %d for runes %q", glyph.offset, glyph.runes[0])
			for _, r := range glyph.runes[1:] {
				fmt.Fprintf(w, ", %q", r)
			}
			fmt.Fprintf(w, "\n")
		}
		fmt.Fprintf(w, "\t\"")
		for _, b := range []byte(glyph.data) {
			fmt.Fprintf(w, "\\x%02x", b)
		}
		fmt.Fprintf(w, "\" +\n")
	}

	// Write end of file.
	w.WriteString("\n\t\"\")\n")

	// Write the file out to disk.
	return os.WriteFile(filename, w.Bytes(), 0o666)
}

// Add a range of runes, from start until end (inclusive).
// Glyphs are deduplicated if there are duplicates (rare, but it might happen).
//
// TODO: while runes generally map to a single glyph in latin1 characters, this
// is not true for many other languages. Right now we assume this mapping but if
// we want to support languages like arabic this simple mapping won't do.
func (d *FontData) addRunes(face font.Face, start, end rune) {
	table := &runeTable{
		start: start,
	}
	d.runeTables = append(d.runeTables, table)
	for r := start; r <= end; r++ {
		// Convert rune to glyph.
		glyphData, ok := d.makeGlyph(face, r)
		if !ok {
			table.glyphs = append(table.glyphs, 0) // replacement character
			d.glyphs[0].runes = append(d.glyphs[0].runes, r)
			continue
		}

		// Deduplicate glyph.
		if index, ok := d.glyphIndices[string(glyphData)]; ok {
			table.glyphs = append(table.glyphs, index)
			d.glyphs[0].runes = append(d.glyphs[0].runes, r)
			continue
		}

		// Add new glyph.
		index := len(d.glyphs)
		d.glyphs = append(d.glyphs, &glyph{
			data:  glyphData,
			runes: []rune{r},
		})
		d.glyphIndices[string(glyphData)] = index
		table.glyphs = append(table.glyphs, index)
	}
}

// Convert a single rune to a glyph.
func (fd *FontData) makeGlyph(face font.Face, r rune) (data []byte, ok bool) {
	advance, ok := face.GlyphAdvance(r)
	if !ok {
		// Glyph not found.
		return nil, false
	}

	// Create image with a single glyph.
	const width, height = 200, 200
	img := image.NewGray(image.Rect(0, 0, width, height))
	d := font.Drawer{
		Dst:  img,
		Src:  image.White,
		Face: face,
		Dot:  fixed.P(width/2, height/2),
	}
	d.DrawString(string(r))

	// We use two bits per pixel.
	const bits = 2
	grayBits := func(gray color.Gray) uint8 {
		return gray.Y >> (8 - bits)
	}

	// Find the bounding box of this glyph.
	top := height
	bottom := -1
	left := width
	right := -1
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			c := grayBits(img.GrayAt(x, y))
			if c != 0 {
				if y < top {
					top = y
				}
				if y >= bottom {
					bottom = y + 1
				}
				if x < left {
					left = x
				}
				if x >= right {
					right = x + 1
				}
			}
		}
	}
	top -= height / 2
	bottom -= height / 2
	left -= height / 2
	right -= height / 2
	if top > bottom {
		// no bitmap (e.g. space)
		top = 0
		bottom = 0
		left = 0
		right = 0
	}

	// Keep track of glyph bounding box.
	fd.glyphAdvance.Add(advance.Round())
	fd.glyphTop.Add(top)
	fd.glyphBottom.Add(bottom)
	fd.glyphLeft.Add(left)
	fd.glyphRight.Add(right)

	// Convert grayscale image to a bitmap.
	maskWidth := right - left
	maskHeight := bottom - top
	var mask []byte
	index := 0
	addBits := func(n uint8) {
		if index/8 >= len(mask) {
			mask = append(mask, 0)
		}
		mask[index/8] |= n << (index % 8)
		index += bits
	}
	for y := 0; y < maskHeight; y++ {
		fd.totalGlyphRows++
		imgY := y + top + height/2
		if y != 0 {
			// Check whether it is identical to the previous row.
			sameAsPrevious := true
			for x := 0; x < maskWidth; x++ {
				imgX := x + left + width/2
				c1 := grayBits(img.GrayAt(imgX, imgY-1))
				c2 := grayBits(img.GrayAt(imgX, imgY))
				if c1 != c2 {
					sameAsPrevious = false
				}
			}
			if sameAsPrevious {
				fd.glyphRowsRepeats++
				fd.glyphRowsRepeatPixels += maskWidth

				// Emit "repeat previous line" command.
				// TODO: it might be profitable to also specify how often to
				// repeat.
				addBits(0b01)
				continue
			}
		}

		// Emit bitmap command.
		addBits(0b00)
		for x := 0; x < maskWidth; x++ {
			imgX := x + left + width/2
			c := grayBits(img.GrayAt(imgX, imgY))
			addBits(c)
		}
	}

	// Make the glyph data.
	// TODO: it's possible to shrink these coordinates somewhat: they currently
	// take up more space than needed.
	data = []byte{
		byte(advance.Round()),
		byte(int8(top)),
		byte(int8(bottom)),
		byte(int8(left)),
		byte(int8(right)),
	}
	data = append(data, mask...)

	return data, true
}

// Range of runes and their glyph index.
type runeTable struct {
	start  rune
	glyphs []int // indices into glyphs table
}

// Single glyph bitmap.
type glyph struct {
	data   []byte // binary data (metadata and bitmap)
	offset uint32 // offset into the file
	runes  []rune // a single glyph might be generated from multiple runes
}

// Convert a font file based on the command line parameters.
func convert() error {
	// Load the font data.
	data, err := os.ReadFile(*flagFont)
	if err != nil {
		return fmt.Errorf("failed to read font file: %w", err)
	}
	parsedFont, err := opentype.Parse(data)
	if err != nil {
		return fmt.Errorf("failed to parse font file: %w", err)
	}
	face, err := opentype.NewFace(parsedFont, &opentype.FaceOptions{
		Size:    *flagSize,
		DPI:     *flagDPI,
		Hinting: font.HintingFull,
	})
	if err != nil {
		return fmt.Errorf("failed to create font face at size=%v and dpi=%v: %w", *flagSize, *flagDPI, err)
	}

	// Create font struct.
	metrics := face.Metrics()
	fontData := FontData{
		size:         uint8(math.Round(*flagSize)),
		height:       uint8(metrics.Height.Round()),
		ascent:       uint8(metrics.Ascent.Round()),
		glyphIndices: make(map[string]int),
	}

	// Load glyphs one by one.
	// TODO: include replacement character as glyph 0.
	fontData.addRunes(face, 32, 126) // code points 32..126

	// Write the font as a Go file.
	if *flagOut != "" {
		err := fontData.writeFile(*flagOut)
		if err != nil {
			return err
		}
	}

	if *flagVerbose {
		fmt.Printf("Glyph metadata ranges:\n")
		fmt.Printf("  advance: %s\n", fontData.glyphAdvance)
		fmt.Printf("  top:     %s\n", fontData.glyphTop)
		fmt.Printf("  bottom:  %s\n", fontData.glyphBottom)
		fmt.Printf("  left:    %s\n", fontData.glyphLeft)
		fmt.Printf("  right:   %s\n", fontData.glyphRight)
		fmt.Printf("Compression stats:\n")
		fmt.Printf("  Total number of rows: %d\n", fontData.totalGlyphRows)
		fmt.Printf("  Rows that repeat:     %d (pixels: %d)\n", fontData.glyphRowsRepeats, fontData.glyphRowsRepeatPixels)
	}

	return nil
}

func encode16(n uint32) string {
	if n&0xff_ff != n {
		panic("cannot encode as 16 bit number") // sanity check
	}
	return fmt.Sprintf("\"\\x%02x\\x%02x\"", uint8(n>>0), uint8(n>>8))
}

func encode24(n uint32) string {
	if n&0xff_ff_ff != n {
		panic("cannot encode as 24 bit number") // sanity check
	}
	return fmt.Sprintf("\"\\x%02x\\x%02x\\x%02x\"", uint8(n>>0), uint8(n>>8), uint8(n>>16))
}

type rangeTrack struct {
	min      int
	max      int
	hasValue bool
}

func (r *rangeTrack) Add(value int) {
	if !r.hasValue {
		// initial value
		r.hasValue = true
		r.min = value
		r.max = value
		return
	}
	if value < r.min {
		r.min = value
	}
	if value > r.max {
		r.max = value
	}
}

func (r rangeTrack) String() string {
	if !r.hasValue {
		return "<unknown>"
	}
	return fmt.Sprintf("%d..%d", r.min, r.max)
}

func main() {
	flag.Parse()
	err := convert()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
