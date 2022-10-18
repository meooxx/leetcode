package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(evalRPN([]string{
		"2", "1", "+", "3", "*",
	}))
}

func evalRPN(tokens []string) int {
	multi := func(a, b int) int {
		return a * b
	}
	div := func(a, b int) int {
		return a / b
	}
	sum := func(a, b int) int {
		return a + b
	}
	minus := func(a, b int) int {
		return a - b
	}

	mp := map[string]func(a, b int) int{
		"*": multi,
		"+": sum,
		"-": minus,
		"/": div,
	}
	stack := []int{}
	for _, vStr := range tokens {

		switch vStr {
		case "*", "+", "-", "/":
			length := len(stack)

			v1 := stack[length-2]
			v2 := stack[length-1]
			r := mp[vStr](v1, v2)
			stack = stack[:length-2]
			stack = append(stack, r)

		default:

			{
				parsedV, _ := strconv.ParseInt(vStr, 10, 32)
				v := int(parsedV)
				stack = append(stack, v)
			}

		}
	}
	return stack[0]
}
