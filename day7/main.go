package day7

import "strings"

func Solve(input string) int {
	levels := strings.Split(input, "\r\n")[1:]

	beams := make([]bool, len(levels[0]))
	beams[strings.Index(input, "S")] = true

	var timesSplit int
	for _, level := range levels {
		for i := 0; i < len(level); i++ {
			if level[i] != '^' || !beams[i] {
				continue
			}

			beams[i] = false
			beams[i-1], beams[i+1] = true, true
			timesSplit++
		}
	}

	return timesSplit
}
