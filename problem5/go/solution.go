package main

import "fmt"

func main() {
	finished := false
	i := 20
	var result int

	for !finished {
		isTarget := true
		for j := 1; j <= 20; j++ {
			if i%j != 0 {
				isTarget = false
				break
			}
		}

		if isTarget {
			result = i
			finished = true
		} else {
			i += 20 // result has to be a multiple of 20
		}
	}

	fmt.Println(result)
}
