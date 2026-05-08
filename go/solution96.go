package solutions

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Solve96() {
	// Read the entire file content into a byte slice.
	fileContent, err := os.ReadFile("../project-euler/resources/problem96sudoku.txt")
	if err != nil {
		log.Fatal(err) // Handle potential errors, such as file not found.
	}

	// Convert the byte slice to a string.
	str := string(fileContent)

	// prepare the games
	games := parseGrids(str)

	sum := 0
	for gameNum := 0; gameNum < 50; gameNum++ {
		var game [9][9]int
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				game[i][j] = games[gameNum][i][j]
			}
		}
		_ = solve(&game)

		sum += game[0][0]*100 + game[0][1]*10 + game[0][2]
	}

	fmt.Printf("The sum of all 3 digit numbers in the solutions of all sudokus is %d", sum)

}

// Parsing the sudoku grids from the original txt
func parseGrids(input string) [][][]int {
	var games [][][]int
	var currentGrid [][]int

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		// If the line starts with "Grid", it's the start of a new game
		if strings.HasPrefix(line, "Grid") {
			if len(currentGrid) > 0 {
				games = append(games, currentGrid)
			}
			currentGrid = [][]int{}
			continue
		}

		// Otherwise, convert the string of digits into an int slice (row)
		row := make([]int, len(line))
		for i, char := range line {
			val, _ := strconv.Atoi(string(char))
			row[i] = val
		}
		currentGrid = append(currentGrid, row)
	}

	// Append the final grid in the buffer
	if len(currentGrid) > 0 {
		games = append(games, currentGrid)
	}

	return games
}

// Sudoku algorith
func solve(board *[9][9]int) bool {
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] == 0 {
				for num := 1; num <= 9; num++ {
					if isValid(board, r, c, num) {
						board[r][c] = num
						if solve(board) {
							return true
						}
						board[r][c] = 0 // Backtrack
					}
				}
				return false // No valid number found for this cell
			}
		}
	}
	return true // All cells filled
}

func isValid(board *[9][9]int, row, col, num int) bool {
	for i := 0; i < 9; i++ {
		// Check row, column, and 3x3 subgrid simultaneously
		if board[row][i] == num || board[i][col] == num ||
			board[3*(row/3)+i/3][3*(col/3)+i%3] == num {
			return false
		}
	}
	return true
}
