package solutions

import (
	"fmt"
	"strconv"
)

func Solve100() {
	numDisks := float64(1000000000000) // 10^12
	numBlueDisks := float64(500000000000)

	finished := false
	for finished == false {
		fmt.Println("numDisks: " + strconv.Itoa(int(numDisks)))
		for numBlueDisks < numDisks {
			probabilityBB := float32((numBlueDisks / numDisks) * ((numBlueDisks - 1) / (numDisks - 1)))
			fmt.Printf("numBlueDisks: %.0f\tprobabilityBB: %.8f\n", numBlueDisks, probabilityBB)

			if probabilityBB == 0.5 {
				fmt.Println("A probability of BB with exactly 50 percent can be found with " + strconv.Itoa(int(numBlueDisks)) + " blue disks")
				finished = true
				break
			}

			if probabilityBB < 0.5 {

			}

			if probabilityBB > 0.5 {

			}

			if numBlueDisks == numDisks {
				fmt.Println("A probability of BB with exactly 50 percent was not found")
				break
			}

			numBlueDisks++
		}

		numBlueDisks = 2
		numDisks++
	}
}
