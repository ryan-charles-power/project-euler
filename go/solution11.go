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
				if product > maxProduct {
					maxProduct = product
					maxProducts = products
				}
			}
			if i > 2 && j < 17 {
				product, products = sumSouthWest(matrix, i, j)
				if product > maxProduct {
					maxProduct = product
					maxProducts = products
				}
			}
			if i < 17 && j > 2 {
				product, products = sumNorthEast(matrix, i, j)
				if product > maxProduct {
					maxProduct = product
					maxProducts = products
				}
			}
			if i < 17 && j < 17 {
				product, products = sumSouthEast(matrix, i, j)
				if product > maxProduct {
					maxProduct = product
					maxProducts = products
				}
			}
			if i > 2 {
				product, products = sumWest(matrix, i, j)
				if product > maxProduct {
					maxProduct = product
					maxProducts = products
				}

			}
			if i < 17 {
				product, products = sumEast(matrix, i, j)
				if product > maxProduct {
					maxProduct = product
					maxProducts = products
				}
			}
			if j > 2 {
				product, products = sumNorth(matrix, i, j)
				if product > maxProduct {
					maxProduct = product
					maxProducts = products
				}
			}
			if j < 17 {
				product, products = sumSouth(matrix, i, j)
				if product > maxProduct {
					maxProduct = product
					maxProducts = products
				}
			}
		}
	}

	fmt.Printf("The max product is %d made of the products %d", maxProduct, maxProducts)
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

func sumSouth(matrix [][]int, i int, j int) (int, [4]int) {
	factor1 := matrix[i][j]
	factor2 := matrix[i][j+1]
	factor3 := matrix[i][j+2]
	factor4 := matrix[i][j+3]
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

func sumWest(matrix [][]int, i int, j int) (int, [4]int) {
	factor1 := matrix[i][j]
	factor2 := matrix[i-1][j]
	factor3 := matrix[i-2][j]
	factor4 := matrix[i-3][j]
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

func sumSouthEast(matrix [][]int, i int, j int) (int, [4]int) {
	factor1 := matrix[i][j]
	factor2 := matrix[i+1][j+1]
	factor3 := matrix[i+2][j+2]
	factor4 := matrix[i+3][j+3]
	product := factor1 * factor2 * factor3 * factor4

	factors := [4]int{factor1, factor2, factor3, factor4}
	return product, factors
}

func sumSouthWest(matrix [][]int, i int, j int) (int, [4]int) {
	factor1 := matrix[i][j]
	factor2 := matrix[i-1][j+1]
	factor3 := matrix[i-2][j+2]
	factor4 := matrix[i-3][j+3]
	product := factor1 * factor2 * factor3 * factor4

	factors := [4]int{factor1, factor2, factor3, factor4}
	return product, factors
}
