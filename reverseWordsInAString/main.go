package main

import (
	"fmt"
)

func main() {
	fmt.Println(reverseWords("the word is crazy"))
}

func reverseWords(s string) string {

	left := 0
	right := 0
	var str string
	for left < len(s) {
		for left < len(s) && s[left] == ' ' {
			left++
		}
		right = left
		for right < len(s) && s[right] != ' ' {
			right++
		}
		if right > left {
			if str == "" {
				str = s[left:right]
			} else {
				str = s[left:right] + " " + str
			}

		}
		left = right
	}

	return str
}
