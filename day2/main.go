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
			if isIdInvalid(strconv.Itoa(productId)) {
				sumOfInvalidIds += productId
			}
		}
	}

	return sumOfInvalidIds
}

func isIdInvalid(productId string) bool {
	length := len(productId)

	if length%2 != 0 {
		return false
	}

	if length == 2 {
		return productId[0] == productId[1]
	}

	right := length / 2
	for left := 0; left < length/2; left++ {
		if productId[left] != productId[right] {
			return false
		}

		right++
	}

	return true
}

func isIdInvalidStringCycling(productId string) bool {
	doubledId := (productId + productId)[1 : len(productId)*2-1]

}

func kmp(text string, pattern string) {
	lps := constructLps(pattern)
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