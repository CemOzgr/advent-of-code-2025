package day6

import (
	"strings"
)

type Thing struct {
	numbers   []int
	operation int32
}

func Solve(input string) int {
	lines := strings.Split(input, "\r\n")

	separators := make(map[int]struct{})

	for i, line := range lines {
		for j, c := range line {
			if c == ' ' && i == 0 {
				separators[j] = struct{}{}
				continue
			}

			if c != ' ' {
				delete(separators, j)
			}
		}
	}

	var operators []Thing
	for _, c := range lines[len(lines)-1] {
		if c != ' ' {
			operators = append(operators, Thing{operation: c})
		}
	}

	var currentOp int
	for i := range lines[0] {
		var current int
		_, ok := separators[i]

		if ok {
			currentOp++
			continue
		}

		for j := 0; j < len(lines)-1; j++ {
			if lines[j][i] != ' ' {
				current = 10*current + int(lines[j][i]-'0')
			}
		}

		operators[currentOp].numbers = append(operators[currentOp].numbers, current)
	}

	var result int
	for _, operation := range operators {
		var operationResult int
		if operation.operation == '*' {
			operationResult = 1
		}

		for _, number := range operation.numbers {
			if operation.operation == '*' {
				operationResult *= number
			} else {
				operationResult += number
			}
		}

		result += operationResult
	}

	return result
}
