package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Square struct {
	side float64
}

type Circle struct {
	radius float64
}

func info(z Shape) {
	fmt.Println(z.area())
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (s Square) area() float64 {
	return s.side * s.side
}

func main() {
	var c = Circle{radius: 33}
	var s = Square{side: 12}
	info(c)
	info(s)
}
