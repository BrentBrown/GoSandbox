package main

import (
	"fmt"
	"time"
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s", name, elapsed)
	fmt.Println()
}

func main() {
	defer timeTrack(time.Now(), "main")
	const s = 10
	const n = 1
	var as, bs int
	var a, b [s]int
	var p = false

	for z := 0; z < n; z++ {
		o := getOriginalData(s)

		for i, j := 0, len(o)-1; j >= 0; i, j = i+1, j-1 {
			a[i], b[j] = as+o[i], bs+o[j]
			as, bs = as+o[i], bs+o[j]
		}

		if p {
			fmt.Println(o)
			fmt.Println(a)
			fmt.Println(b)
			fmt.Println()
		}
	}

}

func getOriginalData(s int) []int {
	o := make([]int, s)
	for i, j := 0, s-1; i <= j; i, j = i+1, j-1 {
		o[i], o[j] = i+1, j+1
	}
	return o
}
