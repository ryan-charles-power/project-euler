package main

import (
	"fmt"
	"strconv"
)

func main() {

	i := 1
	j := 1
	sum := 0
	var temp int

	for j < 4000000 {
		if j%2 == 0 {
			sum += j
		}
		temp = j
		j = i + j
		i = temp
	}

	fmt.Println("The sum of all even fibonacci numbers less than 4,000,000 is " + strconv.Itoa(sum))
}
