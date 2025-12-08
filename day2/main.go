package day2

import (
	"os"
	"strconv"
	"strings"
)

func Solve() int {
	file, fileReadError := os.ReadFile("day2/input.txt")
	if fileReadError != nil {
		panic(fileReadError)
	}

	ranges := strings.Split(strings.ReplaceAll(string(file[3:]), "\r\n", ""), ",")

	var sumOfInvalidIds int
	for _, r := range ranges {
		limits := strings.Split(r, "-")

		start, err1 := strconv.Atoi(limits[0])
		if err1 != nil {
			panic(err1)
		}

		end, err2 := strconv.Atoi(limits[1])
		if err2 != nil {
			panic(err2)
		}

		for productId := start; productId <= end; productId++ {
			if isIdInvalidStringCycling(strconv.Itoa(productId)) {
				sumOfInvalidIds += productId
			}
		}
	}

	return sumOfInvalidIds
}

func isIdInvalidStringCycling(productId string) bool {
	/*
		When we add a periodic string to itself,
		the last occurrence of the pattern in the original string
		and the first occurrence of the pattern in the added string
		will create the original string, thus we can say:
		a string is periodic if it contains itself when doubled
		note: since every string contains itself, we only check the middle part
	*/
	doubledId := (productId + productId)[1 : len(productId)*2-1]

	return kmp(doubledId, productId)
}

func kmp(text string, pattern string) bool {
	lps := constructLps(pattern)

	var i, j int

	for i < len(text) {
		if text[i] == pattern[j] {
			i++
			j++

			if j == len(pattern) {
				return true
			}

			continue
		}

		if j == 0 {
			i++
			continue
		}

		j = lps[j-1]
	}

	return false
}

func constructLps(pattern string) []int {
	lps := make([]int, len(pattern))
	length := 0

	for i := 1; i < len(pattern); {
		if pattern[length] == pattern[i] {
			length++
			lps[i] = length
			i++

			continue
		}

		if length == 0 {
			lps[i] = 0
			i++

			continue
		}

		length = lps[length-1]
	}

	return lps
}
