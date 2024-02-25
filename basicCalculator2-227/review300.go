package main

func calculate(s string) int {
	s += "+"
	stack := []int{}
	lastSign := byte('+')
	n := 0
	for i := range s {
		if s[i] >= '0' && s[i] <= '9' {
			n = n*10 + int(s[i]-'0')
		} else if s[i] == '+' || s[i] == '-' || s[i] == '*' || s[i] == '/' {
			switch lastSign {
			case '+':
				{
					stack = append(stack, n)
				}
			case '-':
				{
					stack = append(stack, -n)
				}
			case '*':
				{
					pre := stack[len(stack)-1]
					stack = stack[:len(stack)-1]
					stack = append(stack, n*pre)
				}
			case '/':
				{
					pre := stack[len(stack)-1]
					stack = stack[:len(stack)-1]
					stack = append(stack, pre/n)
				}
			}
			n = 0
			lastSign = s[i]

		}
	}
	result := 0
	for i := range stack {
		result += stack[i]
	}
	return result
}
