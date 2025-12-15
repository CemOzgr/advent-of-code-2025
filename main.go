package main

import (
	today "AdventOfCode/day9"
	"fmt"
	"os"
)

func main() {
	file, err := os.ReadFile("day9/input.txt")
	if err != nil {
		panic(err)
	}

	// removing BOM
	solution := today.Solve(string(file)[3:])

	fmt.Printf("solution: %d", solution)
}
