package day4

import "strings"

func Solve(input string) int {
	warehouse := strings.Split(input, "\r\n")

	var neighborCount, accessibleRolls int
	for i, row := range warehouse {
		for j, _ := range row {
			if warehouse[i][j] != '@' {
				continue
			}

			if j != len(row)-1 && warehouse[i][j+1] == '@' {
				neighborCount++
			}

			if j != 0 && warehouse[i][j-1] == '@' {
				neighborCount++
			}

			if i != 0 {
				if warehouse[i-1][j] == '@' {
					neighborCount++
				}

				if j != len(row)-1 && warehouse[i-1][j+1] == '@' {
					neighborCount++
				}

				if j != 0 && warehouse[i-1][j-1] == '@' {
					neighborCount++
				}
			}

			if i < len(warehouse)-1 {
				if warehouse[i+1][j] == '@' {
					neighborCount++
				}

				if j < len(row)-1 && warehouse[i+1][j+1] == '@' {
					neighborCount++
				}

				if j != 0 && warehouse[i+1][j-1] == '@' {
					neighborCount++
				}
			}

			if neighborCount < 4 {
				accessibleRolls++
			}

			neighborCount = 0
		}
	}

	return accessibleRolls
}
