package main

import (
	"fmt"
	"sort"
)

func main() {
	g := []int{10, 9, 8, 7}
	s := []int{5, 6, 7, 8}
	fmt.Println(findContentChildren(s, g))
}

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	i, j := 0, 0
	for i < len(g) && j < len(s) {
		if s[j] >= g[i] {
			i++
		}
		j++
	}
	return i
}
