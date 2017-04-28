package main

import (
	"fmt"

	"./tempconv0"
)

func main() {
	fmt.Println(tempconv0.CToF(tempconv0.AbsoluteZeroC))
	f := new(tempconv0.Fahrenheit)
	*f = 5
	c := tempconv0.Celsius(*f)
	c = 10
	fmt.Println(c)
}
