package main

import "fmt"

func filter(numbers []int, callback func(int) bool) []int {
	var xs []int
	for _, n := range numbers {
		if callback(n) {
			xs = append(xs, n)
		}
	}
	return xs
}

func main() {
	a := func(n int) bool {
		return n > 1
	}
	xs := filter([]int{1, 2, 3, 4}, a)
	fmt.Println(xs) // [2 3 4]
}
