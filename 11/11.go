package main

import "fmt"

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}

	fmt.Println(maxArea(height))
}

func maxArea(height []int) int {
	area := 0
	left := 0
	tempArea := 0
	right := len(height) - 1
	for width := len(height) - 1; width > 0; width-=1 {
		if height[left] < height[right] {
			tempArea = width * height[left]
			left += 1
		} else {
			tempArea = width * height[right]
			right -= 1
		}
		area = max(area, tempArea)
	}

	return area
}

func max(x, y int)int{
    if x >= y {
        return x
    }
    return y
}
