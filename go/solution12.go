package solutions

import (
	"fmt"
	"project-euler/go/utils"
)

func Solve12() {
	solved := false
	i := 0
	triangle_i := 0
	var divisors int
	for !solved {
		i++
		triangle_i += i
		divisors = utils.NumDivisors(triangle_i)

		if divisors > 500 {
			solved = true
		}
	}

	fmt.Printf("The number of divisors for %d is %d", triangle_i, divisors)
}
