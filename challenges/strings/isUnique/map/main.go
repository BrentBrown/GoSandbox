package main

import "fmt"

//Is Unique: Implement an algorithm  to  determine  if  a string has  all  unique characters.

func isUnique(array []string) bool {
	m := make(map[string]int)
	for i := 0; i < len(array); i++ {
		if m[array[i]] == 0 {
			m[array[i]]++
		} else {
			return false
		}
	}
	return true
}

func main() {
	u := isUnique([]string {"a", "b", "b"})
	fmt.Println(u)
}
