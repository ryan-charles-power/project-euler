package solutions

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strings"
)

func Solve13() {
	// Open the file
	file, err := os.Open("../project-euler/resources/problem13list.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	total := new(big.Int) // Initializes to 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// Clean up the line (remove whitespace/newlines)
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		// Parse the string into a big.Int
		currentNum := new(big.Int)
		currentNum, ok := currentNum.SetString(line, 10)
		if !ok {
			fmt.Printf("Skipping invalid number: %s\n", line)
			continue
		}

		// Add to total: total = total + currentNum
		total.Add(total, currentNum)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}

	fmt.Printf("The total sum is: %s\n", total.String())
}
