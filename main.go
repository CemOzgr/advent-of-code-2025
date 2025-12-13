package main

import (
	today "AdventOfCode/day7"
	"fmt"
	"os"
)

func main() {
	file, err := os.ReadFile("day7/input.txt")
	if err != nil {
		panic(err)
	}

	// removing BOM
	solution := today.Solve(string(file)[3:])

	fmt.Printf("solution: %d", solution)
}
