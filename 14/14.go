package main

// import (
// 	"fmt"
// 	"strings"
// )

func main() {
	var strs []string
	strs = append(strs, "a")
	// strs = append(strs, "flow")
	// strs = append(strs, "flight")
	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	smallestWord := getSmallestWord(strs)
	prefix := ""
	prefixCounter := 0
	for i := 1; i <= len(smallestWord); i++ {
		for _, element := range strs {
			if element[0:i] == smallestWord[0:i] {
				prefixCounter += 1
			}
		}
		if prefixCounter == len(strs) {
			prefix = smallestWord[0:i]
		}
		prefixCounter = 0
	}

	return prefix
}

func getSmallestWord(strs []string) string {
	smallestWord := strings.Repeat("#", 201)
	for _, element := range strs {
		if len(smallestWord) > len(element) {
			smallestWord = element
		}
	}

	return smallestWord
}
