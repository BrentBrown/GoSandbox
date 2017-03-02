package main

import (
	"fmt"
)

func main() {
	var list = []rune("This is a string that needs to be reversed.")
	for index := len(list) - 1; index >=0; index-- {
		fmt.Printf("%c", list[index])
	}
}
