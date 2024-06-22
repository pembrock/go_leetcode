package main

import "fmt"

func main() {
	word := "abcdefd"
	ch := 'd'
	fmt.Printf(reversePrefix(word, byte(ch)))
}

func reversePrefix(word string, ch byte) string {
	var result []byte

	chIndex := -1
	for i, c := range word {
		if byte(c) == ch {
			chIndex = i
			break
		}
	}

	if chIndex == -1 {
		return word
	}

	for i := chIndex; i >= 0; i-- {
		result = append(result, word[i])
	}

	result = append(result, []byte(word)[chIndex+1:]...)

	return string(result)
}
