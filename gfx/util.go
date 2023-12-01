package gfx

import "math/bits"

// Get the square root of an integer. The result is always exactly the square
// root of the input number rounded down.
func isqrt32(n uint32) uint32 {
	// Accurate and (hopefully) fast integer square root for slow chips like the
	// Cortex-M0+. Source:
	// https://stackoverflow.com/questions/34187171/fast-integer-square-root-approximation/63452286#63452286
	// Changed the code a bit:
	//   * assume 32-bit integers everywhere
	//   * use `32 - clz(n)` instead of `64 - clz(n)`
	//   * move `shift -= 2` to the check at the end, to avoid a compare
	//     instruction (`shift -= 2` already sets the compare bits)
	shift := 32 - bits.LeadingZeros32(n)
	shift += shift & 1 // round up to next multiple of 2

	result := uint32(0)
	for {
		result <<= 1                            // leftshift the result to make the next guess
		result |= 1                             // guess that the next bit is 1
		if result*result > (n >> uint(shift)) { // revert if guess too high
			result ^= 1
		}
		shift -= 2
		if shift < 0 {
			break
		}
	}

	return result
}
