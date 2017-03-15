package main

import "fmt"

func main() {
	answer := recursive(4)
	fmt.Printf("Recursive: %d\n", answer)

	tailAnswer := tailRecursive(4, 1)
	fmt.Printf("Tail Recursive: %d\n", tailAnswer)

	result := make(chan int)
	recursiveChannel(4, 1, result)
	channelAnswer := <-result
	fmt.Printf("Channel Recursive: %d\n", channelAnswer)
}

func recursive(number int) int {
	if number == 1 {
		return number
	}
	return number * recursive(number-1)
}

//compiler doesn't actually optimize stack in Go.
func tailRecursive(number int, product int) int {
	product = product * number

	if number == 1 {
		return product
	}
	return tailRecursive(number-1, product)
}

func recursiveChannel(number int, product int, result chan int) {
	product = product * number

	if number == 1 {
		result <- product
		return
	}
	go recursiveChannel(number-1, product, result)
}
