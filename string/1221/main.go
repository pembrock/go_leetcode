package main

import "fmt"

func main() {
	s := "RLRRLLRLRL"
	fmt.Println(balancedStringSplit(s))
}

func balancedStringSplit(s string) int {
	lCount := 0
	rCount := 0
	bsCount := 0
	for _, c := range s {
		if c == 'R' {
			rCount += 1
		} else if c == 'L' {
			lCount += 1
		}

		if rCount == lCount {
			bsCount += 1
			rCount = 0
			lCount = 0
		}
	}

	return bsCount
}
