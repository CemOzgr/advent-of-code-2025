package day3

import (
	"strings"
)

const NumberOfBatteries = 12

func Solve(input string) int {
	banks := strings.Split(input, "\r\n")

	var totalJolts int
	for _, bank := range banks {
		totalJolts += findMaxJolts(bank, NumberOfBatteries)
	}

	return totalJolts
}

func findMaxJolts(bank string, numberOfBatteries int) int {
	batteries := make([]int, len(bank))
	for i, jolt := range bank {
		batteries[i] = int(jolt - '0')
	}

	stack := NewStack(numberOfBatteries)

	for i, jolt := range batteries {
		if stack.IsEmpty() {
			stack.Push(jolt)
			continue
		}

		allowedSizeToPop := len(batteries) + stack.top - i - numberOfBatteries
		for stack.ShouldPop(jolt) && allowedSizeToPop > 0 {
			stack.Pop()
		}

		if stack.top < NumberOfBatteries {
			stack.Push(jolt)
		}
	}

	var result int
	for _, jolt := range stack.arr {
		result = 10*result + jolt
	}

	return result
}
