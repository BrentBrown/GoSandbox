package main

import "fmt"

//slices, maps and channels are referenceTypes
func main() {
	m := make([]string, 1, 25)
	fmt.Printf("address: %p \n", &m)
	fmt.Println(m)
	changeMe(m)
	fmt.Println(m)
}

func changeMe(z []string) {
	z[0] = "Hi there!"
	fmt.Printf("address: %p \n", &z) //new pointer, pointed at same underlying array
	fmt.Println(z)
}
