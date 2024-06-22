package main

import (
	"fmt"
	"strings"
)

func main() {
	word1 := []string{"ab", "c"}
	word2 := []string{"a", "bc"}
	fmt.Println(arrayStringsAreEqual(word1, word2))
}

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	sb1 := &strings.Builder{}
	sb2 := &strings.Builder{}
	for _, elem := range word1 {
		sb1.WriteString(elem)
	}

	for _, elem := range word2 {
		sb2.WriteString(elem)
	}

	return sb1.String() == sb2.String()
}
