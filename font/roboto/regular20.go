// File generated using:
// 	go run ./generate -font=roboto/Roboto-Regular.ttf -size=20 -dpi=72 -package=roboto

package roboto

import (
	"github.com/aykevl/tinygl/font"
)

// Font statistics:
// - total size:      3196
// - glyph metadata:  475
// - glyph mask data: 2520

var Regular20 = font.Make("" +
	"\x00" + // version: 0
	"\x14" + // size:   20
	"\x18" + // height: 24
	"\x13" + // ascent: 19

	// Runes 32..126 (95 runes)
	"\x5f\x00" + // number of runes (95)
	"\x20\x00\x00" + // start rune (32)
	"\xc9\x00" + // " " at index 201
	"\xce\x00" + // "!" at index 206
	"\xdb\x00" + // "\"" at index 219
	"\xe5\x00" + // "#" at index 229
	"\x14\x01" + // "$" at index 276
	"\x44\x01" + // "%" at index 324
	"\x7e\x01" + // "&" at index 382
	"\xb0\x01" + // "'" at index 432
	"\xb8\x01" + // "(" at index 440
	"\xdc\x01" + // ")" at index 476
	"\x01\x02" + // "*" at index 513
	"\x16\x02" + // "+" at index 534
	"\x2d\x02" + // "," at index 557
	"\x37\x02" + // "-" at index 567
	"\x3f\x02" + // "." at index 575
	"\x46\x02" + // "/" at index 582
	"\x6d\x02" + // "0" at index 621
	"\x8d\x02" + // "1" at index 653
	"\x9c\x02" + // "2" at index 668
	"\xcb\x02" + // "3" at index 715
	"\xf6\x02" + // "4" at index 758
	"\x20\x03" + // "5" at index 800
	"\x47\x03" + // "6" at index 839
	"\x6d\x03" + // "7" at index 877
	"\x9c\x03" + // "8" at index 924
	"\xc7\x03" + // "9" at index 967
	"\xf0\x03" + // ":" at index 1008
	"\xfc\x03" + // ";" at index 1020
	"\x0b\x04" + // "<" at index 1035
	"\x27\x04" + // "=" at index 1063
	"\x39\x04" + // ">" at index 1081
	"\x55\x04" + // "?" at index 1109
	"\x78\x04" + // "@" at index 1144
	"\xce\x04" + // "A" at index 1230
	"\x04\x05" + // "B" at index 1284
	"\x2e\x05" + // "C" at index 1326
	"\x5b\x05" + // "D" at index 1371
	"\x82\x05" + // "E" at index 1410
	"\x9f\x05" + // "F" at index 1439
	"\xb7\x05" + // "G" at index 1463
	"\xe9\x05" + // "H" at index 1513
	"\xfe\x05" + // "I" at index 1534
	"\x07\x06" + // "J" at index 1543
	"\x1f\x06" + // "K" at index 1567
	"\x4e\x06" + // "L" at index 1614
	"\x5e\x06" + // "M" at index 1630
	"\x9b\x06" + // "N" at index 1691
	"\xce\x06" + // "O" at index 1742
	"\xfe\x06" + // "P" at index 1790
	"\x20\x07" + // "Q" at index 1824
	"\x5a\x07" + // "R" at index 1882
	"\x89\x07" + // "S" at index 1929
	"\xb8\x07" + // "T" at index 1976
	"\xca\x07" + // "U" at index 1994
	"\xe1\x07" + // "V" at index 2017
	"\x14\x08" + // "W" at index 2068
	"\x55\x08" + // "X" at index 2133
	"\x84\x08" + // "Y" at index 2180
	"\xa8\x08" + // "Z" at index 2216
	"\xd4\x08" + // "[" at index 2260
	"\xe3\x08" + // "\\" at index 2275
	"\x0a\x09" + // "]" at index 2314
	"\x1a\x09" + // "^" at index 2330
	"\x2d\x09" + // "_" at index 2349
	"\x37\x09" + // "`" at index 2359
	"\x40\x09" + // "a" at index 2368
	"\x61\x09" + // "b" at index 2401
	"\x86\x09" + // "c" at index 2438
	"\xa7\x09" + // "d" at index 2471
	"\xc2\x09" + // "e" at index 2498
	"\xe1\x09" + // "f" at index 2529
	"\xf5\x09" + // "g" at index 2549
	"\x17\x0a" + // "h" at index 2583
	"\x2e\x0a" + // "i" at index 2606
	"\x3c\x0a" + // "j" at index 2620
	"\x50\x0a" + // "k" at index 2640
	"\x74\x0a" + // "l" at index 2676
	"\x7e\x0a" + // "m" at index 2686
	"\x99\x0a" + // "n" at index 2713
	"\xac\x0a" + // "o" at index 2732
	"\xd0\x0a" + // "p" at index 2768
	"\xf7\x0a" + // "q" at index 2807
	"\x12\x0b" + // "r" at index 2834
	"\x22\x0b" + // "s" at index 2850
	"\x40\x0b" + // "t" at index 2880
	"\x52\x0b" + // "u" at index 2898
	"\x65\x0b" + // "v" at index 2917
	"\x86\x0b" + // "w" at index 2950
	"\xb7\x0b" + // "x" at index 2999
	"\xd8\x0b" + // "y" at index 3032
	"\x03\x0c" + // "z" at index 3075
	"\x21\x0c" + // "{" at index 3105
	"\x42\x0c" + // "|" at index 3138
	"\x4d\x0c" + // "}" at index 3149
	"\x6b\x0c" + // "~" at index 3179

	// mark the end of the rune tables
	"\x00\x00" +

	// glyph 201 for rune ' '
	"\x05\x00\x00\x00\x00" +

	// glyph 206 for rune '!'
	"\x05\xf2\x00\x01\x04\xb4\xd1\x55\x15\x1c\x40\x70\xb4" +

	// glyph 219 for rune '"'
	"\x06\xf1\xf6\x01\x06\x28\x17\xca\x84\x31" +

	// glyph 229 for rune '#'
	"\x0c\xf2\x00\x01\x0c\x00\x28\x28\x00\x2c\x2c\x00\x1c\x1c\x00\x1d\x1d\xf0\xff\xff\x00\x0e\x0e\x00\x0b\x0b\x00\x07\x07\x54\x57\x17\xfc\xff\x7f\x80\x83\x03\xc0\xc2\x02\xc0\xc1\x01\xc0\xc0\x00" +

	// glyph 276 for rune '$'
	"\x0b\xef\x02\x01\x0a\x00\x14\x00\xc0\x02\x01\xff\x0b\xbc\xf4\xc1\x02\x3c\x1d\xc0\xc3\x02\x00\xbc\x00\x00\xfe\x01\x40\xfe\x00\x00\x2e\x00\xc0\xf3\x00\x38\x1e\xc0\xc3\x9b\x1f\xf0\x7f\x00\x70\x40" +

	// glyph 324 for rune '%'
	"\x0f\xf1\x00\x01\x0e\x40\x01\x00\x00\xff\x00\x00\x2c\x2c\x20\xc0\x81\x83\x03\x1c\x38\x1c\x80\xc3\xb2\x00\xf0\x4b\x03\x00\x00\x1c\x00\x00\xe0\x50\x00\x40\xd3\x3f\x00\x2c\x0b\x0b\xe0\x70\xe0\x00\x07\x07\x0a\x24\xf0\x70\x00\x00\xfc\x02" +

	// glyph 382 for rune '&'
	"\x0c\xf1\x00\x01\x0c\x00\x14\x00\x80\xff\x00\xe0\xd2\x03\xf0\x80\x03\xe0\xc0\x03\xd0\xe2\x01\xc0\x7f\x00\x80\x1f\x00\xe0\x3e\x60\xb8\xb4\x70\x3c\xe0\x76\x3c\x80\x3f\x3c\x00\x2f\xf0\xd5\x3f\xc0\xff\xf1" +

	// glyph 432 for rune '\''
	"\x03\xf1\xf6\x01\x03\x5c\x85\x00" +

	// glyph 440 for rune '('
	"\x07\xf0\x04\x01\x07\x00\x08\x80\x03\x34\x00\x0b\xe0\x00\x3c\x00\x0b\xd0\x01\x38\x40\x85\x07\xd0\x01\xb0\x00\x3c\x00\x1d\x00\x0f\x00\x0b\x40\x07" +

	// glyph 476 for rune ')'
	"\x07\xf0\x04\x00\x06\x14\x00\x2c\x00\x1d\x00\x0f\x40\x07\xc0\x03\xe0\x00\x78\x00\x5d\x00\x2d\x40\x17\x80\x03\xf0\x00\x2c\x80\x03\xb0\x00\x0e\xd0\x00" +

	// glyph 513 for rune '*'
	"\x09\xf2\xfa\x00\x08\x00\x0d\x44\xd6\x24\xfd\xff\x00\x2f\x00\xee\x00\x1d\x0f\x24\x24" +

	// glyph 534 for rune '+'
	"\x0b\xf4\xff\x01\x0b\x00\x1c\x00\x40\x0b\x50\xa8\xbe\x1a\xff\xff\x0b\xd0\x02\x54\x00\x18\x00" +

	// glyph 557 for rune ','
	"\x04\xfe\x03\x00\x03\xf0\xc1\xd2\x51\x00" +

	// glyph 567 for rune '-'
	"\x06\xf9\xfb\x00\x05\xf8\x4f\xaa" +

	// glyph 575 for rune '.'
	"\x05\xfe\x00\x01\x04\x74\x01" +

	// glyph 582 for rune '/'
	"\x08\xf2\x01\x00\x08\x00\xc0\x01\x80\x03\x00\x0b\x00\x0d\x00\x38\x00\x70\x00\xd0\x00\xc0\x02\x00\x07\x00\x0e\x00\x2c\x00\x34\x00\xf0\x00\xc0\x01\x80\x03\x00" +

	// glyph 621 for rune '0'
	"\x0b\xf1\x00\x01\x0a\x00\x14\x00\xfc\x1f\xf0\xd1\x43\x07\xb0\x38\x00\x1f\x0f\x80\x57\x38\x00\x1f\x2d\xd0\xc1\x5b\x0f\xe0\x3f\x00" +

	// glyph 653 for rune '1'
	"\x0b\xf2\x00\x01\x07\x00\x3d\xf8\x4f\x8b\x03\xe0\x55\x55\x05" +

	// glyph 668 for rune '2'
	"\x0b\xf1\x00\x01\x0b\x00\x15\x00\xf4\x7f\x40\x1f\x7d\xf0\x00\x2c\x3c\x00\x0f\x00\xc0\x02\x00\x38\x00\x40\x0b\x00\xf0\x00\x00\x0f\x00\xf0\x00\x00\x1f\x00\xe0\x01\x00\x7e\x55\xc1\xff\xff\x01" +

	// glyph 715 for rune '3'
	"\x0b\xf1\x00\x01\x0a\x00\x15\x00\xfd\x1f\xf4\xd0\x83\x03\xb4\x14\x00\x0b\x00\xb4\x00\xd0\x03\xf0\x0f\x00\xe9\x03\x00\xb4\x00\x00\x8f\x02\xf0\x3c\x00\x4b\x5f\x3d\xd0\xff\x00" +

	// glyph 758 for rune '4'
	"\x0b\xf2\x00\x00\x0b\x00\x80\x0f\x00\xc0\x0f\x00\xf0\x0f\x00\x38\x0f\x00\x2c\x0f\x00\x0f\x0f\x40\x07\x0f\xc0\x02\x0f\xe0\x00\x0f\xf4\xff\xbf\xa4\xaa\xaf\x00\x00\x0f\x05" +

	// glyph 800 for rune '5'
	"\x0b\xf2\x00\x01\x0b\xd0\xff\x0f\xb4\x55\x01\x0e\x00\x14\xfc\xbf\x00\x6f\xbd\x00\x00\x3c\x00\x00\x1e\x00\x40\x47\x02\xe0\xc1\x02\x3c\xe0\xd6\x0b\xe0\x7f\x00" +

	// glyph 839 for rune '6'
	"\x0b\xf2\x00\x01\x0a\x00\xfd\x01\xf8\x06\xd0\x02\x00\x0f\x00\x70\x10\x40\xfb\x2f\xf8\xd2\x87\x0b\xf0\x38\x00\x1e\x1d\x80\xc3\x03\x3c\xf4\xf5\x01\xfd\x07" +

	// glyph 877 for rune '7'
	"\x0b\xf2\x00\x00\x0b\xf4\xff\x7f\x50\x55\x3d\x00\x00\x2c\x00\x00\x0e\x00\x00\x0b\x00\x40\x07\x00\xc0\x03\x00\xd0\x01\x00\xf0\x00\x00\x74\x00\x00\x3c\x00\x00\x2d\x00\x00\x0e\x00\x00\x0b\x00" +

	// glyph 924 for rune '8'
	"\x0b\xf1\x00\x01\x0a\x00\x14\x00\xfc\x1f\xf0\xd1\x47\x07\xb0\x78\x00\x4f\x07\xb0\xf0\xd1\x03\xfc\x0f\xe0\xe6\x43\x07\xb0\x3c\x00\xcf\x03\xe0\x78\x00\x0f\x5f\x7d\xc0\xff\x01" +

	// glyph 967 for rune '9'
	"\x0b\xf1\x00\x01\x0a\x00\x05\x00\xfd\x0f\xf0\xd1\x83\x03\x78\x3c\x00\xcb\x03\xf0\xe1\x01\x3d\x7c\xf9\x43\xff\x2d\x00\xc0\x01\x00\x0e\x00\xb8\x00\xfa\x02\xf0\x06\x00" +

	// glyph 1008 for rune ':'
	"\x05\xf5\x00\x01\x04\x20\x78\x10\x00\x55\xd1\x05" +

	// glyph 1020 for rune ';'
	"\x04\xf5\x03\x00\x03\x90\xf0\x40\x00\x55\x81\x07\x4f\x07\x02" +

	// glyph 1035 for rune '<'
	"\x0a\xf5\xfe\x00\x09\x00\x00\x04\x00\xb8\x00\xfd\x02\xfd\x02\xf4\x01\x00\xbf\x00\x00\xbe\x01\x00\xbe\x00\x00\x0a" +

	// glyph 1063 for rune '='
	"\x0b\xf6\xfc\x01\x0a\xa4\xaa\x86\xff\xbf\x00\x00\x10\x54\x55\xe1\xff\x2f" +

	// glyph 1081 for rune '>'
	"\x0a\xf5\xfe\x01\x0a\x04\x00\x80\x1f\x00\xe0\x1f\x00\xd0\x2f\x00\xc0\x0b\x40\x7f\x90\x7f\x80\x2f\x00\x28\x00\x00" +

	// glyph 1109 for rune '?'
	"\x09\xf1\x00\x01\x09\x00\x05\xc0\xff\xc1\x4b\x0f\x0b\xb4\x00\xc0\x02\x40\x07\x00\x0f\x00\x0f\x00\x1f\x00\x2d\x00\x74\x00\x00\x00\x01\x1d\x10" +

	// glyph 1144 for rune '@'
	"\x12\xf2\x05\x01\x11\x00\x90\xbf\x01\x00\xf0\x56\x3e\x00\xb4\x00\x80\x03\xb0\x00\x00\x38\xe0\x00\x29\xc0\xc1\x01\xee\x07\x4a\x03\x0e\x1c\x34\x0e\x2c\x30\xd0\x28\x34\xd0\x40\xb3\xe0\x40\x03\xcd\x82\x03\x0d\x24\x0a\x0e\x3c\x70\x38\xb4\xf9\xe1\xd0\x80\x2f\xfe\x00\x0b\x00\x00\x00\x74\x00\x00\x00\x40\x0b\x40\x00\x00\xe4\xff\x03\x00\x00\x54\x00\x00" +

	// glyph 1230 for rune 'A'
	"\x0d\xf2\x00\x00\x0d\x00\xe0\x02\x00\x00\x3f\x00\x00\xb4\x07\x00\x80\xf3\x00\x00\x2c\x0e\x00\xd0\xc1\x02\x00\x0f\x3c\x00\xb0\x80\x07\x80\x57\xb5\x00\xfc\xff\x0f\xd0\x56\xe5\x01\x0f\x00\x3c\xb0\x00\x80\x83\x07\x00\xb4" +

	// glyph 1284 for rune 'B'
	"\x0c\xf2\x00\x01\x0c\xf4\xff\x07\xb4\x95\x1f\xb4\x00\x3c\x45\x0b\xf4\x40\xff\x7f\x40\x5b\xf9\x41\x0b\xc0\x43\x0b\x80\x47\x0b\x80\x43\x0b\xc0\x43\xaf\xfa\x41\xff\x2f\x00" +

	// glyph 1326 for rune 'C'
	"\x0d\xf1\x00\x01\x0c\x00\x50\x00\x40\xff\x0f\xd0\x07\x3e\xf0\x00\xb4\xb4\x00\xf0\x78\x00\x50\x3c\x00\x00\x85\x03\x00\x80\x07\x00\x49\x0b\x00\x0f\x1f\x80\x07\xbc\xe5\x03\xe0\xbf\x00" +

	// glyph 1371 for rune 'D'
	"\x0d\xf2\x00\x01\x0c\xf4\xff\x02\xb4\x95\x0f\xb4\x00\x3d\xb4\x00\x78\xb4\x00\xb0\xb4\x00\xf0\x15\x2d\x00\x2d\x2d\x00\x1f\x2d\x80\x0b\xbd\xfa\x02\xfd\x6f\x00" +

	// glyph 1410 for rune 'E'
	"\x0b\xf2\x00\x01\x0b\xf4\xff\x2f\x6d\x55\x41\x0b\x00\x54\xf4\xff\x0b\x6d\x55\x41\x0b\x00\x54\xf4\xaa\x1a\xfd\xff\x0b" +

	// glyph 1439 for rune 'F'
	"\x0b\xf2\x00\x01\x0b\xf4\xff\x1f\x6d\x55\x41\x0b\x00\x54\xf4\xff\x07\xbd\xaa\x41\x0b\x00\x54\x05" +

	// glyph 1463 for rune 'G'
	"\x0e\xf1\x00\x01\x0c\x00\x50\x00\x40\xff\x0f\xd0\x07\x3d\xf0\x00\xb4\xb4\x00\xf0\x78\x00\x00\x38\x00\x00\x3c\x00\x00\x3c\xc0\xff\x38\x40\xf5\x78\x00\xe0\xf4\x00\xe0\xf0\x01\xf0\xc0\x5b\xfe\x00\xfe\x1f" +

	// glyph 1513 for rune 'H'
	"\x0e\xf2\x00\x01\x0d\xb4\x00\xd0\x56\x45\xff\xff\x2f\x6d\x55\xb9\xb4\x00\xd0\x56\x05" +

	// glyph 1534 for rune 'I'
	"\x05\xf2\x00\x02\x04\x6c\x55\x55\x55" +

	// glyph 1543 for rune 'J'
	"\x0b\xf2\x00\x00\x0a\x00\x00\x6d\x55\x15\x00\x40\x47\x07\xd0\xc1\x03\x3c\xe0\xd6\x0b\xe0\xbf\x00" +

	// glyph 1567 for rune 'K'
	"\x0d\xf2\x00\x01\x0c\xb4\x00\x7c\xb4\x00\x2e\xb4\x80\x0b\xb4\xd0\x03\xb4\xf4\x00\xb4\x3d\x00\xb4\x3f\x00\xf4\xbb\x00\xf4\xf0\x01\xb4\xc0\x03\xb4\x40\x0f\xb4\x00\x2e\xb4\x00\x7c\xb4\x00\xf0" +

	// glyph 1614 for rune 'L'
	"\x0b\xf2\x00\x01\x0b\xb4\x00\x40\x55\x55\x45\xaf\xaa\xd0\xff\x7f" +

	// glyph 1630 for rune 'M'
	"\x11\xf2\x00\x01\x10\xf4\x01\x00\xfc\xf4\x03\x00\xfd\xf4\x03\x00\xfe\xb4\x0b\x00\xfb\x74\x0f\x80\xf7\x74\x1d\xc0\xf3\x74\x3c\xd0\xf1\x74\x38\xe0\xf0\xb4\xb0\xb0\xf0\xb4\xf0\x78\xf0\xb4\xd0\x3d\xf0\xb4\xc0\x1f\xf0\xb4\x80\x0f\xf0\xb4\x00\x0b\xf0" +

	// glyph 1691 for rune 'N'
	"\x0e\xf2\x00\x01\x0d\xf4\x00\xd0\xd2\x0b\x40\x4b\x7f\x00\x2d\xed\x03\xb4\xb4\x2e\xd0\xd2\xf2\x41\x4b\x0b\x0f\x2d\x2d\xb8\xb4\xb4\xc0\xd3\xd2\x02\x7d\x4b\x0b\xe0\x2f\x2d\x00\xbf\xb4\x00\xf4\xd2\x02\x80\x0b" +

	// glyph 1742 for rune 'O'
	"\x0e\xf1\x00\x01\x0d\x00\x50\x00\x00\xfd\x2f\x00\xbd\xe5\x03\x3c\x00\x2d\xb4\x00\xf0\xe0\x01\x80\xc7\x03\x00\x6d\xe1\x00\x40\x87\x07\x00\x1e\x2d\x00\x3c\xf0\x01\xb8\x00\x6f\xfe\x00\xe0\xbf\x00" +

	// glyph 1790 for rune 'P'
	"\x0d\xf2\x00\x01\x0c\xf4\xff\x0b\xb4\x55\x3e\xb4\x00\xb8\xb4\x00\xf0\xd1\x02\xd0\xd2\x02\xf4\xd0\xff\x7f\xd0\x56\x01\xd0\x02\x00\x54\x01" +

	// glyph 1824 for rune 'Q'
	"\x0e\xf1\x02\x01\x0d\x00\x50\x00\x00\xfd\x2f\x00\x7d\xe5\x03\x3c\x00\x2e\xb4\x00\xf0\xe0\x00\x80\xc7\x03\x00\x1d\x0f\x00\xb4\x3c\x00\xd0\x85\x03\x00\x1e\x2d\x00\x3c\xf0\x01\xb8\x00\x6f\xbe\x00\xe0\xff\x01\x00\x00\x2f\x00\x00\xf0\x00" +

	// glyph 1882 for rune 'R'
	"\x0c\xf2\x00\x01\x0c\xf4\xff\x07\xb4\x55\x1f\xb4\x00\x3c\xb4\x00\x38\xb4\x00\x78\xb4\x00\x3c\xb4\x55\x1f\xf4\xff\x07\xb4\xd0\x03\xb4\x40\x0b\xb4\x00\x0f\xb4\x00\x2d\xb4\x00\x3c\xb4\x00\xb4" +

	// glyph 1929 for rune 'S'
	"\x0c\xf1\x00\x01\x0b\x00\x14\x00\xf0\xff\x01\x1f\xf4\xe1\x01\xf0\x38\x00\x3c\x2e\x00\x00\x6f\x00\x00\xfe\x02\x00\xf8\x0b\x00\xd0\x0b\x00\xc0\xf3\x00\xf0\x7c\x00\x3c\xbc\xe5\x07\xf8\x7f\x00" +

	// glyph 1976 for rune 'T'
	"\x0c\xf2\x00\x00\x0c\xf8\xff\xff\x41\xd5\x57\x01\x00\x0f\x40\x55\x55\x05" +

	// glyph 1994 for rune 'U'
	"\x0d\xf2\x00\x01\x0c\x38\x00\xb0\x55\x55\xd1\x01\xd0\xc1\x03\xf0\x80\x5f\xbd\x00\xfd\x1f\x00" +

	// glyph 2017 for rune 'V'
	"\x0d\xf2\x00\x00\x0c\xb4\x00\xc0\xc3\x03\x00\x0f\x0e\x00\x1e\xb4\x00\x3c\xc0\x03\xb0\x00\x1e\xe0\x01\xb0\xc0\x03\xc0\x03\x0b\x00\x1e\x1e\x00\xb0\x3c\x00\xc0\xb7\x00\x00\xfd\x00\x00\xf0\x03\x00\x80\x07\x00" +

	// glyph 2068 for rune 'W'
	"\x12\xf2\x00\x01\x11\x2c\x00\x0f\x80\xf3\x00\x7d\x00\xcf\x03\xf8\x02\x2c\x1e\xf0\x0f\x74\xb4\xc0\x3a\xe0\xc0\x43\xc3\xc1\x03\x0f\x0e\x0b\x0b\x38\x2c\x38\x1d\xd0\x75\xd0\x38\x00\xeb\x00\xf7\x00\xfc\x03\xec\x02\xe0\x07\xe0\x07\x40\x0f\x40\x1f\x00\x3c\x00\x3c\x00" +

	// glyph 2133 for rune 'X'
	"\x0d\xf2\x00\x01\x0c\x7c\x00\x78\xf0\x00\x3d\xe0\x02\x0f\xc0\x83\x0b\x40\xdf\x03\x00\xfe\x01\x00\xbc\x00\x00\xfd\x00\x00\xff\x01\x80\xcb\x03\xc0\x43\x0f\xf0\x01\x1f\xb4\x00\x3c\x3c\x00\xb8" +

	// glyph 2180 for rune 'Y'
	"\x0c\xf2\x00\x00\x0c\xb4\x00\xe0\xc1\x03\xc0\x03\x2d\x80\x07\xf0\x00\x0f\x40\x0b\x0e\x00\x38\x2d\x00\xc0\x3f\x00\x00\xbe\x00\x00\xf0\x00\x54\x05" +

	// glyph 2216 for rune 'Z'
	"\x0c\xf2\x00\x01\x0b\xfc\xff\x3f\x55\x95\x0b\x00\xf0\x00\x00\x1f\x00\xe0\x02\x00\x3d\x00\xc0\x03\x00\x78\x00\x40\x0f\x00\xf0\x00\x00\x1e\x00\xd0\x02\x00\xbc\xaa\x2a\xff\xff\x0f" +

	// glyph 2260 for rune '['
	"\x05\xef\x03\x01\x05\x50\xe1\x8f\x47\x55\x55\x55\x85\x1b\xfe" +

	// glyph 2275 for rune '\\'
	"\x08\xf2\x01\x00\x08\x74\x00\xc0\x03\x00\x0e\x00\xb0\x00\xc0\x03\x00\x1d\x00\xb0\x00\x80\x03\x00\x1d\x00\xf0\x00\x80\x03\x00\x2c\x00\xe0\x00\x40\x07\x00\x2c" +

	// glyph 2314 for rune ']'
	"\x05\xef\x03\x00\x04\x50\xf0\x0f\x3d\xf0\x55\x55\x55\x45\x3d\xff" +

	// glyph 2330 for rune '^'
	"\x08\xf2\xf9\x01\x08\x80\x03\xc0\x07\xe0\x0f\xb0\x1d\x34\x2c\x3c\x38\x1c\x70" +

	// glyph 2349 for rune '_'
	"\x09\x00\x02\x00\x09\xfc\xff\x4f\x55\x55" +

	// glyph 2359 for rune '`'
	"\x06\xf1\xf4\x01\x05\x3c\xc0\x02\x1d" +

	// glyph 2368 for rune 'a'
	"\x0b\xf5\x00\x01\x0a\x40\x6a\x00\xbe\x2f\xb4\xc0\x03\x01\x74\x00\x95\x07\xfe\x7f\xb4\x40\xc7\x03\x74\x3c\xc0\x87\x9f\x7f\xe0\x6f\x0b" +

	// glyph 2401 for rune 'b'
	"\x0b\xf1\x00\x01\x0b\x38\x00\x40\x85\x93\x06\xe0\xff\x0f\xf8\x80\x0b\x1e\xc0\x83\x03\xe0\xe0\x00\x78\x38\x00\x4e\xb8\x00\x0f\xbe\xf9\x81\xf3\x1f\x00" +

	// glyph 2438 for rune 'c'
	"\x0a\xf5\x00\x01\x0a\x40\x6a\x00\xfe\x2f\xb4\x80\xc7\x03\xb0\x3c\x00\xc0\x02\x00\x3c\x00\xc0\x03\x10\x78\x00\x0b\x6f\x3d\x80\xbf\x00" +

	// glyph 2471 for rune 'd'
	"\x0b\xf1\x00\x01\x0a\x00\x00\x5f\x01\x69\x3c\xf8\xff\xd3\x03\x3e\x0f\xc0\x57\xe1\x01\x3d\xbc\xf9\x03\xff\x3d" +

	// glyph 2498 for rune 'e'
	"\x0b\xf5\x00\x01\x0a\x00\x6a\x00\xfd\x2f\xb4\x80\x83\x03\xb0\x7c\x55\xcb\xff\xff\x3c\x00\x10\x1e\x40\xc0\x5b\x1f\xe0\x7f\x00" +

	// glyph 2529 for rune 'f'
	"\x07\xf1\x00\x00\x07\x00\xfd\x40\x5f\x80\x03\x81\xaf\xd0\xff\x01\x0e\x54\x55\x01" +

	// glyph 2549 for rune 'g'
	"\x0b\xf5\x04\x01\x0a\x40\x1a\x05\xfe\xff\xf4\x80\xcf\x03\xf0\x55\x78\x40\x0f\x6f\xfe\xc0\x7f\x0f\x00\xb0\x20\x40\x4b\x5f\x3e\x90\xbf\x00" +

	// glyph 2583 for rune 'h'
	"\x0b\xf1\x00\x01\x0a\x38\x00\x50\xe1\xa4\x01\xee\xff\xe0\x03\x1e\x1e\xc0\xe2\x00\x6c\x55\x01" +

	// glyph 2606 for rune 'i'
	"\x05\xf1\x00\x01\x04\x10\x78\x20\x00\x24\x74\x55\x55\x01" +

	// glyph 2620 for rune 'j'
	"\x05\xf1\x04\xff\x04\x00\x01\x78\x40\x02\x00\x40\x02\x38\x55\x55\x15\xf4\xe0\x07" +

	// glyph 2640 for rune 'k'
	"\x0a\xf1\x00\x01\x0a\x38\x00\x50\xe1\x00\x09\x0e\xb8\xe0\xe0\x02\x4e\x0f\xe0\x3e\x00\xfe\x03\xe0\xf7\x00\x0e\x2e\xe0\xc0\x07\x0e\xf0\xe0\x00\x2d" +

	// glyph 2676 for rune 'l'
	"\x05\xf1\x00\x01\x04\x74\x55\x55\x55\x05" +

	// glyph 2686 for rune 'm'
	"\x12\xf5\x00\x01\x10\x24\x69\x40\x1a\xb8\xff\xf7\x7f\xf8\x80\x1f\xf0\x38\x40\x0f\xe0\x38\x00\x0b\xe0\x55\x05" +

	// glyph 2713 for rune 'n'
	"\x0b\xf5\x00\x01\x0a\x24\x69\x80\xfb\x3f\xf8\x80\x87\x07\xb0\x38\x00\x5b\x55" +

	// glyph 2732 for rune 'o'
	"\x0b\xf5\x00\x01\x0b\x00\x6a\x00\xf4\xbf\x40\x0f\xb8\xe0\x00\x3c\x3c\x00\x1d\x0b\x40\xc7\x03\xd0\xf1\x00\x78\x74\x00\x0f\xbc\xf5\x01\xf8\x1f\x00" +

	// glyph 2768 for rune 'p'
	"\x0b\xf5\x04\x01\x0b\x24\x69\x00\xee\xff\x80\x0f\xb8\xe0\x00\x3c\x38\x00\x0e\x0e\x80\x87\x03\xe0\xe0\x00\x3c\x78\x00\x0f\xbe\xf5\x81\xf7\x1f\xe0\x00\x00\x15" +

	// glyph 2807 for rune 'q'
	"\x0b\xf5\x04\x01\x0a\x40\x1a\x05\xfe\xff\xf4\x80\xcf\x03\xf0\x55\x78\x00\x0f\x6f\xfe\xc0\x7f\x0f\x00\xf0\x15" +

	// glyph 2834 for rune 'r'
	"\x07\xf5\x00\x01\x07\x24\x19\xfe\x87\x1f\xe0\x01\x38\x40\x55\x01" +

	// glyph 2850 for rune 's'
	"\x0a\xf5\x00\x01\x09\x40\x2a\xc0\xff\x87\x07\x3c\x0e\x90\xf4\x01\x40\xff\x01\x40\x2f\x01\xe0\x3c\x80\xd3\x97\x0f\xfd\x0b" +

	// glyph 2880 for rune 't'
	"\x07\xf3\x00\x00\x06\xc0\x43\xe8\x1b\xff\x0b\x3c\x54\x15\xf0\x06\xf4\x03" +

	// glyph 2898 for rune 'u'
	"\x0b\xf5\x00\x01\x0a\x24\x00\x86\x03\xb0\x55\x85\x07\xb4\xf0\xe5\x0b\xfd\xb7" +

	// glyph 2917 for rune 'v'
	"\x0a\xf5\x00\x00\x09\x24\x00\x49\x0b\xf0\xf0\x00\x0b\x0e\x38\xc0\xc2\x03\x3c\x2c\x40\xd3\x00\xb0\x0f\x00\x7f\x00\xd0\x03\x00\x2c\x00" +

	// glyph 2950 for rune 'w'
	"\x0f\xf5\x00\x00\x0f\x24\x00\x02\x60\x74\x80\x0b\x78\xf0\xc0\x0f\x3c\xe0\xc0\x0e\x2c\xd0\xd1\x1c\x1d\xc0\xb2\x38\x0e\xc0\x73\x34\x0f\x80\x3b\x70\x0b\x00\x2f\xf0\x03\x00\x1f\xd0\x03\x00\x0e\xc0\x02" +

	// glyph 2999 for rune 'x'
	"\x0a\xf5\x00\x00\x09\x60\x00\x0a\x0f\xb4\xc0\xc3\x03\x74\x1e\x00\xbf\x00\xc0\x03\x00\x7e\x00\xb4\x0f\xc0\xc3\x03\x1e\x78\xf4\x00\x0f" +

	// glyph 3032 for rune 'y'
	"\x09\xf5\x04\x00\x09\x24\x00\x49\x0b\xf0\xf0\x40\x0b\x0e\x38\xd0\xc2\x03\x3c\x1d\x80\xe3\x00\xb0\x0f\x00\x7f\x00\xe0\x03\x00\x2c\x00\xd0\x01\x00\x0e\x00\xb9\x00\xf0\x02\x00" +

	// glyph 3075 for rune 'z'
	"\x0a\xf5\x00\x01\x09\xa8\xaa\xf2\xff\x0f\x00\x1e\x00\x2e\x00\x3d\x00\x3c\x00\x78\x00\xb8\x00\xf4\x00\xf0\x56\xc5\xff\x3f" +

	// glyph 3105 for rune '{'
	"\x07\xf0\x04\x00\x07\x00\x10\x00\x7c\x00\x1e\x00\x0b\x01\x1d\x14\xf0\x00\x7d\x40\xc0\x03\x40\x07\x05\xb0\x10\x80\x07\x00\x1f\x00\x04" +

	// glyph 3138 for rune '|'
	"\x05\xf2\x03\x01\x03\x74\x55\x55\x55\x05\x02" +

	// glyph 3149 for rune '}'
	"\x07\xf0\x04\x00\x06\x04\x00\x2e\x00\x1e\x00\x0f\x15\xe0\x00\xb4\x00\xf8\x01\x2d\x80\x03\xf0\x50\x81\x07\xb8\x00\x01\x00" +

	// glyph 3179 for rune '~'
	"\x0e\xf8\xfc\x01\x0c\x80\x0b\x80\xf0\x7f\xd0\x38\xf0\xba\x24\x80\x2f" +

	"")
