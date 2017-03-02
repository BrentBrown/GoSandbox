package main

import (
	"fmt"
)

func zero(x int) {
	fmt.Printf("zero x value - type: %v - %T \n", x, x)
	fmt.Printf("zero x address - type: %v - %T \n", &x, &x)
	x = 0
}

func zeropointer(x *int) {
	fmt.Printf("zeropointer x value - type: %v - %T \n", x, x)
	*x = 0
}

func main() {
	x := 5
	fmt.Printf("main x value - type: %v - %T \n", x, x)
	fmt.Printf("main x address - type: %v - %T \n", &x, &x)

	zero(x)
	fmt.Printf("main x value - type after call to zero - type: %v - %T \n", x, x)

	zeropointer(&x)
	fmt.Printf("main x value - type after call to zeropointer: %v - %T \n", x, x)
}
