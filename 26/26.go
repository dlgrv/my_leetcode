package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 5, 5, 5, 6, 6, 6}
	fmt.Println(removeDuplicates(a))
}

type Element struct {
	Index int
	Exist bool
}

func removeDuplicates(nums []int) int {
	elementCount := 0

	for i := 0; i < len(nums); i++ {
		if i < len(nums)-1 && nums[i] == nums[i+1] {
			continue
		}
		nums[elementCount] = nums[i]
		elementCount++
	}

	return elementCount
}

// func removeDuplicates(nums []int) int {
// 	presence := map[int]bool{}

// 	for _, value := range nums {
// 		if presence[value] != true {
// 			presence[value] = true
// 		}
// 	}

// 	return len(presence)
// }
