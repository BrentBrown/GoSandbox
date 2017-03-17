package main

import "fmt"

var data = make([]int, 10, 10)

func init() {
	for i := 0; i < 10; i++ {
		data[i] = i * 10
	}
}

func main() {
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data); j++ {
			fmt.Printf("[%d][%d]\n", data[i], data[j])
		}
	}
}
