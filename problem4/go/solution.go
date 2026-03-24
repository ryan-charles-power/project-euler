package main

import (
	"fmt"
	"strconv"
)

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
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
				productStrReverse := Reverse(productStr)
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
