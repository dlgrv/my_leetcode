package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 6}
	target := 7
	fmt.Println(searchInsert(nums, target))
}

func searchInsert(nums []int, target int) int {
	for i, element := range nums {
		if element == target || element > target {
			return i
		}

		if i+1 == len(nums) && element < target {
			return i + 1
		}
	}

	return -1
}
