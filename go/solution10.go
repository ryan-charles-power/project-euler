package solutions

import (
	"fmt"
	"project-euler/go/utils"
)

func Solve10() {
	sum := 0

	for i := 2; i < 2000000; i++ {
		if utils.IsPrimeNumber(i) {
			sum += i
		}
	}

	fmt.Printf("The sum of all prime numbers below 2 million is %d", sum)

}
