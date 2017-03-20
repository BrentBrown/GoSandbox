package main

import (
	"fmt"
	"sort"
)

//Given two strings, write a method to decide if one is a permutation of the other.

type RuneSlice []rune
func(r RuneSlice) Less(i, j int) bool { return r[i] < r[j] }
func(r RuneSlice) Swap(i, j int) { r[i], r[j] = r[j], r[i] }
func(r RuneSlice) Len() int { return len(r) }

func main() {
	p := isPerm("abcd", "bcad")
	fmt.Println(p)
}

func isPerm(first string, second string) bool {
	if len(first) != len(second) { return false }
	f := RuneSlice([]rune(first))
	s := RuneSlice([]rune(second))
	sort.Sort(f)
	sort.Sort(s)
	return string(f) == string(s)
}