package main

import (
	"fmt"
	"math"
)

func main() {
	s := "abcde"
	t := "edbac"
	fmt.Println(findPermutationDifference(s, t))
}

func findPermutationDifference(s string, t string) int {
	tMap := make(map[rune]int)

	for i, c := range t {
		tMap[c] = i
	}
	result := 0
	for i, c := range s {
		result += int(math.Abs(float64(i - tMap[c])))
	}

	return result
}
