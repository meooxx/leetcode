package main

// +1 - 2  + 2 +
//    |       	push(+1)
//      	 |  	push(-2)
// if  + -:  push(+- n)
// if * /: push(pop() * n)
// sum v in stack
func calculate(s string) int {
	n := 0
	stack := []int{}
	sign := byte('+')
	result := 0
	s += "+"
	for i := range s {
		if s[i] >= '0' && s[i] <= '9' {
			n = int(s[i]-'0') + n*10
		} else if s[i] == '+' || s[i] == '-' || s[i] == '*' || s[i] == '/' {
			switch sign {
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
					m := pre * n
					stack[len(stack)-1] = m
				}
			case '/':
				{
					pre := stack[len(stack)-1]
					m := pre / n
					stack[len(stack)-1] = m
				}
			}
			n = 0
			sign = s[i]

		}

	}
	for _, v := range stack {
		result += v
	}
	return result
}
