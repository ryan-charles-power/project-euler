package solutions

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func Solve8() {
	// Read the entire file content into a byte slice.
	fileContent, err := os.ReadFile("../project-euler/resources/problem8number.txt")
	if err != nil {
		log.Fatal(err) // Handle potential errors, such as file not found.
	}

	// Convert the byte slice to a string.
	numStr := string(fileContent)
	lengthStr := len(numStr)

	numProducts := 13
	maxProduct := 0

	for i := 0; i < lengthStr-numProducts; i++ {
		product := 1
		//
		for j := 0; j < numProducts; j++ {
			num, err := strconv.Atoi(string(numStr[i+j]))
			if err != nil {
				fmt.Printf("\nERROR: %s\n", err.Error())
			}
			product *= num
		}

		if product > maxProduct {
			maxProduct = product
		}
	}

	fmt.Printf("The max product of %d numbers in a row in the provided string is %d", numProducts, maxProduct)

}
