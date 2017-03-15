package main

import (
	"fmt"
)

func main() {
	c, _ := foo(2)
	fmt.Println(c)
}

func foo(x int) (y, z int) {
	y = x * 10
	z = x * 20
	return
}
