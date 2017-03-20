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

func urlify(s string, l int) string {
	var buffer bytes.Buffer
	d := len(s) - l
	a := []rune(s)
	for i := len(s) - 1 - d; i >= 0; i-- {
		if a[i] != ' ' {
			buffer.WriteString(string(a[i]))
		} else {
			buffer.WriteString("02%")
		}
	}
	u := []rune(buffer.String())

	for i, j := 0, len(u) - 1; i < j; i, j = i+1, j-1 {
		u[i], u[j] = u[j], u[i]
	}
	return string(u)
}

func main() {
	u := urlify("mr john doe    ", 11)
	fmt.Println(u)
}