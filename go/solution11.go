package solutions

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Solve11() {
	// Read the entire file content into a byte slice.
	fileContent, err := os.ReadFile("../project-euler/resources/problem11grid.txt")
	if err != nil {
		log.Fatal(err) // Handle potential errors, such as file not found.
	}

	// Convert the byte slice to a string.
	gridStr := string(fileContent)

	// 1. Split into lines (rows)
	lines := strings.Split(strings.TrimSpace(gridStr), "\n")

	// Create the 2D slice
	var matrix [][]int

	for _, line := range lines {
		// Split each line by whitespace
		fields := strings.Fields(line)
		var row []int

		for _, field := range fields {
			// Convert string to integer
			num, _ := strconv.Atoi(field)
			row = append(row, num)
		}
		// Add the completed row to the matrix
		matrix = append(matrix, row)
	}

	maxProduct := 0
	product := 0
	var maxProducts [4]int
	var products [4]int

	// i = x, j = y
	for i := range 19 {
		for j := range 19 {
			if i > 2 && j > 2 {
				product, products = sumNorthWest(matrix, i, j)
				maxProduct, maxProducts = compareProducts(product, maxProduct, products, maxProducts)
			}
			if i < 17 && j > 2 {
				product, products = sumNorthEast(matrix, i, j)
				maxProduct, maxProducts = compareProducts(product, maxProduct, products, maxProducts)
			}
			if i < 17 {
				product, products = sumEast(matrix, i, j)
				maxProduct, maxProducts = compareProducts(product, maxProduct, products, maxProducts)
			}
			if j > 2 {
				product, products = sumNorth(matrix, i, j)
				maxProduct, maxProducts = compareProducts(product, maxProduct, products, maxProducts)
			}
		}
	}

	fmt.Printf("The max product is %d made of the products %d", maxProduct, maxProducts)
}

func compareProducts(product int, maxProduct int, products [4]int, maxProducts [4]int) (int, [4]int) {
	if product > maxProduct {
		return product, products
	}
	return maxProduct, maxProducts
}

func sumNorth(matrix [][]int, i int, j int) (int, [4]int) {
	factor1 := matrix[i][j]
	factor2 := matrix[i][j-1]
	factor3 := matrix[i][j-2]
	factor4 := matrix[i][j-3]
	product := factor1 * factor2 * factor3 * factor4

	factors := [4]int{factor1, factor2, factor3, factor4}
	return product, factors
}

func sumEast(matrix [][]int, i int, j int) (int, [4]int) {
	factor1 := matrix[i][j]
	factor2 := matrix[i+1][j]
	factor3 := matrix[i+2][j]
	factor4 := matrix[i+3][j]
	product := factor1 * factor2 * factor3 * factor4

	factors := [4]int{factor1, factor2, factor3, factor4}
	return product, factors
}

func sumNorthEast(matrix [][]int, i int, j int) (int, [4]int) {
	factor1 := matrix[i][j]
	factor2 := matrix[i+1][j-1]
	factor3 := matrix[i+2][j-2]
	factor4 := matrix[i+3][j-3]
	product := factor1 * factor2 * factor3 * factor4

	factors := [4]int{factor1, factor2, factor3, factor4}
	return product, factors
}

func sumNorthWest(matrix [][]int, i int, j int) (int, [4]int) {
	factor1 := matrix[i][j]
	factor2 := matrix[i-1][j-1]
	factor3 := matrix[i-2][j-2]
	factor4 := matrix[i-3][j-3]
	product := factor1 * factor2 * factor3 * factor4

	factors := [4]int{factor1, factor2, factor3, factor4}
	return product, factors
}
