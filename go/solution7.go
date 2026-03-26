package solutions

import (
	"fmt"
	utils "project-euler/go/utils"
	"strconv"
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

	fmt.Println("The 10001st prime number is: " + strconv.Itoa(primeNumber))
}
