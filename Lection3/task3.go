package main

import (
	"fmt"
	"sort"
)

func printSorted(unsorted map[int]string) {
	var keys = make([]int, 0, len(unsorted))
	for key := range unsorted {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	for _, key := range keys {
		fmt.Println(unsorted[key])
	}
}

func main() {
	var i = map[int]string{2: "a", 0: "b", 1: "c"}
	printSorted(i)
}
