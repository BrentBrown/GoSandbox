package main

import (
	"fmt"
)


func zero(x int) {
	fmt.Printf("%v - %T \n", x, x)
	x = 0
}

func zerop(x *int) {
	fmt.Printf("%v - %T \n", x, x)
	*x = 0
}

func main() {
	x := 5
	fmt.Printf("%v - %T \n", x, x)
	zero(x)
	fmt.Printf("%v - %T \n", x, x)

	zerop(&x)
	fmt.Printf("%v - %T \n", x, x)
}
