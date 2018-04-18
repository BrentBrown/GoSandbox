package main

import (
	"fmt"
)

//meaningful comment
func main() {
	var s, sep string
	hack := []string{"main", "foo", "bar"}
	for i := 1; i < len(hack); i++ {
		s += sep + hack[i]
		sep = " "
	}
	fmt.Println(s)
}
