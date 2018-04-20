package main

import (
	"fmt"
	"os"
	"strings"
)

//meaningful comment
func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
