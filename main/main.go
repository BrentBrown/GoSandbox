package main

import (
	"fmt"
)

var x = 0

func main() {

	x := 40

	foo := func() {
		fmt.Println(x)
	}

	fmt.Println(x)
	{
		x := 1
		y := x
		fmt.Println(x)
		{
			fmt.Println(y)
			foo()
		}

		increment := wrapper()
		fmt.Println(increment())
		fmt.Println(increment())
	}
}


func wrapper() func() int {
	//x := 10
	return func() int {
		x++
		return x
	}
}

