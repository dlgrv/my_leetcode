package main

import "fmt"

func main() {
	nums := []int{-2,1,-3,4,-1,2,1,-5,4}
	fmt.Println(maxSubArray(nums))
}

func maxSubArray(nums []int) int {
    maxSum, tempMax := nums[0], nums[0]
    for i := 1; i < len(nums); i++ {
        tempMax += nums[i]
		tempMax = max(tempMax, nums[i])
		maxSum = max(maxSum, tempMax)
    }
    return maxSum
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}