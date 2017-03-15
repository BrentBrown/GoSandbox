package main

import "fmt"

func main() {

	//conversion
	num := 3.14
	fmt.Printf("%T\n", num)
	fmt.Printf("%T\n", int(num))

	//assertion
	var val interface{} = 7
	fmt.Printf("%T\n", val)
	//fmt.Printf("%T\n", int(val))  //cannot convert val (type interface {}) to type int: need type assertion
	fmt.Printf("%T\n", val.(int))
}
