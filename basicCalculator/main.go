package main

/*
* sign = 1
* 1 + 2 - 3 +  ( 2 + 3 )
* if s[i] == '+' || '-'
     sum the n before sign
			"-11 + 2" ==> '-', +0 * sign(1), sign = -1 
								==> '+', 11 * -1(sign), sign = 1
	if s[i] =='(', save result and sign, calculate n which in parenthesis
		1-(1+2) stack{1, -1(sign)}
				|==> 3 == result 
	if s[i] == ')', pop stack, get pre result and sign
	    	==> result *= sign(pre) += result(pre)
 */  
func calculate(s string) int {
	sign := 1
	result := 0
	n := 0
	stack := []int{sign}
	for i := range s {
		if s[i] >= '0' && s[i] <= '9' {
			n = n*10 + int(s[i]-'0')
		} else if s[i] == '+' {

			result += n * sign
			n = 0
			sign = 1
		// if s[i] == '+' || s[i] == '-'
		//   result += n *  sign
		//   sign = stack[len(stack) - 1] * (1 | -1 depend on s[i])
		// e.g.
		// - (2+3) ==> -2 -3, 
		// - (2-3) ==> -2+3
		} else if s[i] == '-' {
			result += n * sign
			n = 0
			sign = -1
		} else if s[i] == '(' {
			stack = append(stack, result, sign)
			result = 0
			n = 0
			sign = 1
		} else if s[i] == ')' {
			result += n * sign
			result *= stack[len(stack)-1]
			result += stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			n = 0
		}
	}

	result += n * sign
	return result
}
