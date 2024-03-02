package main

func calculate(s string) int {
	lastSign := 1
	stack := []int{}

	n := 0
	result := 0
	for i := range s {
		if s[i] >= '0' && s[i] <= '9' {
			n = n*10 + int(s[i]-'0')
		} else if s[i] == '+' {
			result += n * lastSign
			n = 0
			lastSign = 1
		} else if s[i] == '-' {
			result += n * lastSign
			n = 0
			lastSign = -1
		} else if s[i] == '(' {
			stack = append(stack, result, lastSign)
      // re-init for expression in bracket
			n = 0
      result = 0
      lastSign = 1
		} else if s[i] == ')' {
			preSign := stack[len(stack)-1]
			preResult := stack[len(stack)-2]

			stack = stack[0 : len(stack)-2]
			// 1 calculate expression in bracket, get sum0
			// 2 sum0 by preceding sign of bracket get sum1
			// 3 sum1 + pre result(out of bracket)
			result += n * lastSign
			result = preSign*result + preResult
			n = 0
			// duplicated
			// lastSign = 1
		}
	}
  result += n * lastSign
	return result

}
