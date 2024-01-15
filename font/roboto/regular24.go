// File generated using:
// 	go run ./generate -font=roboto/Roboto-Regular.ttf -size=24 -dpi=72 -package=roboto

package roboto

import (
	"github.com/aykevl/tinygl/font"
)

// Font statistics:
// - total size:      4186
// - glyph metadata:  475
// - glyph mask data: 3510

var Regular24 = font.Make("" +
	"\x00" + // version: 0
	"\x18" + // size:   24
	"\x1d" + // height: 29
	"\x17" + // ascent: 23

	// Runes 32..126 (95 runes)
	"\x5f\x00" + // number of runes (95)
	"\x20\x00\x00" + // start rune (32)
	"\xc9\x00" + // " " at index 201
	"\xce\x00" + // "!" at index 206
	"\xdb\x00" + // "\"" at index 219
	"\xe8\x00" + // "#" at index 232
	"\x2d\x01" + // "$" at index 301
	"\x72\x01" + // "%" at index 370
	"\xc4\x01" + // "&" at index 452
	"\x09\x02" + // "'" at index 521
	"\x11\x02" + // "(" at index 529
	"\x3f\x02" + // ")" at index 575
	"\x6d\x02" + // "*" at index 621
	"\x89\x02" + // "+" at index 649
	"\x9e\x02" + // "," at index 670
	"\xaa\x02" + // "-" at index 682
	"\xb3\x02" + // "." at index 691
	"\xba\x02" + // "/" at index 698
	"\xef\x02" + // "0" at index 751
	"\x1d\x03" + // "1" at index 797
	"\x2f\x03" + // "2" at index 815
	"\x6f\x03" + // "3" at index 879
	"\xaa\x03" + // "4" at index 938
	"\xe1\x03" + // "5" at index 993
	"\x11\x04" + // "6" at index 1041
	"\x4e\x04" + // "7" at index 1102
	"\x8b\x04" + // "8" at index 1163
	"\xc1\x04" + // "9" at index 1217
	"\xfc\x04" + // ":" at index 1276
	"\x0a\x05" + // ";" at index 1290
	"\x1c\x05" + // "<" at index 1308
	"\x40\x05" + // "=" at index 1344
	"\x59\x05" + // ">" at index 1369
	"\x7f\x05" + // "?" at index 1407
	"\xb1\x05" + // "@" at index 1457
	"\x2f\x06" + // "A" at index 1583
	"\x78\x06" + // "B" at index 1656
	"\xb5\x06" + // "C" at index 1717
	"\xf7\x06" + // "D" at index 1783
	"\x31\x07" + // "E" at index 1841
	"\x51\x07" + // "F" at index 1873
	"\x6e\x07" + // "G" at index 1902
	"\xb3\x07" + // "H" at index 1971
	"\xcd\x07" + // "I" at index 1997
	"\xd7\x07" + // "J" at index 2007
	"\xf1\x07" + // "K" at index 2033
	"\x32\x08" + // "L" at index 2098
	"\x41\x08" + // "M" at index 2113
	"\x93\x08" + // "N" at index 2195
	"\xd4\x08" + // "O" at index 2260
	"\x0f\x09" + // "P" at index 2319
	"\x3a\x09" + // "Q" at index 2362
	"\x83\x09" + // "R" at index 2435
	"\xc0\x09" + // "S" at index 2496
	"\x04\x0a" + // "T" at index 2564
	"\x18\x0a" + // "U" at index 2584
	"\x35\x0a" + // "V" at index 2613
	"\x7e\x0a" + // "W" at index 2686
	"\xdd\x0a" + // "X" at index 2781
	"\x1e\x0b" + // "Y" at index 2846
	"\x4e\x0b" + // "Z" at index 2894
	"\x8c\x0b" + // "[" at index 2956
	"\x9f\x0b" + // "\\" at index 2975
	"\xd9\x0b" + // "]" at index 3033
	"\xec\x0b" + // "^" at index 3052
	"\x06\x0c" + // "_" at index 3078
	"\x11\x0c" + // "`" at index 3089
	"\x1b\x0c" + // "a" at index 3099
	"\x45\x0c" + // "b" at index 3141
	"\x73\x0c" + // "c" at index 3187
	"\x9a\x0c" + // "d" at index 3226
	"\xc8\x0c" + // "e" at index 3272
	"\xf4\x0c" + // "f" at index 3316
	"\x10\x0d" + // "g" at index 3344
	"\x49\x0d" + // "h" at index 3401
	"\x63\x0d" + // "i" at index 3427
	"\x6f\x0d" + // "j" at index 3439
	"\x83\x0d" + // "k" at index 3459
	"\xb3\x0d" + // "l" at index 3507
	"\xbd\x0d" + // "m" at index 3517
	"\xdd\x0d" + // "n" at index 3549
	"\xf3\x0d" + // "o" at index 3571
	"\x20\x0e" + // "p" at index 3616
	"\x51\x0e" + // "q" at index 3665
	"\x7f\x0e" + // "r" at index 3711
	"\x8f\x0e" + // "s" at index 3727
	"\xb8\x0e" + // "t" at index 3768
	"\xcf\x0e" + // "u" at index 3791
	"\xe5\x0e" + // "v" at index 3813
	"\x11\x0f" + // "w" at index 3857
	"\x54\x0f" + // "x" at index 3924
	"\x7d\x0f" + // "y" at index 3965
	"\xb8\x0f" + // "z" at index 4024
	"\xdf\x0f" + // "{" at index 4063
	"\x0a\x10" + // "|" at index 4106
	"\x15\x10" + // "}" at index 4117
	"\x42\x10" + // "~" at index 4162

	// mark the end of the rune tables
	"\x00\x00" +

	// glyph 201 for rune ' '
	"\x06\x00\x00\x00\x00" +

	// glyph 206 for rune '!'
	"\x06\xef\x00\x02\x05\x3c\x55\x55\x15\x40\xf1\xf1\x00" +

	// glyph 219 for rune '"'
	"\x08\xee\xf4\x01\x07\x74\x1c\x0d\x57\x34\x0c\x09\x03" +

	// glyph 232 for rune '#'
	"\x0f\xef\x00\x01\x0f\x00\xe0\xc0\x03\x00\x38\xb0\x00\x00\x0f\x2c\x00\xc0\x42\x07\x00\xb4\xe0\x01\xf8\xff\xff\x07\xd5\x57\x5f\x00\xf0\xc0\x02\x00\x2c\x74\x00\x40\x07\x0e\x00\xe5\xd6\x17\xe0\xff\xff\x2f\x00\x0f\x2d\x00\xc0\x42\x07\x00\x70\xd0\x00\x00\x0d\x38\x00\x80\x03\x0f\x00" +

	// glyph 301 for rune '$'
	"\x0d\xec\x03\x01\x0c\x00\xa0\x00\x00\xf0\x00\x01\xfc\x2f\x00\xbf\xfe\x80\x0b\xf0\xc1\x03\xd0\xc3\x03\xc0\xc3\x07\x00\x80\x1f\x00\x00\xff\x01\x00\xf4\x2f\x00\x40\xff\x00\x00\xf4\x02\x00\xd0\xe3\x02\xc0\xd3\x03\xc0\xc3\x07\xe0\x83\xbf\xfe\x00\xfd\x2f\x00\xc0\x02\x04\x00\x05\x00" +

	// glyph 370 for rune '%'
	"\x12\xee\x00\x01\x11\x00\x01\x00\x00\x40\xff\x00\x00\x40\x4b\x0f\x00\x00\x0e\x38\xc0\x01\x3c\xd0\xc0\x03\xe0\x40\x43\x07\x40\x07\x0f\x0b\x00\xfc\x1f\x0e\x00\x40\x06\x1d\x00\x00\x00\x2c\x00\x00\x00\x38\xa4\x01\x00\xb0\xf8\x2f\x00\xf0\xb0\xe0\x01\xd0\xd1\x00\x07\xc0\x42\x03\x2c\x80\x03\x1d\x70\x00\x04\xf0\xf1\x00\x00\x00\xff\x01" +

	// glyph 452 for rune '&'
	"\x0f\xef\x00\x01\x0f\x00\xfe\x07\x00\xe0\xeb\x07\x00\x3c\xd0\x02\x40\x0f\xf0\x00\xd0\x03\x2d\x00\xf0\xd1\x03\x00\xf4\x7e\x00\x00\xfc\x03\x00\x40\xbf\x00\x00\xf8\xbd\xc0\x42\x1f\x7c\xb4\xe0\x02\x7d\x2e\x7c\x00\xfe\x03\x2e\x00\xbe\x40\x1f\x80\x2f\x80\xaf\xfe\x2f\x40\xff\x0b\x1f" +

	// glyph 521 for rune '\''
	"\x04\xee\xf4\x01\x03\x7c\x15\x05" +

	// glyph 529 for rune '('
	"\x08\xed\x05\x01\x08\x00\xa0\x00\x78\x00\x2e\x00\x0f\x80\x07\xc0\x03\xe0\x02\xf0\x01\xf0\x00\xd1\x03\xd0\x02\xd0\x03\x04\x0f\x50\xe0\x02\xd0\x03\xc0\x03\x40\x0b\x00\x1e\x00\x3c\x00\xb0" +

	// glyph 575 for rune ')'
	"\x08\xed\x05\x00\x07\x34\x00\xf0\x00\xd0\x02\x80\x07\x00\x0f\x00\x1e\x00\x2d\x00\x3c\x01\xe0\x01\xe0\x56\x00\x78\x00\x7c\x00\x3c\x00\x3d\x00\x1e\x00\x0f\x40\x0b\xc0\x03\xf0\x00\x74\x00" +

	// glyph 621 for rune '*'
	"\x0a\xef\xf9\x00\x0a\x00\x38\x40\xd1\xe6\xb8\xf4\xff\x2f\x80\x7f\x00\xf0\x0f\x00\x5e\x0f\xc0\x83\x07\x10\x80\x00" +

	// glyph 649 for rune '+'
	"\x0e\xf2\xff\x01\x0d\x00\xf4\x00\x54\xf1\xff\xff\x1b\x00\x3d\x00\x55\x00\x10\x00\x00" +

	// glyph 670 for rune ','
	"\x05\xfd\x03\x00\x04\x90\x81\x1b\x78\xf0\xd0\x01" +

	// glyph 682 for rune '-'
	"\x07\xf8\xfa\x00\x06\xf8\x3f\xa9\x0a" +

	// glyph 691 for rune '.'
	"\x06\xfd\x00\x02\x04\x14\x1f" +

	// glyph 698 for rune '/'
	"\x0a\xef\x02\x00\x09\x00\x00\x0f\x00\xb4\x00\x80\x03\x00\x3c\x00\xd0\x01\x00\x0f\x00\xb0\x00\x80\x07\x00\x3c\x00\xd0\x02\x00\x0e\x00\xf0\x00\x40\x07\x00\x3c\x00\xc0\x02\x00\x0e\x00\xf0\x00\x40\x07\x00\x14\x00\x00" +

	// glyph 751 for rune '0'
	"\x0d\xee\x00\x01\x0c\x00\x10\x00\x40\xff\x0b\xd0\xaf\x2f\xf0\x02\x3c\xf0\x00\xb8\xb4\x00\xf4\xb8\x00\xf0\x55\xd1\x02\xc0\xd3\x03\xd0\xc3\x03\xe0\x82\x0b\xf0\x01\xbf\xbe\x00\xfc\x2f\x00" +

	// glyph 797 for rune '1'
	"\x0d\xef\x00\x02\x09\x00\xb8\x90\xbf\xfc\xbf\x2c\xb8\x00\xb8\x55\x55\x55" +

	// glyph 815 for rune '2'
	"\x0d\xee\x00\x01\x0d\x00\x10\x00\x00\xfd\x2f\x00\xbe\xfa\x03\x3d\x00\x1f\xb8\x00\xb8\xf0\x01\xd0\x03\x00\x80\x0b\x00\x00\x0f\x00\x00\x2e\x00\x00\x3d\x00\x00\x3d\x00\x00\x7c\x00\x00\x7c\x00\x00\xbc\x00\x00\xbc\x00\x00\xb8\x00\x00\xf4\xff\xff\xe1\xff\xff\x0b" +

	// glyph 879 for rune '3'
	"\x0d\xee\x00\x01\x0c\x00\x10\x00\x40\xff\x07\xe0\xab\x2f\xf4\x00\x7c\xb8\x00\x78\x14\x00\xb8\x00\x00\x78\x00\x00\x3d\x00\xe9\x0f\x00\xfd\x0b\x00\x40\x3f\x00\x00\xbc\x00\x00\xf4\x14\x00\xf4\x78\x00\xb4\xf4\x00\x7c\xe0\xab\x2f\x40\xff\x07" +

	// glyph 938 for rune '4'
	"\x0d\xef\x00\x00\x0d\x00\x00\xbc\x00\x00\xe0\x0b\x00\x40\xbf\x00\x00\xbc\x0b\x00\xe0\xb9\x00\x40\x8f\x0b\x00\x3c\xb8\x00\xe0\x81\x0b\x40\x0f\xb8\x00\x3c\x80\x0b\xe0\x02\xb8\x00\xbf\xea\xaf\xf4\xff\xff\x0f\x00\x80\x0b\x15" +

	// glyph 993 for rune '5'
	"\x0d\xef\x00\x02\x0d\xf0\xff\x3f\xc1\x03\x00\x14\x2d\x00\x00\xfd\xbf\x00\xfd\xff\x03\x29\xc0\x0f\x00\x00\x0f\x00\x00\x1f\x00\x00\x1e\x05\x00\x1f\x0f\x00\x0f\x2e\x80\x0f\xfc\xfa\x03\xe0\xbf\x00" +

	// glyph 1041 for rune '6'
	"\x0d\xef\x00\x01\x0d\x00\xe0\x0f\x00\xf4\x2f\x00\xf4\x02\x00\xf0\x01\x00\xe0\x02\x00\xc0\x03\x00\x00\xcf\xbf\x00\xfd\xeb\x0f\xf4\x03\xbc\xd0\x03\xd0\x43\x0b\x00\x0f\x3d\x00\x7c\xf0\x00\xf0\xc0\x07\xc0\x03\x3d\xc0\x0b\xe0\xeb\x0f\x00\xfd\x0b\x00" +

	// glyph 1102 for rune '7'
	"\x0d\xef\x00\x01\x0d\xfc\xff\xff\xa1\xaa\xea\x07\x00\x00\x0f\x00\x00\x1e\x00\x00\x3c\x00\x00\xb8\x00\x00\xf0\x00\x00\xe0\x02\x00\xc0\x03\x00\x40\x0b\x00\x00\x1f\x00\x00\x3d\x00\x00\x7c\x00\x00\xf4\x00\x00\xf0\x01\x00\xc0\x03\x00\x80\x0b\x00\x00" +

	// glyph 1163 for rune '8'
	"\x0d\xef\x00\x01\x0c\x40\xff\x0b\xd0\xaf\x2f\xf0\x02\x7c\xf0\x00\xb8\xf0\x00\xb4\xf0\x00\xb8\xe0\x02\x3d\xc0\xef\x1f\x40\xff\x0b\xe0\x07\x3e\xf0\x00\xb8\xb8\x00\xf0\xd1\x02\xc0\xc3\x07\xf0\x82\xbf\xfe\x00\xfd\x2f\x00" +

	// glyph 1217 for rune '9'
	"\x0d\xee\x00\x01\x0c\x00\x10\x00\x40\xff\x03\xd0\xeb\x1f\xf0\x01\x3e\xb8\x00\x7c\x78\x00\xb8\x7c\x00\xb4\x78\x00\xf4\xb8\x00\xf4\xf0\x01\xfd\xe0\xab\xff\x80\xff\xb5\x00\x00\x78\x00\x00\x3c\x00\x00\x2e\x00\xd0\x0f\x40\xff\x03\x40\x6f\x00" +

	// glyph 1276 for rune ':'
	"\x06\xf3\x00\x01\x04\xa0\xf4\xa0\x00\x55\x05\x41\x0f\x0f" +

	// glyph 1290 for rune ';'
	"\x05\xf3\x03\x00\x04\x80\x81\x0f\x18\x00\x55\x05\x18\xb4\x81\x0b\x0f\x1d" +

	// glyph 1308 for rune '<'
	"\x0c\xf3\xfe\x01\x0b\x00\x00\x14\x00\xe0\x07\x80\xff\x00\xfe\x02\xf8\x07\x00\x2f\x00\x80\xbf\x00\x00\xfd\x07\x00\xf4\x1f\x00\xd0\x07\x00\x40\x01" +

	// glyph 1344 for rune '='
	"\x0d\xf4\xfc\x02\x0c\x54\x55\x05\xff\xff\x4b\x55\x55\x01\x00\x00\xa1\xaa\x6a\xfc\xff\x2f\x55\x55\x01" +

	// glyph 1369 for rune '>'
	"\x0d\xf3\xfe\x01\x0c\x14\x00\x00\xf4\x02\x00\xe0\x6f\x00\x00\xfd\x07\x00\x90\x7f\x00\x00\xbd\x00\xe0\x2f\x40\xfe\x02\xf4\x2f\x00\xf4\x01\x00\x14\x00\x00" +

	// glyph 1407 for rune '?'
	"\x0b\xee\x00\x01\x0b\x00\x04\x00\xf4\x7f\x00\xbf\xbe\xf0\x02\x3d\x3c\x00\x1f\x00\xc0\x07\x00\xf0\x00\x00\x2e\x00\xe0\x03\x00\x3e\x00\xc0\x03\x00\xb8\x00\x00\x1f\x00\x00\x00\x50\x00\x1f\x00\x80\x07\x00" +

	// glyph 1457 for rune '@'
	"\x16\xef\x06\x01\x15\x00\x00\xa9\x1a\x00\x00\x40\xff\xff\x07\x00\x80\x1f\x00\x7d\x00\x80\x0b\x00\x80\x07\x40\x07\x00\x00\x38\x00\x0f\x40\x06\xc0\x02\x0d\xc0\xff\x02\x0e\x3c\xc0\x07\x0f\x34\xb0\x80\x07\x3c\xc0\xd0\x01\x0f\xb0\x00\x47\x03\x2c\xc0\x02\x1c\x0e\x74\x00\x07\x70\x38\xe0\x01\x1d\xd0\xe0\x40\x07\x78\x80\x43\x07\x2d\xf0\x01\x0b\x2c\xf0\xba\x5f\x0f\xf0\x40\xbf\xf4\x0b\x80\x07\x00\x00\x00\x00\x3c\x00\x00\x00\x00\xd0\x07\x00\x00\x00\x00\xfc\x45\x29\x00\x00\x40\xff\x7f\x00\x00\x00\x00\x01\x00\x00" +

	// glyph 1583 for rune 'A'
	"\x10\xef\x00\x00\x0f\x00\x40\x0f\x00\x00\xc0\x1f\x00\x00\xc0\x3f\x00\x00\xe0\x3e\x00\x00\xf0\xb8\x00\x00\xf4\xf0\x00\x00\x78\xf0\x01\x00\x3c\xd0\x02\x00\x2d\xc0\x03\x00\x1f\x80\x07\x00\x5f\x95\x0f\x80\xff\xff\x0f\xc0\xab\xaa\x2f\xd0\x03\x00\x3c\xe0\x02\x00\x7c\xf0\x00\x00\xb8\xf4\x00\x00\xf0" +

	// glyph 1656 for rune 'B'
	"\x0f\xef\x00\x02\x0e\xfc\xff\x0b\xf0\xff\xff\xc0\x03\xc0\x0b\x0f\x00\x3d\x3c\x00\xf0\xf0\x00\xd0\xc3\x03\xd0\x07\xff\xff\x07\xfc\xff\x2f\xf0\x00\xf4\xc2\x03\x00\x1f\x0f\x00\xbc\x3c\x00\xe0\xf2\x00\xc0\xc7\x03\xc0\x0f\xff\xff\x0f\xfc\xff\x0b\x00" +

	// glyph 1717 for rune 'C'
	"\x10\xee\x00\x01\x0f\x00\x00\x01\x00\x00\xfe\x1f\x00\xf4\xab\x3f\x00\x2f\x00\x2f\xf0\x02\x40\x0f\x3c\x00\xc0\x47\x0f\x00\x50\xe0\x03\x00\x00\xb8\x00\x00\x40\xe1\x03\x00\x00\xf4\x00\x00\x05\x3c\x00\xc0\x07\x2e\x00\xf4\x00\x2f\x00\x1f\x40\xbf\xfe\x02\x00\xfe\x1f\x00" +

	// glyph 1783 for rune 'D'
	"\x10\xef\x00\x02\x0f\xfc\xff\x02\xc0\xff\xff\x01\x3c\x00\xbe\xc0\x03\x40\x0f\x3c\x00\xe0\xc2\x03\x00\x3d\x3c\x00\xc0\xc3\x03\x00\x7c\xc5\x03\x00\x3c\x3c\x00\xd0\xc3\x03\x00\x2f\x3c\x00\xf4\xc0\x03\xe0\x07\xfc\xff\x1f\xc0\xff\x1f\x00" +

	// glyph 1841 for rune 'E'
	"\x0e\xef\x00\x02\x0d\xfc\xff\xbf\xf1\x00\x00\x54\xf1\xab\x6a\xf0\xff\xbf\xf0\x01\x00\xf0\x00\x00\x54\xf1\xff\xff\xf2\xff\xff\x03" +

	// glyph 1873 for rune 'F'
	"\x0d\xef\x00\x02\x0d\xfc\xff\xbf\xfc\xff\x7f\x3c\x00\x00\x55\xbc\x55\x05\xfc\xff\x1f\xbc\x55\x05\x3c\x00\x00\x55\x05" +

	// glyph 1902 for rune 'G'
	"\x10\xee\x00\x01\x0f\x00\x00\x01\x00\x40\xfe\x2f\x00\xf4\xeb\x3f\x00\x2f\x00\x2f\xf0\x02\x00\x0f\x7c\x00\xc0\x47\x0f\x00\x00\x84\x0b\x00\x00\xe0\x02\xa4\x6a\xb8\x00\xfe\x2f\x3d\x00\xd5\x4b\x0f\x00\xe0\xc2\x07\x00\xb8\xe0\x03\x00\x2e\xf0\x03\xd0\x0b\xf0\xaf\xff\x00\xe0\xff\x07" +

	// glyph 1971 for rune 'H'
	"\x11\xef\x00\x02\x0f\x3c\x00\x00\x5f\x55\xfc\xaa\xea\xcf\xff\xff\xff\x7c\x00\x40\xcf\x03\x00\xf0\x55\x05" +

	// glyph 1997 for rune 'I'
	"\x07\xef\x00\x02\x05\x7c\x55\x55\x55\x55" +

	// glyph 2007 for rune 'J'
	"\x0d\xef\x00\x01\x0c\x00\x00\x7c\x55\x55\x15\x0a\x00\x1f\x0f\x00\x0f\x2f\xc0\x0f\xfd\xfa\x03\xf0\xbf\x00" +

	// glyph 2033 for rune 'K'
	"\x0f\xef\x00\x02\x0f\x3c\x00\xe0\xc3\x03\x80\x0f\x3c\x00\x7d\xc0\x03\xf4\x01\x3c\xd0\x07\xc0\x03\x2f\x00\x3c\xbc\x00\xc0\xf7\x03\x00\xfc\xbf\x00\xc0\x7f\x1f\x00\xfc\xe0\x03\xc0\x07\xfc\x00\x3c\x00\x2f\xc0\x03\xd0\x07\x3c\x00\xf8\xc0\x03\x00\x3f\x3c\x00\xc0\x0b" +

	// glyph 2098 for rune 'L'
	"\x0d\xef\x00\x02\x0d\x7c\x00\x00\x55\x55\x55\xc5\xff\xff\x17" +

	// glyph 2113 for rune 'M'
	"\x15\xef\x00\x02\x13\xfc\x00\x00\xc0\xcf\x2f\x00\x00\xfe\xfc\x03\x00\xf0\xcf\x7f\x00\x40\xff\xbc\x0b\x00\xbc\xcf\xf3\x00\xc0\xf3\x3c\x2e\x00\x2e\xcf\xd3\x03\xf0\xf0\x3c\x7c\x40\x0f\xcf\x43\x0b\x7c\xf4\x3c\xf0\xc0\x43\xcf\x03\x2e\x2e\xf4\x3c\xc0\xf3\x40\xcf\x03\xbc\x0f\xf4\x3c\x40\x7f\x40\xcf\x03\xf0\x03\xf4\x3c\x00\x2e\x40\x0f" +

	// glyph 2195 for rune 'N'
	"\x11\xef\x00\x02\x0f\xbc\x00\x00\xcf\x1f\x00\xf0\xfc\x03\x00\xcf\xbf\x00\xf0\xbc\x0f\x00\xcf\xe7\x03\xf0\x7c\xbc\x00\xcf\x47\x0f\xf0\x7c\xe0\x03\xcf\x07\x7c\xf0\x7c\x40\x0f\xcf\x07\xf0\xf3\x7c\x00\xbc\xcf\x07\x40\xff\x7c\x00\xf0\xcf\x07\x00\xfc\x7c\x00\x80\x0f" +

	// glyph 2260 for rune 'O'
	"\x11\xee\x00\x01\x0f\x00\x00\x01\x00\x00\xfe\x1f\x00\xf4\xef\x2f\x00\x2f\x40\x2f\xe0\x03\x40\x0f\x7c\x00\xc0\x4b\x0f\x00\xd0\x87\x0b\x00\xc0\x17\x3d\x00\x40\x1f\x7c\x00\xc0\x0b\x3e\x00\xf4\x00\x3f\x40\x2f\x00\xff\xfe\x02\x00\xfe\x1f\x00" +

	// glyph 2319 for rune 'P'
	"\x0f\xef\x00\x02\x0e\xfc\xff\x1b\xf0\xff\xff\xc2\x03\x40\x1f\x0f\x00\xf4\x3c\x00\xc0\xc7\x03\x00\x3d\x0f\x00\xbc\xbc\xaa\xfe\xf0\xff\xbf\xc0\x57\x05\x00\x0f\x00\x00\x55\x01" +

	// glyph 2362 for rune 'Q'
	"\x11\xef\x03\x01\x0f\x00\xf8\x7f\x00\xd0\xbf\xbf\x00\xbc\x00\xbd\xc0\x0b\x00\x3d\xf0\x00\x00\x1f\x3d\x00\x80\x8f\x0b\x00\xd0\x87\x0b\x00\xc0\xe3\x02\x00\xf4\xd1\x03\x00\xf8\xf0\x00\x00\x1f\xbc\x00\xd0\x03\xbc\x00\xbd\x00\xfd\xfb\x0b\x00\xf8\xff\x02\x00\x00\xf4\x03\x00\x00\xf4\x03\x00\x00\x20" +

	// glyph 2435 for rune 'R'
	"\x0f\xef\x00\x02\x0e\xfc\xff\x0b\xf0\xff\xff\xc0\x03\xc0\x0f\x0f\x00\x7c\x3c\x00\xf0\xf2\x00\x80\xcb\x03\x00\x1f\x0f\x00\x3f\xfc\xff\x3f\xf0\xff\x2f\xc0\x07\xf8\x00\x0f\xc0\x07\x3c\x00\x3d\xf0\x00\xf0\xc1\x03\x40\x0f\x0f\x00\xbc\x3c\x00\xd0\x03" +

	// glyph 2496 for rune 'S'
	"\x0e\xee\x00\x01\x0e\x00\x40\x00\x00\xe0\xff\x02\xd0\xaf\xfe\x00\x2f\x00\x2f\xf4\x00\xd0\x43\x0f\x00\x3c\xf4\x00\x00\x00\x7f\x00\x00\xc0\xbf\x01\x00\xd0\xff\x02\x00\x80\xff\x00\x00\x40\x3f\x00\x00\xd0\xc3\x07\x00\x7c\xbc\x00\xc0\x43\x1f\x00\x3e\xd0\xaf\xfe\x00\xe0\xff\x02" +

	// glyph 2564 for rune 'T'
	"\x0e\xef\x00\x00\x0e\xf4\xff\xff\x3f\xfd\xff\xff\x0b\x00\x7c\x00\x54\x55\x55\x15" +

	// glyph 2584 for rune 'U'
	"\x10\xef\x00\x01\x0e\xf4\x00\x40\x5f\x55\x55\xc1\x03\x00\x3d\x7c\x00\xe0\x82\x1f\x80\x0f\xf0\xaf\x7f\x00\xf4\xbf\x00" +

	// glyph 2613 for rune 'V'
	"\x0f\xef\x00\x00\x0f\xf4\x00\x00\xf8\xf0\x01\x00\x7c\xe0\x02\x00\x3d\xd0\x03\x00\x2e\xc0\x07\x00\x1f\x80\x0b\x40\x0f\x40\x0f\xc0\x0b\x00\x1f\xc0\x03\x00\x2e\xd0\x03\x00\x3c\xf0\x01\x00\x7c\xf0\x00\x00\xf4\xb4\x00\x00\xf0\x7c\x00\x00\xe0\x3e\x00\x00\xd0\x2f\x00\x00\xc0\x0f\x00\x00\x80\x0f\x00" +

	// glyph 2686 for rune 'W'
	"\x15\xef\x00\x01\x15\x3c\x00\xf4\x00\xe0\xf2\x01\xe0\x03\xc0\x87\x0b\xc0\x1f\x00\x0f\x3d\x00\xbf\x00\x3d\xf0\x00\xed\x03\xb8\xc0\x03\x78\x1e\xe0\x01\x1e\xf0\xb4\xc0\x03\xb4\xd0\xc2\x03\x0f\xc0\x83\x07\x0f\x2d\x00\x0f\x0f\x78\x78\x00\x78\x3c\xd0\xf2\x00\xe0\xb6\x00\xcf\x03\x40\xef\x00\x78\x0b\x00\xfc\x03\xd0\x1f\x00\xf0\x0b\x00\x7f\x00\x80\x1f\x00\xfc\x00\x00\x3d\x00\xe0\x03\x00" +

	// glyph 2781 for rune 'X'
	"\x0f\xef\x00\x01\x0e\xfc\x00\x80\x0f\x1f\x00\x7d\xe0\x03\xf0\x02\xbc\x80\x0f\x40\x1f\x7c\x00\xe0\xf3\x02\x00\xfc\x0f\x00\x40\x7f\x00\x00\xf0\x03\x00\x80\xbf\x00\x00\xbc\x0f\x00\xf0\xe3\x03\x40\x0f\xbc\x00\xbc\x40\x0f\xe0\x03\xf0\x43\x1f\x00\x7c\xbc\x00\x80\x0f" +

	// glyph 2846 for rune 'Y'
	"\x0e\xef\x00\x00\x0e\xf8\x00\x00\x3e\xbc\x00\xc0\x07\x3d\x00\xf8\x00\x2f\x00\x0f\x40\x0f\xe0\x02\xc0\x0b\x3c\x00\xc0\xc3\x0b\x00\xe0\xf6\x00\x00\xf0\x1f\x00\x00\xf8\x03\x00\x00\x7c\x00\x54\x15" +

	// glyph 2894 for rune 'Z'
	"\x0e\xef\x00\x01\x0e\xfc\xff\xff\x83\xff\xff\x3f\x00\x00\xf0\x01\x00\xc0\x0b\x00\x00\x3e\x00\x00\xf0\x01\x00\xc0\x0b\x00\x00\x3e\x00\x00\xf4\x01\x00\xc0\x0b\x00\x00\x3e\x00\x00\xf4\x00\x00\xc0\x0b\x00\x00\x3e\x00\x00\xf4\x00\x00\xc0\xff\xff\x7f\x01" +

	// glyph 2956 for rune '['
	"\x06\xec\x04\x01\x06\xa0\x4a\xff\xf4\x45\x0f\x55\x55\x55\x55\x45\xff\xa0\x0a" +

	// glyph 2975 for rune '\\'
	"\x0a\xef\x02\x00\x0a\xb4\x00\x00\x3c\x00\x00\x2e\x00\x00\x0f\x00\xc0\x07\x00\xd0\x02\x00\xf0\x00\x00\x78\x00\x00\x3d\x00\x00\x0f\x00\x40\x0b\x00\xc0\x03\x00\xe0\x01\x00\xb4\x00\x00\x3c\x00\x00\x1e\x00\x00\x0f\x00\xc0\x07\x00\x40\x01" +

	// glyph 3033 for rune ']'
	"\x06\xec\x04\x00\x05\xa4\xc6\xbf\x94\x0b\xb8\x55\x55\x55\x55\xc5\xbf\xa8\x06" +

	// glyph 3052 for rune '^'
	"\x0a\xef\xf8\x01\x09\x00\x0f\x00\x7d\x00\xfc\x03\xb0\x0e\xe0\xb5\xc0\xc3\x43\x0b\x1e\x0f\xf0\x14\x40\x01" +

	// glyph 3078 for rune '_'
	"\x0b\x00\x02\x00\x0b\xfc\xff\xff\xf8\xff\xbf" +

	// glyph 3089 for rune '`'
	"\x07\xee\xf2\x01\x05\xbc\xc0\x07\x3c\x40" +

	// glyph 3099 for rune 'a'
	"\x0d\xf3\x00\x01\x0c\x00\xfa\x02\xd0\xff\x0f\xf0\x01\x3e\xb4\x00\x3c\x00\x00\x7c\x00\xa9\x7e\xd0\xff\x7f\xf0\x02\x7c\xb8\x00\x7c\xe1\x03\xf8\xc1\xef\xff\x01\xff\xe3\x02" +

	// glyph 3141 for rune 'b'
	"\x0d\xee\x00\x01\x0d\xf4\x00\x00\x54\xd1\x93\x1b\x40\xff\xff\x02\xfd\x40\x1f\xf4\x00\xf4\xd0\x03\xc0\x13\x3d\x00\x7c\xd1\x03\xc0\x43\x0f\x40\x0f\xbd\x00\x2f\xf4\xaf\x3f\xd0\xf2\x2f\x00" +

	// glyph 3187 for rune 'c'
	"\x0d\xf3\x00\x01\x0c\x00\xfd\x02\xc0\xff\x1f\xf0\x03\x3d\xf4\x00\x78\xb8\x00\x60\x7c\x00\x00\x85\x07\x00\x40\x0f\x40\x0b\x1f\xc0\x07\xfd\xfa\x02\xf0\x7f\x00" +

	// glyph 3226 for rune 'd'
	"\x0e\xee\x00\x01\x0c\x00\x00\xf4\x55\x00\xbe\xf5\xc0\xff\xff\xf0\x03\xfe\xf4\x00\xf8\xb8\x00\xf4\x78\x00\xf4\x7c\x00\xf4\xe1\x01\xd0\xd3\x03\xd0\xc3\x0b\xf0\x43\xbf\xff\x03\xfd\xcb\x03" +

	// glyph 3272 for rune 'e'
	"\x0d\xf3\x00\x01\x0c\x00\xfd\x02\xc0\xff\x0f\xe0\x03\x3e\xf4\x00\x7c\xb8\x00\xb4\xbc\x55\xb9\xfc\xff\xff\xbc\x55\x15\x78\x00\x00\xf4\x00\x00\xf0\x02\x74\xd0\xaf\x3f\x00\xfe\x0b" +

	// glyph 3316 for rune 'f'
	"\x08\xee\x00\x00\x09\x00\xf8\x07\xe0\x2b\x00\x1f\x00\xf0\x00\x40\x0f\x00\xfe\x1a\xf4\xff\x02\xf4\x00\x55\x55\x05" +

	// glyph 3344 for rune 'g'
	"\x0d\xf3\x05\x01\x0c\x00\xbe\xa1\xc0\xff\xfb\xf0\x03\xfe\xf4\x00\xf8\xb8\x00\xf4\x78\x00\xf4\x7c\x00\xf4\xe1\x02\xd0\xd3\x03\xd0\xc3\x0b\xf0\x43\xbf\xff\x03\xfd\xdb\x03\x00\xd0\x02\x00\xe0\xc2\x07\xf8\x80\xff\x7f\x00\xf8\x0b\x00" +

	// glyph 3401 for rune 'h'
	"\x0d\xee\x00\x01\x0c\xf4\x00\x00\x55\xf4\xe4\x07\xf4\xfe\x2f\xf4\x07\x3d\xf4\x00\x78\xf4\x00\xb8\x55\x55" +

	// glyph 3427 for rune 'i'
	"\x06\xef\x00\x02\x04\x7c\x40\x28\x5f\x55\x55\x01" +

	// glyph 3439 for rune 'j'
	"\x06\xef\x05\xff\x04\x40\x1f\x00\x04\xa0\x40\x5f\x55\x55\x15\xe0\xe2\x1f\xbe\x00" +

	// glyph 3459 for rune 'k'
	"\x0c\xee\x00\x01\x0c\xf4\x00\x00\x55\xf4\x00\x29\xf4\x00\x1f\xf4\xc0\x0b\xf4\xf0\x02\xf4\xbc\x00\xf4\x3f\x00\xf4\x7f\x00\xf4\xfb\x00\xf4\xf0\x03\xf4\xc0\x0b\xf4\x40\x1f\xf4\x00\x3e\xf4\x00\xfc" +

	// glyph 3507 for rune 'l'
	"\x06\xee\x00\x02\x04\x7c\x55\x55\x55\x55" +

	// glyph 3517 for rune 'm'
	"\x15\xf3\x00\x01\x14\x64\xe4\x07\xf8\x06\xf4\xff\x2f\xff\x1f\xf4\x03\xfe\x02\x3e\xf4\x00\xfc\x00\x7c\xf4\x00\xb8\x00\x7c\x55\x55" +

	// glyph 3549 for rune 'n'
	"\x0d\xf3\x00\x01\x0c\x60\xe4\x07\xf4\xfe\x2f\xf4\x07\x3d\xf4\x00\x78\xf4\x00\xb8\x55\x55" +

	// glyph 3571 for rune 'o'
	"\x0e\xf3\x00\x01\x0d\x00\xf9\x02\x00\xff\x7f\x00\x3e\xd0\x07\x3d\x00\x3d\xb8\x00\xf0\xf1\x01\x80\xc7\x07\x00\x6e\x78\x00\xf0\xd1\x03\xc0\x03\x2f\xc0\x0b\xf0\xeb\x0f\x00\xfe\x0b\x00" +

	// glyph 3616 for rune 'p'
	"\x0d\xf3\x05\x01\x0d\x60\xe8\x06\xd0\xff\xbf\x40\x3f\xe0\x07\x3d\x00\x3e\xf4\x00\xf0\x44\x0f\x00\x5f\xf4\x00\xf0\xd0\x03\xd0\x43\x2f\xc0\x0b\xfd\xeb\x0f\xf4\xfd\x0b\xd0\x03\x00\x50\xd1\x02\x00\x00" +

	// glyph 3665 for rune 'q'
	"\x0e\xf3\x05\x01\x0c\x00\xbe\xa1\xc0\xff\xfb\xf0\x03\xfd\xf4\x00\xf8\xb8\x00\xf4\x7c\x00\xf4\x85\x07\x40\x4f\x0f\x40\x0f\x2f\xc0\x0f\xfd\xfa\x0f\xf4\x6f\x0f\x00\x40\x5f\x01\x00\xd0\x02" +

	// glyph 3711 for rune 'r'
	"\x08\xf3\x00\x01\x08\xa0\xb8\xf4\xff\xf4\x17\xf4\x00\x55\x55\x01" +

	// glyph 3727 for rune 's'
	"\x0c\xf3\x00\x01\x0b\x00\xbe\x01\xf4\xff\x03\x1f\xf0\xd2\x03\xf0\xf4\x00\x00\xfc\x06\x00\xf8\x2f\x00\x90\x7f\x00\x00\x3d\x1e\x00\x8f\x0f\xd0\xc3\xaf\x7e\x80\xff\x07" +

	// glyph 3768 for rune 't'
	"\x08\xf0\x00\x00\x07\x40\x07\x80\x0b\xa1\xbf\xf2\xff\x03\x2e\x54\x55\x40\x0f\x40\xaf\x00\xfe" +

	// glyph 3791 for rune 'u'
	"\x0d\xf3\x00\x01\x0c\xa4\x00\x64\xf4\x00\xb8\x55\x55\xf0\x01\xbd\xe0\xef\xbf\x80\xff\xb9" +

	// glyph 3813 for rune 'v'
	"\x0c\xf3\x00\x00\x0b\x64\x00\xa0\xf0\x00\xf4\xf0\x01\x78\xd0\x02\x3c\xc0\x03\x2d\x80\x07\x1e\x40\x0b\x0f\x00\x4f\x0b\x00\x9e\x07\x00\xed\x03\x00\xfc\x02\x00\xf8\x00\x00\xf0\x00" +

	// glyph 3857 for rune 'w'
	"\x12\xf3\x00\x00\x12\xa4\x00\x28\x00\x1a\x3c\x40\x0f\xc0\x03\x0f\xe0\x0b\xf0\x80\x07\xfc\x03\x2d\xd0\x03\xeb\x80\x07\xf0\xe0\x75\xf0\x00\x78\x3c\x3c\x2c\x00\x2d\x0b\x4e\x07\x00\xdf\x41\xe7\x00\xc0\x3f\xc0\x3e\x00\xe0\x0b\xf0\x0b\x00\xf4\x01\xf4\x00\x00\x3c\x00\x3c\x00" +

	// glyph 3924 for rune 'x'
	"\x0c\xf3\x00\x01\x0b\x28\x00\x29\x3e\xc0\x07\x1f\xf8\x40\x0f\x0f\x80\xfb\x01\xc0\x3f\x00\xd0\x07\x00\xfc\x02\x80\xff\x01\xf0\xf4\x00\x2f\xb8\xe0\x03\x7c\x7c\x00\x3d" +

	// glyph 3965 for rune 'y'
	"\x0b\xf3\x05\x00\x0b\xa4\x00\xa0\xf4\x00\xb8\xf0\x01\x7c\xe0\x02\x3c\xc0\x03\x2e\xc0\x07\x1f\x80\x0b\x0f\x00\x4f\x0b\x00\xdf\x03\x00\xfd\x03\x00\xfc\x02\x00\xf8\x00\x00\xf4\x00\x00\x74\x00\x00\x3c\x00\x00\x2d\x00\xf0\x0f\x00\xf0\x03\x00" +

	// glyph 4024 for rune 'z'
	"\x0c\xf3\x00\x01\x0b\xa8\xaa\x2a\xff\xff\x0b\x00\xf4\x01\x00\x2f\x00\xf0\x03\x00\x3e\x00\xd0\x07\x00\xbc\x00\xc0\x0b\x00\xf8\x00\x40\x1f\x00\xf0\xff\xff\x01" +

	// glyph 4063 for rune '{'
	"\x08\xed\x04\x01\x08\x00\x50\x00\xbc\x00\x2f\x40\x0f\x80\x0b\x80\x07\xc0\x07\x05\x3c\x00\x3e\xc0\x0f\xc0\x1f\x00\x3d\x00\x7c\x50\x01\x1e\x00\x2e\x00\x3c\x00\xf8\x01\xd0\x02" +

	// glyph 4106 for rune '|'
	"\x06\xef\x03\x02\x04\x7c\x55\x55\x55\x55\x05" +

	// glyph 4117 for rune '}'
	"\x08\xed\x04\x00\x08\x14\x00\xe0\x03\x00\x3d\x00\xf0\x01\x40\x0b\x00\x3d\x50\x01\x3c\x00\xe0\x07\x00\x7e\x00\xfc\x00\x7c\x00\xf0\x00\xd0\x03\x15\xe0\x02\xc0\x03\xc0\x0b\x80\x07\x00" +

	// glyph 4162 for rune '~'
	"\x10\xf6\xfb\x01\x0f\x00\x04\x00\x00\xf0\x3f\x00\x0b\xbf\x7f\xd0\xc2\x02\xbe\x3e\x74\x00\xfd\x03" +

	"")