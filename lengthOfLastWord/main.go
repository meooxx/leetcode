package main

import "fmt"

func main() {
	fmt.Println(lengthOfLastWord("hello world"))
}

func lengthOfLastWord(s string) int {
	i := len(s) - 1
	for i >= 0 {
		if s[i] != ' ' {
			break
		}
		i--
	}
	j := i
	for j >= 0 {
		if s[j] == ' ' {
			break
		}
		j--
	}
	return i - j
}
