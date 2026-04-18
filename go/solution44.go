package solutions

import (
	"fmt"
	utils "project-euler/go/utils"
)

func Solve44() {
	var P []int

	solved := false
	n := 1
	var solution int

	for !solved {
		Pn := n * (3*n - 1) / 2
		P = append(P, Pn)

		if len(P) > 1 {
			// Check all combinations to see if both sum and dif are Pentagon Numbers
			for i := len(P) - 2; i >= 0; i-- { // Iterating backwards finds the smallest D faster
				Pj := P[i]
				Psum := Pn + Pj
				Pdif := Pn - Pj

				if utils.IsPentagonalNumber(Psum) && utils.IsPentagonalNumber(Pdif) {
					solution = Pdif
					solved = true
					break
				}
			}

		}

		n++
	}

	fmt.Printf("The solution for D = |Pk - Pj| for problem 44 is %d", solution)
}
