package main

import "fmt"

func main() {
	fmt.Println(addBinary("11", "1"))
}

func addBinary(string1, string2 string) string {

	if len(string1) > len(string2) {
		return addBinary(string2, string1)
	}

	diff := len(string2) - len(string1)
	for i := 0; i < diff; i++ {
		string1 = "0" + string1
	}

	carry := false

	answer := ""

	for i := len(string1) - 1; i >= 0; i-- {
		if string1[i] == '1' && string2[i] == '1' {
			if carry == true {
				answer = "1" + answer
			} else {
				answer = "0" + answer
				carry = true
			}
		} else if string1[i] == '0' && string2[i] == '0' {
			if carry == true {
				answer = "1" + answer
				carry = false
			} else {
				answer = "0" + answer
			}
		} else if string1[i] != string2[i] {
			if carry == true {
				answer = "0" + answer
			} else {
				answer = "1" + answer
			}
		}
	}

	if carry {
		answer = "1" + answer
	}

	return answer
}
