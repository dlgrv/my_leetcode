package main

import (
	"fmt"
	"strings"
)

func main() {
	sentence := "leetcode exercises sound delightful"
	fmt.Println(isCircularSentence(sentence))

	sentence = "eetcode"
	fmt.Println(isCircularSentence(sentence))

	sentence = "Leetcode is cool"
	fmt.Println(isCircularSentence(sentence))
}

func isCircularSentence(sentence string) bool {
	words := strings.Fields(sentence)
	for i, _ := range words {
		if i == 0 &&
			words[i][0] != words[len(words)-1][len(words[len(words)-1])-1] {
			return false
		}
		if i != 0 &&
			words[i-1][len(words[i-1])-1] != words[i][0] {
			return false
		}
	}
	return true
}
