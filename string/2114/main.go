package main

import (
	"fmt"
	"strings"
)

func main() {
	sentences := []string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"}
	fmt.Println(mostWordsFound(sentences))
}

func mostWordsFound(sentences []string) int {
	result := 0
	for _, sentence := range sentences {
		words := strings.Split(sentence, " ")
		if len(words) > result {
			result = len(words)
		}
	}

	return result
}
