package main

import (
	"fmt"

	"strings"
)

//  0      0
//  1   1  1
/// 2 2    2
//  3      3

func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}
	r := make([]string, numRows)
	step := 1
	row := 0
	for _, v := range s {
		r[row] += string(v)
		row += step
		if row == numRows {
			step = -1
			row = numRows - 2
		}
		if row == 0 {
			step = 1
		}
	}
	return strings.Join(r, "")
}

func main() {
	v := convert("abcabcabc", 2)
	fmt.Println(v)
}
