package main

import (
	"fmt"
)

//meaningful comment
func main() {
	var s, sep string
	hack := []string{"main", "foo", "bar"}
	for i := foo(); i < len(hack); i++ {
		s += sep + hack[i]
		sep = " "
	}
	fmt.Println(s)
}

func foo() int {
	return 1
}
