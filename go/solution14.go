package solutions

import (
	"fmt"
)

func Solve14() {
	maxLength := 0
	startingNum := 1

	for i := 1; i < 1000000; i++ {
		currentLength := 1
		currentNum := i

		for currentNum != 1 {
			currentLength++
			if currentNum%2 == 0 {
				currentNum = currentNum / 2
			} else {
				currentNum = 3*currentNum + 1
			}
		}

		if currentLength > maxLength {
			maxLength = currentLength
			startingNum = i
			fmt.Printf("maxLength = %d, n = %d\n", maxLength, startingNum)
		}

	}

	fmt.Printf("The longest chain starting where n is under 1,000,000 is %d long, with a starting number of %d", maxLength, startingNum)
}
