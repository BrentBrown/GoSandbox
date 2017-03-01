package main

import (
	"fmt"
	_ "foo"
)

func init() {
	fmt.Println("init")
}

func main() {
	fmt.Println("main")
}
