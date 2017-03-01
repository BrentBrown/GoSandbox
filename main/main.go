package main

import (
	"fmt"
)

func main() {
	a := 42

	fmt.Printf("%v - %T \n", a, a)
	fmt.Printf("%v - %T \n", &a, &a)

	var b *int = &a
	fmt.Printf("%v - %T \n", b, b)
	fmt.Printf("%v - %T \n", *b, *b)

	*b = 27
	fmt.Printf("%v - %T \n", a, a)

	*&a = 12
	fmt.Printf("%v - %T \n", *b, *b)


}
