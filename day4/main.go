package day4

import "strings"

type Pair struct {
	i int
	j int
}

func Solve(input string) int {
	warehouse := strings.Split(input, "\r\n")
	columns := strings.Index(input, "\r\n")

	matrix := make([][]bool, len(warehouse))

	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]bool, columns)
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < columns; j++ {
			matrix[i][j] = warehouse[i][j] == '@'
		}
	}

	var accessibleRolls int

	for {
		accessed, toRemove := Visit(matrix)
		if accessed == 0 {
			break
		}

		accessibleRolls += accessed
		for _, pair := range toRemove {
			matrix[pair.i][pair.j] = false
		}
	}

	return accessibleRolls
}

func Visit(matrix [][]bool) (int, []Pair) {
	var neighborCount, accessibleRolls int
	var removeBeforeNextVisit []Pair

	for i, row := range matrix {
		for j, _ := range row {
			if !matrix[i][j] {
				continue
			}

			if j != len(row)-1 && matrix[i][j+1] {
				neighborCount++
			}

			if j != 0 && matrix[i][j-1] {
				neighborCount++
			}

			if i != 0 {
				if matrix[i-1][j] {
					neighborCount++
				}

				if j != len(row)-1 && matrix[i-1][j+1] {
					neighborCount++
				}

				if j != 0 && matrix[i-1][j-1] {
					neighborCount++
				}
			}

			if i < len(matrix)-1 {
				if matrix[i+1][j] {
					neighborCount++
				}

				if j < len(row)-1 && matrix[i+1][j+1] {
					neighborCount++
				}

				if j != 0 && matrix[i+1][j-1] {
					neighborCount++
				}
			}

			if neighborCount < 4 {
				accessibleRolls++
				removeBeforeNextVisit = append(removeBeforeNextVisit, Pair{i, j})
			}

			neighborCount = 0
		}
	}

	return accessibleRolls, removeBeforeNextVisit
}
