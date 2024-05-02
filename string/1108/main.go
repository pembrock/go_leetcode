package main

import (
	"fmt"
	"strings"
)

/*
Given a valid (IPv4) IP address, return a defanged version of that IP address.
A defanged IP address replaces every period "." with "[.]".

Example 1:
Input: address = "1.1.1.1"
Output: "1[.]1[.]1[.]1"

Example 2:
Input: address = "255.100.50.0"
Output: "255[.]100[.]50[.]0"
*/

func main() {
	address := "255.100.50.0"

	fmt.Println(defangIPaddr(address))
}

func defangIPaddr(address string) string {
	runesArray := []rune(address)
	var newAddr []rune
	lastIndex := 0
	for i, ch := range runesArray {
		if ch == '.' {
			newAddr = append(newAddr, runesArray[lastIndex:i]...)
			newAddr = append(newAddr, '[', '.', ']')
			lastIndex = i + 1
		}
	}
	newAddr = append(newAddr, runesArray[lastIndex:]...)
	return string(newAddr)
}

func defangIPaddrBuiltIn(address string) string {
	return strings.Replace(address, ".", "[.]", -1)
}
