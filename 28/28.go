package main

import "fmt"

func main() {
	fmt.Println(strStr("svaadbutssa", "sa"))
}

func strStr(haystack string, needle string) int {
	for i := 0; i <= (len(haystack) - len(needle)); i++ {
		fmt.Println(haystack[i : i+len(needle)])
		if needle == haystack[i:i+len(needle)] {
			return i
		}
	}

	return -1
}
