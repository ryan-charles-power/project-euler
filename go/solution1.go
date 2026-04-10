package solutions

import (
	"fmt"
)

func Solve1() {
	sum := 0
	for i := 1; i < 1000; i++ {
		if (i%3 == 0) || (i%5 == 0) {
			sum += i
		}
	}

	fmt.Printf("%d is the sum of all multiples of 3 or 5", sum)
}
