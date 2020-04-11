package main

import "fmt"

func median(i []int) float64 {
	// Defining one extra variable > calling array length 3 times
	var arrLength = len(i) / 2

	if len(i)%2 == 1 {
		return float64(i[arrLength])
	} else {
		return float64(i[arrLength]+i[arrLength-1]) / 2
	}
}

func main() {
	var i = []int{1, 2, 3, 4, 5}
	fmt.Println(median(i))
}
