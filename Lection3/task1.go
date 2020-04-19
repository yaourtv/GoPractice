package main

import "fmt"

func average(arr [6]int) float64 {
	var sum int
	for i := 0; i < 6; i++ {
		sum += arr[i]
	}
	return float64(sum) / 6
}

func main() {
	var i = [6]int{5, 4, 1, 3, 2, 8}
	fmt.Println(average(i))
}
