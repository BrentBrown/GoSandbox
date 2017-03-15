package main

import "fmt"

func main() {

	fmt.Println(average(233, 432434, 223, 967867, 54674))

	s := []float64{233, 432434, 223, 967867, 54674}
	fmt.Println(average(s...))
}

func average(sf ...float64) float64 {
	//fmt.Println(sf)
	//fmt.Printf("%T \n", sf)
	var total float64
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}
