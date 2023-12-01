package gfx

import (
	"math"
	"testing"
)

func TestIsqrt(t *testing.T) {
	// Test isqrt, either exhaustively or just a limited subset of smaller
	// numbers.
	//const until = 1 << 32
	const until = 1 << 20
	for i := uint64(0); i < until; i++ {
		result := uint64(isqrt32(uint32(i)))
		expected := uint64(math.Sqrt(float64(i)))
		if result != expected {
			t.Errorf("isqrt(%d): %d %.2f", i, result, math.Sqrt(float64(i)))
			break
		}
	}
}
