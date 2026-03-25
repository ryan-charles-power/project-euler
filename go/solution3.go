package solutions

import (
	"fmt"
)

func Solve3() {
	var num int64 = 600851475143
	original := num
	var divisor int64 = 2

	// Loop until divisor squared exceeds the remaining number
	for divisor^2 <= num {
		// If the divisor is a factor
		if num%divisor == 0 {
			// Divide out all occurrences of this factor
			for num%divisor == 0 {
				num /= divisor
			}
		}
		divisor++
	}

	// If num > 1, the remaining num is the largest prime factor
	largest := num
	if num == 1 {
		largest = divisor - 1
	}

	fmt.Printf("The largest prime factor of %d is %d\n", original, largest)
}
