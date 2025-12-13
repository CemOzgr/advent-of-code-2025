package day7

import "strings"

func Solve(input string) int {
	levels := strings.Split(input, "\r\n")[1:]
	beamPosition := strings.Index(input, "S")

	cache := make([][]int, len(levels))
	for i := 0; i < len(cache); i++ {
		cache[i] = make([]int, len(levels[0]))
	}

	return CountWorlds(cache, levels, 0, beamPosition)
}

func CountWorlds(cache [][]int, levels []string, levelIndex int, beamPosition int) int {
	if levelIndex == len(levels)-1 {
		if levels[levelIndex][beamPosition] == '^' {
			return 2
		}

		return 1
	}

	if cache[levelIndex][beamPosition] != 0 {
		return cache[levelIndex][beamPosition]
	}

	for i, c := range levels[levelIndex] {
		if c == '^' && beamPosition == i {
			worlds := CountWorlds(cache, levels, levelIndex+1, i-1) + CountWorlds(cache, levels, levelIndex+1, i+1)
			cache[levelIndex][beamPosition] = worlds

			return worlds
		}
	}

	return CountWorlds(cache, levels, levelIndex+1, beamPosition)
}
