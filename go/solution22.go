package solutions

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

var charToValue = map[string]int{
	"A": 1,
	"B": 2,
	"C": 3,
	"D": 4,
	"E": 5,
	"F": 6,
	"G": 7,
	"H": 8,
	"I": 9,
	"J": 10,
	"K": 11,
	"L": 12,
	"M": 13,
	"N": 14,
	"O": 15,
	"P": 16,
	"Q": 17,
	"R": 18,
	"S": 19,
	"T": 20,
	"U": 21,
	"V": 22,
	"W": 23,
	"X": 24,
	"Y": 25,
	"Z": 26,
}

func Solve22() {
	nameScore := 0

	// Read the entire file content into a byte slice.
	fileContent, err := os.ReadFile("../project-euler/resources/problem22names.txt")
	if err != nil {
		log.Fatal(err) // Handle potential errors, such as file not found.
	}

	// Convert the byte slice to a string.
	nameStr := string(fileContent)

	// Split into names
	nameStr = strings.ReplaceAll(nameStr, "\"", "")
	nameList := strings.Split(strings.TrimSpace(nameStr), ",")

	// Sort names into alphabetical order
	slices.Sort(nameList)

	// Alphabetical value * sorted number
	for i := 0; i < len(nameList); i++ {
		thisNameScore := 0
		currentName := nameList[i]

		for j := 0; j < len(currentName); j++ {
			char := string(currentName[j])
			thisNameScore += charToValue[char]
		}

		thisNameScore *= (i + 1)
		nameScore += thisNameScore
	}

	fmt.Printf("The total name score is %d", nameScore)
}
