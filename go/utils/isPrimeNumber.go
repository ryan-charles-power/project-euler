package utils

import "math"

func IsPrimeNumber(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	// Eliminate even numbers and multiples of 3 immediately
	if n%2 == 0 || n%3 == 0 {
		return false
	}

	// Check divisibility starting from 5, skipping even numbers
	// Every prime > 3 is of the form 6k ± 1
	limit := int(math.Sqrt(float64(n)))
	for i := 5; i <= limit; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}
