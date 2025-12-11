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

	var current int
	for _, line := range lines[:len(lines)-1] {
		var currentOp int
		for j, c := range line {
			_, ok := separators[j]

			if !ok && c != ' ' {
				current = current*10 + int(c-'0')
			}
			if ok || j == len(line)-1 {
				operators[currentOp].numbers = append(operators[currentOp].numbers, current)
				current = 0
				currentOp++
			}
		}
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
