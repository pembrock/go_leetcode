package main

import "fmt"

func main() {
	command := "(al)G(al)()()G"
	fmt.Println(interpret(command))
}

func interpret(command string) string {
	result := ""
	for i, c := range command {
		if c == '(' && command[i+1] == ')' {
			result += "o"
		} else {
			if c != '(' && c != ')' {
				result += string(c)
			}
		}
	}
	return result
}
