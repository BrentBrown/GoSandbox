package main

import "fmt"

/*
Write a method to replace all spaces in a string with '%20. You may assume that the string
has sufficient space at the end to hold the additional characters, and that you are given the "true"
length of the string.
*/

func urlify(input string, truelength int) string {
	inputSlice := []rune(input)
	numSpaces := 0
	for i := truelength - 1; i >= 0; i-- {
		if inputSlice[i] == ' ' {
			numSpaces++
		}
	}

	index := truelength + (numSpaces * 2)
	urlSlice := make([]rune, index, index)

	for i := truelength - 1; i >= 0; i-- {
		if inputSlice[i] == ' ' {
			urlSlice[index-1] = '0'
			urlSlice[index-2] = '2'
			urlSlice[index-3] = '%'
			index = index - 3
		} else {
			urlSlice[index-1] = inputSlice[i]
			index--
		}
	}
	return string(urlSlice)
}

func main() {
	u := urlify("mr john doe    ", 11)
	fmt.Println(u)
}
