package main

import (
	"fmt"
)

func main() {
	fmt.Println(foo(2))

}

func foo(x int) (y, z int) {
	y = x * 10
	z = x * 20
	return
}
