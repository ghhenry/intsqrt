package intsqrt

import (
	"math"
)

// Sqrt32 returns floor(sqrt(n)).
func Sqrt32(n uint32) uint32 {
	return uint32(math.Sqrt(float64(n)))
}

// Sqrt64 returns floor(sqrt(n)).
func Sqrt64(n uint64) uint64 {
	if n < 1<<52 {
		return uint64(math.Sqrt(float64(n)))
	}
	var y = 1 + uint64(math.Sqrt(float64(n)))
	var x = y + 1
	for y < x {
		x = y
		y = (n/x + x) / 2
	}
	return x
}
