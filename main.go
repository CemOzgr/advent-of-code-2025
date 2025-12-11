package main

import (
	"AdventOfCode/day4"
	"fmt"
	"os"
)

func main() {
	file, err := os.ReadFile("day4/input.txt")
	if err != nil {
		panic(err)
	}

	// removing BOM
	solution := day4.Solve(string(file)[3:])

	fmt.Printf("day4: %d", solution)
}
