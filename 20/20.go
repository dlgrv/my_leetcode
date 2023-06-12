package main

import (
	"fmt"
	"strings"
)

// func main() {
// 	s := "[]{}(())"
// 	fmt.Println(isValid(s))
// }

func isValid(s string) bool {
	characterMonitor := ""
	character := ""
	for i := 0; i < len(s); i++ {
		character = s[i : i+1]
		if character == "{" || character == "[" || character == "(" {
			characterMonitor = fmt.Sprint(characterMonitor + character)
		}
		if len(characterMonitor) == 0 {
			return false
		}
		if character == "}" {
			if lastSymbol(characterMonitor) == fmt.Sprint("{") {
				characterMonitor = removeLastSymbol(characterMonitor, "{")
			} else {
				return false
			}
		}
		if character == "]" {
			if lastSymbol(characterMonitor) == fmt.Sprint("[") {
				characterMonitor = removeLastSymbol(characterMonitor, "[")
			} else {
				return false
			}
		}
		if character == ")" {
			if lastSymbol(characterMonitor) == fmt.Sprint("(") {
				characterMonitor = removeLastSymbol(characterMonitor, "(")
			} else {
				return false
			}
		}
		fmt.Println(characterMonitor)
	}

	if characterMonitor == fmt.Sprint("") {
		return true
	}
	return false
}

func lastSymbol(str string) string {
	symbolsCount := len(str)
	return str[symbolsCount-1 : symbolsCount]
}

func removeLastSymbol(str string, symbol string) string {
	return strings.TrimSuffix(str, symbol)
}
