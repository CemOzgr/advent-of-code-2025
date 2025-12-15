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

	var result int
	unionFind := MakeSet(len(points))
	for i := 0; i < len(pairs); i++ {
		if unionFind.Find(pairs[i].point1) != unionFind.Find(pairs[i].point2) {
			result = points[pairs[i].point1].x * points[pairs[i].point2].x
		}

		unionFind.Union(pairs[i].point1, pairs[i].point2)
	}

	return result
}

func getDistance(p1 Point, p2 Point) int {
	dx := p1.x - p2.x
	dy := p1.y - p2.y
	dz := p1.z - p2.z

	return dx*dx + dy*dy + dz*dz
}
