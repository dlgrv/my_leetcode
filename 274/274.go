package main

import (
	"fmt"
	"sort"
)

func main() {
	// citations := []int{3, 0, 6, 1, 5}
	// citations := []int{2, 2, 2, 2, 2}
	// citations := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 9, 10, 11, 11, 11}
	citations := []int{11, 15}
	fmt.Println(hIndex(citations))
}

func hIndex(citations []int) int {
    citationsSlice := citations[:]
	sort.Ints(citationsSlice)

	citationsSize := len(citations)

	if citations[citationsSize-1] == 0 {
		return 0
	} else if citationsSize == 1 {
		return 1
	}

	h_index := 0

	for h:=1; h<=citationsSize; h++  {
		if citations[citationsSize-h] >= h {
			h_index = h
		} else {
			break
		}
	}

	return h_index
}