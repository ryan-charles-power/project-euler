package solutions

import (
	"fmt"
)

func Solve6() {
	var sum_x, sum_y int

	for i := 1; i <= 100; i++{
		sum_x += i
		sum_y += i*i
	}

	sum_x = sum_x*sum_x
	solution := sum_x - sum_y

	fmt.Printf("sum_x: %d, sum_y: %d\nsolution: %d", sum_x, sum_y, solution)
}