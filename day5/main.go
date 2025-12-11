package day5

import (
	"slices"
	"strconv"
	"strings"
)

func Solve(input string) int {
	db := strings.Split(input, "\r\n\r\n")
	freshRanges := strings.Split(db[0], "\r\n")

	slices.Sort(freshRanges)

	var ranges Ranges

	for _, freshRange := range freshRanges {
		hmm := strings.Split(freshRange, "-")
		start, _ := strconv.Atoi(hmm[0])
		end, _ := strconv.Atoi(hmm[1])

		r := Range{start, end}
		ranges.Add(&r)
	}

	var result int
	for _, r := range ranges {
		result += r.end - r.start + 1
	}

	return result
}
