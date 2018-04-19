package main

import "fmt"

func main() {
	a := []string{"main", "foo", "bar"}

	//index always returned first
	for i := range a {
		fmt.Println(a[i])
	}

	//can also return value at each index
	for _, i := range a {
		fmt.Println(i)
	}
}
