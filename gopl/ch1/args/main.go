package main

import (
	"fmt"
	"os"
)

func main() {
	a := os.Args[0]
	b := os.Args[1]
	c := os.Args[2]
	fmt.Println(a, b, c)
	foo()
}

func foo() {
	fmt.Println(os.Args[1])
}
