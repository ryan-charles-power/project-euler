package main

import (
	"fmt"
	"os"
	solution "project-euler/go"
)

func main() {
	// Map problem numbers to their solution functions
	solutions := map[string]func(){
		"1": solution.Solve1,
		"2": solution.Solve2,
		"3": solution.Solve3,
		"4": solution.Solve4,
		"5": solution.Solve5,
		// 6
		"7": solution.Solve7,
		// ...
		"11": solution.Solve11,
		// ...
		"100": solution.Solve100,
	}

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <problem_number>")
		return
	}

	problem := os.Args[1]
	if solveFunc, ok := solutions[problem]; ok {
		solveFunc()
	} else {
		fmt.Printf("Problem %s not implemented yet.\n", problem)
	}
}
