package main

import (
	"fmt"
	"bytes"
)

/*
Write a method to replace all spaces in a string with '%20. You may assume that the string
has sufficient space at the end to hold the additional characters, and that you are given the "true"
length of the string.
 */

func urlify(input string, length int) string {
	var buffer bytes.Buffer
	delta := len(input) - length
	url := []rune(input)
	for i := 0; i < len(input); i++ {
		if buffer.Len() >= length + delta { break }
		if url[i] != ' ' {
			buffer.WriteString(string(url[i]))
		} else {
			buffer.WriteString("%20")
		}
	}
	return buffer.String()
}

func main() {
	u := urlify("Mr John Doe Sr      ", 13)
	fmt.Println(u)
}