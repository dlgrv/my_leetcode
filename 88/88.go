package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge(nums1, 3, nums2, 3)
	fmt.Println(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	if len(nums1) == 0 && len(nums2) == 0 {
		return
	}
	for i:=0; i<len(nums2) && i<n; i++{
		nums1[m+i] = nums2[i]
	}
	sort.Ints(nums1)
}