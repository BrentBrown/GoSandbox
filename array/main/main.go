package main

import "fmt"

func main() {
	const s = 20
	var as, bs int
	var a, b [s]int

	o := getdata(s)

	for i, j := 0, len(o)-1; j >= 0; i, j = i+1, j-1 {
		a[i], b[j] = as+o[i], bs+o[j]
		as, bs = as+o[i], bs+o[j]
	}
	fmt.Println(o)
	fmt.Println(a)
	fmt.Println(b)
}

func getdata(s int) []int {
	o := make([]int, s)
	for i, j := 0, s-1; i <= j; i, j = i+1, j-1 {
		o[i], o[j] = i+1, j+1
	}
	return o
}
