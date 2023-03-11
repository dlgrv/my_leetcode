package main

import (
	"sort"
	"strconv"
)

func main() {
	a := []int{10, 3, 8, 9, 4}
	findRelativeRanks(a)
}

func findRelativeRanks(score []int) []string {
	sortedScore := []int{}
	ansValues := []string{}
	ans := []string{}

	for _, val := range score {
		sortedScore = append(sortedScore, val)
	}
	sort.Ints(sortedScore)

	for i := 0; i < len(score); i++ {
		if i == len(score)-3 {
			ansValues = append(ansValues, "Bronze Medal")
		} else if i == len(score)-2 {
			ansValues = append(ansValues, "Silver Medal")
		} else if i == len(score)-1 {
			ansValues = append(ansValues, "Gold Medal")
		} else {
			ansValues = append(ansValues, strconv.Itoa(len(score)-i))
		}
	}

	for i := 0; i < len(score); i++ {
		for j := 0; j < len(score); j++ {
			if score[i] == sortedScore[j] {
				ans = append(ans, ansValues[j])
			}
		}
	}

	return ans
}
