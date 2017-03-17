package main

import "fmt"

var dataA = make([]int, 10, 10)
var dataB = make([]int, 10, 10)

func init() {
	for i := 0; i < 10; i++ {
		dataA[i] = i * 10
		dataB[i] = i * 20
	}
}

func main() {
	for i := 0; i < len(dataA); i++ {
		for j := 0; j < len(dataB); j++ {
			fmt.Printf("[%d][%d]\n", dataA[i], dataB[j])
		}
	}
}

//Not O(N^2) because of two different input, both of which must be considered
