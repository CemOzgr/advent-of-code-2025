package day8

import (
	"sort"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
	z int
}

type Pair struct {
	point1   int
	point2   int
	distance int
}

func Solve(input string) int {
	lines := strings.Split(input, "\r\n")
	points := make([]Point, len(lines))

	for i, line := range lines {
		coordinates := strings.Split(line, ",")
		x, _ := strconv.Atoi(coordinates[0])
		y, _ := strconv.Atoi(coordinates[1])
		z, _ := strconv.Atoi(coordinates[2])

		points[i] = Point{x, y, z}
	}

	pairs := make([]Pair, len(points)*(len(points)-1)/2)
	var pairCount int
	for i, point := range points {
		for j := i + 1; j < len(points); j++ {
			pairs[pairCount] = Pair{
				i,
				j,
				getDistance(point, points[j]),
			}
			pairCount++
		}
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].distance < pairs[j].distance
	})

	unionFind := MakeSet(len(points))
	for i := 0; i < 1000; i++ {
		unionFind.Union(pairs[i].point1, pairs[i].point2)
	}

	var max1, max2, max3 int
	for _, x := range unionFind.sizes {
		if x > max1 {
			max3 = max2
			max2 = max1
			max1 = x
		} else if x > max2 {
			max3 = max2
			max2 = x
		} else if x > max3 {
			max3 = x
		}
	}

	return max1 * max2 * max3
}

func getDistance(p1 Point, p2 Point) int {
	dx := p1.x - p2.x
	dy := p1.y - p2.y
	dz := p1.z - p2.z

	return dx*dx + dy*dy + dz*dz
}
