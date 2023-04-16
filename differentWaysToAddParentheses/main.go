package main

import "strconv"

// @see:https://leetcode.com/problems/different-ways-to-add-parentheses/solutions/1294189/easy-solution-faster-than-100-explained-with-diagrams/?orderBy=most_votes
// partion at *, get leftPart andt rightPart,
//                  for i in left, for j in right, => append(i +-* j)  
// partion at -
// partion at +
// 2     * 3 				-    2 + 1             *          |     *
//       |             	               2        -     | 2      +  
// left        		rgith                 		 3    +   |      -    1
//                  |                           2   1 |    3  2
// 									|
//          left       right
// 

func diffWaysToCompute(expression string) []int {
	isNumber := true
	result := []int{}
	for i := range expression {
		if expression[i] == '+' || expression[i] == '*' || expression[i] == '-' {
			isNumber = false
			leftPart := diffWaysToCompute(expression[:i])
			rightPart := diffWaysToCompute(expression[i+1:])
			for _, left := range leftPart {
				for _, right := range rightPart {
					switch expression[i] {
					case '-':
						result = append(result, left-right)
					case '+':
						result = append(result, left+right)
					case '*':
						result = append(result, left*right)
					}
				}
			}
		}

	}
	if isNumber {
		n, _ := strconv.ParseInt(expression, 10, 64)
		result = append(result, int(n))
	}
	return result
}
