package main

import "fmt"

func main() {
	a := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	// a := [][]int{{-10}}
	fmt.Println(minimumTotal(a))
}

func minimumTotal(triangle [][]int) int {
	return minimumTotalRecur(triangle, 0, 0)
}

// func minimumTotalRecur(triangle [][]int, row int, i int) int {
// 	if row == (len(triangle) - 1) {
// 		return triangle[row][i]
// 	}

// 	a := triangle[row][i] + minimumTotalRecur(triangle, row+1, i)
// 	b := triangle[row][i] + minimumTotalRecur(triangle, row+1, i+1)

// 	return min(a, b)
// }

// // Time Limit Exceeded
func minimumTotalRecur(triangle [][]int, i int) int {
	if len(triangle) == 1 {
		return triangle[0][i]
	}

	a := triangle[0][i] + minimumTotalRecur(triangle[1:], i)
	b := triangle[0][i] + minimumTotalRecur(triangle[1:], i+1)

	return min(a, b)
}

func min(num1, num2 int) int {
	if num1 < num2 {
		return num1
	}
	return num2
}
