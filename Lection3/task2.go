package main

import "fmt"

func max(arr []string) string {
	longest, length := "", 0

	for i := 0; i < len(arr); i++ {
		if len(arr[i]) > length {
			longest = arr[i]
			length = len(arr[i])
		}
	}
	return longest
}

func reverse(arr []int64) []int64 {
	var copiedArr = make([]int64, len(arr))
	copy(copiedArr, arr)

	for i, j := 0, len(copiedArr)-1; i < j; i, j = i+1, j-1 {
		copiedArr[i], copiedArr[j] = copiedArr[j], copiedArr[i]
	}
	return copiedArr
}

func main() {
	var i = []int64{5, 4, 1, 3, 2}
	var j = []string{"so", "you", "are", "programmer", "name", "every", "code"}
	fmt.Println(reverse(i))
	fmt.Println(max(j))
}
