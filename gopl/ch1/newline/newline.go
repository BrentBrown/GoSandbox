package main

import "fmt"

func main() {
	x, y := 1, 2

	//valid
	a := x + y

	//valid
	b := x +
		y

	//invalid
	/*
		c := x
		+y
	*/

	fmt.Println(a)
	fmt.Println(b)
}
