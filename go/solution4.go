package solutions

import (
	"fmt"
	utils "project-euler/go/utils"
	"strconv"
)

func Solve4() {
	product := 0
	largestPalindrome := 0
	iMax := 0
	jMax := 0

	// run through all products 1-999, and 1-999
	for i := 1; i < 1000; i++ {
		for j := 1; j < 1000; j++ {
			product = i * j
			// Check size first to save time
			if product > largestPalindrome {
				// Then do string conversion and comparison
				productStr := strconv.Itoa(product)
				productStrReverse := utils.ReverseString(productStr)
				if productStr == productStrReverse {
					largestPalindrome = product
					iMax = i
					jMax = j
				}
			}
		}
	}

	fmt.Printf("The largest palindrome is %d, which is a product of %d and %d", largestPalindrome, iMax, jMax)
}
