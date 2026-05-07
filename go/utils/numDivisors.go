package utils

import "math"

func NumDivisors(n int) int {
	count := 0
	limit := int(math.Sqrt(float64(n)))

	for i := 1; i <= limit; i++ {
		if n%i == 0 {
			if i*i == n {
				count++ // Perfect square (e.g., 10*10=100), only count once
			} else {
				count += 2 // Count both i and n/i
			}
		}
	}
	return count
}
