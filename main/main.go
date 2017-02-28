package main

import (
	"fmt"
	"stringutils"
)

var h string = "hi"

func main() {

	a := 10
	b := "golang"
	c := 4.17
	d := true
	e := "Hello"
	f := `Do you like my hat?`
	g := 'M'
	//a = 11
	//h := 12

	fmt.Printf("%v - %T \n", a, a)
	fmt.Printf("%v - %T \n", b, b)
	fmt.Printf("%v - %T \n", c, c)
	fmt.Printf("%v - %T \n", d, d)
	fmt.Printf("%v - %T \n", e, e)
	fmt.Printf("%v - %T \n", f, f)
	fmt.Printf("%v - %T \n", g, g)
	//fmt.Printf("%v - %T \n", h, h)
	otherPrint()

	fmt.Println(stringutils.Reverse("Hi There!"))

	max := max(7)
	fmt.Println(max)
	fmt.Printf("%v - %T \n", X, X)
}
