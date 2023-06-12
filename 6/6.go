package main

import "fmt"

func main() {
	fmt.Println(convert("PAYPALISHIRING", 4))
}

func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}

	zigzagString := ""
	for row := 1; row <= numRows; row++ {
		i := row
		isZigzagMode := false
		for i-1 < len(s) {
			zigzagString += string(s[i-1])
			if row == 1 || row == numRows {
				i += numRows + (numRows-2)
			} else if isZigzagMode {
				i += row + (row-2)
			} else {
				i += numRows + (numRows-2*row)
			}
			isZigzagMode = !isZigzagMode
		}
	}

	return zigzagString
}
