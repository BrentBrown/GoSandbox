package main

import "fmt"

//Given a string, write a function to check  if it  is  a permutation of a palindrome.

func isPal(input string) bool {
	inputSlice := []rune(input)
	length := len(inputSlice)
	palmap := make(map[rune]int)

	for i := 0; i < length; i++ {
		palmap[inputSlice[i]]++
	}

	numOdd := 0
	for k, v := range palmap {
		if k == ' ' {
			continue
		}
		if v%2 != 0 && numOdd != 0 {
			return false
		}
		if v%2 != 0 {
			numOdd++
		}
	}
	return true
}

func main() {
	p := isPal("tact coa")
	fmt.Println(p)
}
