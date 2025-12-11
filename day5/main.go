package day5

import (
	"strconv"
	"strings"
)

type Range struct {
	start int
	end   int
}

func Solve(input string) int {
	db := strings.Split(input, "\r\n\r\n")
	freshRanges := strings.Split(db[0], "\r\n")
	availableIngredients := strings.Split(db[1], "\r\n")

	ingredients := make([]int, len(availableIngredients))
	ranges := make([]Range, len(freshRanges))

	for i, ingredient := range availableIngredients {
		hmm, _ := strconv.Atoi(ingredient)
		ingredients[i] = hmm
	}

	for i, freshRange := range freshRanges {
		hmm := strings.Split(freshRange, "-")
		start, _ := strconv.Atoi(hmm[0])
		end, _ := strconv.Atoi(hmm[1])

		ranges[i] = Range{
			start: start,
			end:   end,
		}
	}

	var result int
	for _, ingredient := range ingredients {
		for _, freshRange := range ranges {
			if ingredient >= freshRange.start && ingredient <= freshRange.end {
				result++
				break
			}
		}
	}

	return result
}
