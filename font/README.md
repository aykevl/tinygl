# Fonts

This package contains a number of fonts and code for parsing/rendering them.

Currently supported:

  * Rendering on a `pixel.Image[T]` buffer for fast drawing.
  * Two bits per pixel, which provides very bare bones antialiasing. This balances code size and appearance: it looks a lot better than no antialiasing at all while not increasing size as much as 4 bits per pixel (which does look a lot better though).  
    I'm betting on screens only getting higher resolution in the future so that antialiasing increasinly has less impact on appearance.

To be implemented:

  * Better compression. I have a few ideas how fonts can be compressed while not increasing rendering time too much.
  * Kerning. I also have a few ideas for this one, but I'll have to see how welll they work.
  * Languages where there isn't a simple mapping between Unicode code point (rune) and glyph, like Arabic.
  * Right-to-left (or other direction) languages.
  * Colored emojis.

## Format

The font file format consists of a few parts:

  * The header, with some very basic metadata.
  * The rune tables, that describe for each rune (Unicode code point) where in the file the glyph can be found.
  * A number of glyphs.

All values are little endian.

### Header

The header consists of the following 4 bytes:

| type  | name    | value
| ----- | ------- | -----
| uint8 | version | Font version, to be incremented with backwards incompatible changes. Currently 0.
| uint8 | size    | Original font size (in pixels, or points at 72dpi).
| uint8 | height  | Recommended line height of the original font. Usually slightly larger than the size.
| uint8 | ascent  | Font ascent value, that is, how far down (from the top) the origin of each glyph is. When drawing text in an area of the font height, this value should be used as the Y value to start drawing.

### Rune tables

The rune tables maps runes to glyphs, as offsets in the file. Each table contains a start rune, the number of runes in the table, and an array of file offsets.

Specifically, a table has the following format:

| type      | name    | value
| --------- | ------- | -----
| uint16    | length  | The number of runes in this table.
| uint24    | start   | The first rune in the table.
| [N]uint16 | offsets | The glyph offsets from the start of the file, corresponding to a particular rune.

To mark the last table, two 0 bytes are inserted after the last rune table which are decoded as a table of length 0.

### Glyphs

A glyph has some metadata followed by the bitmap for that glyph.

The header looks like this:

| type      | name    | value
| --------- | ------- | -----
| uint8     | advance | The "width" of the character, that is, how much to increase the X value to draw the next glyph. This does not take kerning into account.
| int8      | top     | Top of the bitmap, relative to the origin (usually negative).
| int8      | bottom  | Bottom of the bitmap, relative to the origin.
| int8      | left    | Left bounding box of the bitmap, relative to the origin. This can be a negative integer for characters like 'j'.
| int8      | right   | Right bounding box for the bitmap, relative to the origin (usually positive).

The bitmap consists of a number of rows, prefixed by a two-bit command. All values are packed as a large bitstring, though care is take to keep commands in a single byte.

  * Command `0b00` simply draws a bitmap row. The data consists of a sequence of bits, where each pair of two bits represents one pixel.
  * Command `0b01` repeats the previous command. This is beneficial in many characters such as 'n' which have many rows that are actually just a repeat of a previous row.

The bitmap is padded to a full byte using zero bits.
