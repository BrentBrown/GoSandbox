package main

import "fmt"

func main() {

	fmt.Printf("%d - %b - %x \n", 42, 42, 42)
	fmt.Printf("%d - %b - %#x \n", 42, 42, 42)
	fmt.Printf("%d - %b - %#X \n", 42, 42, 42)
	fmt.Printf("%d \t %b \t %#X \n", 42, 42, 42)

	for index := 1; index < 256; index++ {
		fmt.Printf("%d \t %b \t %x \n", index, index, index)
	}
}
