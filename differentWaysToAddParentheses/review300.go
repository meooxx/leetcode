package main

import "strconv"

// 2 * 3
func diffWaysToCompute(expression string) []int {
	result := []int{}
	isNumber := true
	for i := range expression {
		if expression[i] == '-' || expression[i] == '+' || expression[i] == '*' {
			isNumber = false
			left := diffWaysToCompute(expression[:i])
			right := diffWaysToCompute(expression[i+1:])
			for j := range left {
				for k := range right {
					switch expression[i] {
					case '-':
						result = append(result, left[j]-right[k])
					case '+':
						result = append(result, left[j]+right[k])
					case '*':
						result = append(result, left[j]*right[k])
					}

				}
			}
		}
	}
	if isNumber {
		v, _ := strconv.Atoi(expression)
		result = append(result, v)
	}
	return result

}
