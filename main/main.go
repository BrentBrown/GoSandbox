package main

import (
	"fmt"
)

func main() {
	c, d := foo(2)
	fmt.Println(c)
	fmt.Println(&c)
	fmt.Println(&d)
	temp := foo
	fmt.Println(&temp)

}

func foo(x int) (y, z int) {
	y = x * 10
	z = x * 20
	return
}
