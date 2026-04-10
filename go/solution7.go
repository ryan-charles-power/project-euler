package solutions

import (
	"fmt"
	utils "project-euler/go/utils"
)

func Solve7() {
	i := 0
	j := 1
	primeNumber := 0

	for i < 10001 {
		if utils.IsPrimeNumber(j) {
			primeNumber = j
			i++
		}

		j++
	}

	fmt.Printf("The 10001st prime number is: %d", primeNumber)
}
