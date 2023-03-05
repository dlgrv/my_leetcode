package main

import "fmt"

func main() {
	nums := []int{10, 20, 30, 40, 50}
	target := 50
	// fmt.Println(nums[1:])
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	for i, iValue := range nums {
		for j, jValue := range nums[(i + 1):] {
			if iValue+jValue == target {
				res := []int{i, j + i + 1}
				return res
			}
		}
	}
	return nums
}
