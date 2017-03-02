package main

import (
	"fmt"
)

func main() {
	list := []rune("This is a string that needs to be reversed.")
	for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
		list[i], list[j] = list[j], list[i]
	}
	fmt.Println(string(list))
}
