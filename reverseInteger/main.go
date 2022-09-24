package main

import (
	"fmt"
	"strconv"
	"strings"
)

// -102
func reverse(x int) int {
	s := fmt.Sprint(x)
	r := ""
	if x < 0 {
		s = s[1:]
	}
	digitals := strings.Split(s, "")
	for i := len(digitals) - 1; i >= 0; i-- {
		// 101000
		if r == "" && digitals[i] == "0" {
			continue
		} else {
			r += digitals[i]
		}
	}
	num, err := strconv.ParseInt(r, 10, 32)

	if err != nil {
		return 0
	}
	if x < 0 {
		return -int(num)
	}
	return int(num)
}
func main() {

	n := reverse(10100)
	fmt.Println(n)
}
