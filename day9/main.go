package day9

import (
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

func Solve(input string) int {
	lines := strings.Split(input, "\r\n")
	points := make([]Point, len(lines))

	for i, line := range lines {
		values := strings.Split(line, ",")
		x, _ := strconv.Atoi(values[0])
		y, _ := strconv.Atoi(values[1])

		points[i] = Point{x, y}
	}

	var area, maxArea int
	for i, point1 := range points {
		for j := i + 1; j < len(points); j++ {
			area = (abs(point1.x-points[j].x) + 1) * (abs(point1.y-points[j].y) + 1)
			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea
}

func abs(number int) int {
	if number > 0 {
		return number
	}

	return -number
}
