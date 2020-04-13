package main

import "fmt"

func median(i []int) float64 {
	var pivot = len(i) / 2

	bubbleSort(&i)
	if len(i)%2 == 1 {
		return float64(i[pivot])
	}

	return float64(i[pivot]+i[pivot-1]) / 2
}

func bubbleSort(arr *[]int) {
	var arrLength = len(*arr)

	for i := 0; i < arrLength; i++ {
		for j := 0; j < arrLength-i-1; j++ {

			if (*arr)[j] > (*arr)[j+1] {
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
			}
		}
	}
}

func main() {
	var i = []int{5, 4, 1, 3, 2}
	var j = []int{7, 1, 5, 4, 3, 6}
	fmt.Println(median(i))
	fmt.Println(median(j))
}
