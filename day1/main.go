package day1

import (
	"os"
	"strconv"
	"strings"
)

func Solve() int {
	content, err := os.ReadFile("./day1/input.txt")
	if err != nil {
		panic(err)
	}

	instructions := strings.Split(string(content)[3:], "\r\n")

	position, zeroCount := 50, 0
	for _, instruction := range instructions {
		isLeft := instruction[0] == 'L'
		rotation, err := strconv.Atoi(instruction[1:])
		if err != nil {
			panic(err)
		}

		zeroCount += calculateZeroCount(position, rotation, isLeft)
		position = calculateNewPosition(position, rotation, isLeft)
	}

	return zeroCount
}

func calculateNewPosition(current int, rotation int, isLeft bool) int {
	const Start, End = 0, 99
	normalizedRotation := rotation % 100

	var newPosition int
	if isLeft {
		newPosition = current - normalizedRotation
	} else {
		newPosition = current + normalizedRotation
	}

	if newPosition < Start {
		return End + 1 + newPosition
	}

	if newPosition > End {
		return newPosition % 100
	}

	return newPosition
}

func calculateZeroCount(current int, rotation int, isLeft bool) int {
	const Start, End = 0, 99
	var rotationsToZero int

	if current == Start {
		rotationsToZero = End + 1
	} else if isLeft {
		rotationsToZero = current
	} else {
		rotationsToZero = End + 1 - current
	}

	if rotationsToZero > rotation {
		return 0
	}

	zeroCount := 1
	zeroCount += (rotation - rotationsToZero) / 100

	return zeroCount
}
