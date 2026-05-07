package solutions

import (
	"fmt"
	"math"
)

func Solve12() {
	solved := false
	i := 0
	triangle_i := 0
	var divisors int
	for !solved {
		i++
		triangle_i += i
		divisors = numDivisors(triangle_i)

		if divisors > 500 {
			solved = true
		}
	}

	fmt.Printf("The number of divisors for %d is %d", triangle_i, divisors)
}

func numDivisors(n int) int {
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
