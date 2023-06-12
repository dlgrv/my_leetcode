package main

import (
	"fmt"
)

// func main() {
// 	a := 1000021
// 	fmt.Println(isPalindrome(a))
// }

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	strNumber := fmt.Sprint(x)
	for strNumber != fmt.Sprint("") {
		if firstDigit(strNumber) != lastDigit(strNumber) {
			return false
		}
		strNumber = separateNumber(strNumber)
	}

	return true
}

func firstDigit(strNumber string) string {
	return strNumber[0:1]
}

func lastDigit(strNumber string) string {
	digitCount := len(strNumber)
	return strNumber[digitCount-1 : digitCount]
}

func separateNumber(number string) string {
	strNumber := fmt.Sprint(number)
	if len(strNumber) <= 1 {
		return fmt.Sprint("")
	}
	fmt.Println(strNumber)
	strNumber = strNumber[1 : len(strNumber)-1]
	return strNumber
}
