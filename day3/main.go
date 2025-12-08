package day3

import (
	"os"
	"strings"
)

func Solve() int {
	file, err := os.ReadFile("day3/input.txt")
	if err != nil {
		panic(err)
	}

	banks := strings.Split(string(file)[3:], "\r\n")

	var totalJolts int
	for _, bank := range banks {
		totalJolts += findMaxJolts(bank)
	}

	return totalJolts
}

func findMaxJolts(bank string) int {
	var l, r = 0, 1

	var maxJolt int
	for l < r && r < len(bank) {
		leftJolt := int(bank[l] - '0')
		rightJolt := int(bank[r] - '0')
		combinedJolt := 10*leftJolt + rightJolt

		if combinedJolt > maxJolt {
			maxJolt = combinedJolt
		}

		if rightJolt >= leftJolt {
			l = r
		}

		r++
	}

	return maxJolt
}
