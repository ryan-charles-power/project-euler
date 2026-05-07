package solutions

import (
	"fmt"
)

func Solve9() {

	product := 0
	var final_a, final_b, final_c int
	productFound := false

	// Initial bounds set of a < b < c
	c := 3
	for !productFound {
		for b := 2; b < c; b++ {
			for a := 1; a < b; a++ {
				if (a*a)+(b*b) == (c * c) {
					if a+b+c == 1000 {
						product = a * b * c
						final_a = a
						final_b = b
						final_c = c
						productFound = true
					}
				}
			}
		}
		c++
	}

	fmt.Printf("The product of abc, where a < b < c AND a^2 + b^2 = c^2 AND a + b + c = 1000 is %d\na = %d, b= %d, c = %d", product, final_a, final_b, final_c)

}
