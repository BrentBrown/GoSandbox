package main

import "fmt"

func main() {
	//var student []string
	var student = make([]string, 0, 1)
	//var student = make([]string, 0, 200)

	println(len(student))
	println(cap(student))

	student = append(student, "Cyclops")
	println(len(student))
	println(cap(student))

	student = append(student, "Rogue")
	println(len(student))
	println(cap(student))

	student = append(student, "Jean")
	println(len(student))
	println(cap(student))

	student = append(student, "Logan")
	println(len(student))
	println(cap(student))

	fmt.Println(student)
}
