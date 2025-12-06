package day2

import (
	"os"
	"strconv"
	"strings"
)

type ranger struct {
	start int
	end   int
}

func Solve() {
	file, fileReadError := os.ReadFile("day2/input.txt")
	if fileReadError != nil {
		panic(fileReadError)
	}

	ranges := strings.Split(string(file), ",")

	rangers := make([]ranger, len(ranges))
	for i, r := range ranges {
		values := strings.Split(r, "-")
		start, err1 := strconv.Atoi(values[0])
		if err1 != nil {
			panic(err1)
		}
		end, err2 := strconv.Atoi(values[1])
		if err2 != nil {
			panic(err2)
		}

		rangers[i] = ranger{
			start: start,
			end:   end,
		}
	}
}
