package utils

import "math"

func IsPentagonalNumber(x int) bool {
	if x <= 0 {
		return false
	}

	// Calculate n using floating-point math
	val := (math.Sqrt(float64(24*x+1)) + 1) / 6

	// Check if val is a whole number (integer)
	return val == math.Trunc(val)
}
