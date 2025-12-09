package main

import (
	"AdventOfCode/day3"
	"fmt"
	"os"
)

func main() {
	file, err := os.ReadFile("day3/input.txt")
	if err != nil {
		panic(err)
	}

	// removing BOM
	solution := day3.Solve(string(file)[3:])

	fmt.Printf("day3: %d", solution)
}
