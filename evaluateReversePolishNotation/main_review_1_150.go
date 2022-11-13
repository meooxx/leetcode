package main
import "strconv"
func evalRPN(tokens []string) int {
	stack := []int{}
	operations := map[string]func(a, b int) int{
		"*": func(a, b int) int {
			return a * b
		},
		"+": func(a, b int) int {
			return a + b
		},
		"-": func(a, b int) int {
			return a - b
		},
		"/": func(a, b int) int {
			return a / b
		},
	}

	for i := range tokens {
		t := tokens[i]
		if t == "*" || t == "-" || t == "+" || t == "/" {
			a := stack[len(stack)-2]
			b := stack[len(stack)-1]
			r := operations[tokens[i]](a, b)
			stack = stack[:len(stack)-2]
			stack = append(stack, r)
		} else {
			n, _ := strconv.ParseInt(t, 10, 32)
			stack = append(stack, int(n))

		}
	}
	return stack[len(stack)-1]
}
