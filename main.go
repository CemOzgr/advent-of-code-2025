package main

import (
	today "AdventOfCode/day6"
	"fmt"
	"os"
)

func main() {
	file, err := os.ReadFile("day6/input.txt")
	if err != nil {
		panic(err)
	}

	// removing BOM
	solution := today.Solve(string(file)[3:])

	fmt.Printf("solution: %d", solution)
}
