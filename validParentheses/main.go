package main

import (
	"fmt"
)

func main() {
	fmt.Println(isValid("{[]}"))
	fmt.Println(isValid("{[]"))
	fmt.Println(isValid("{}"))
	fmt.Println(isValid("){"))

	fmt.Println(isValid("(){}}{"))
}

// case1 "[]"
// case2 "["
// case3 "]"
func isValid(s string) bool {
	stack := []rune{}
	for _, v := range s {
		switch v {
		case '(':
			stack = append(stack, ')')
		case '[':
			stack = append(stack, ']')
		case '{':
			stack = append(stack, '}')
		}

		if v == ')' || v == ']' || v == '}' {
			if len(stack) == 0 || stack[len(stack)-1] != v {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}
	if len(stack) > 0 {
		return false
	}
	return true
}
