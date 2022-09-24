package main

import "fmt"

//
// 遇到右括号,推出左括号 计算此时长度 
// stack 为空的时候, 直接推一个作为起始位置index      
//    | -  - - - - |
//    | -------    | 
//    | |     |    |
// -1 ( ( (   )    )    )
func longestValidParentheses(s string) int {
	stack := []int{-1}
	maxLen := 0
	for index, v := range s {
		if v == '(' {
			stack = append(stack, index)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				// 第一位有效括号左边一个元素index
				stack = append(stack, index)
				continue
			}
			preIndex := stack[len(stack)-1]
			if index-preIndex > maxLen {
				maxLen = index - preIndex
			}
		}
	}
	return maxLen
}

func main() {
	fmt.Println(longestValidParentheses("()()"))
	fmt.Println(longestValidParentheses(")()())"))

	fmt.Println(longestValidParentheses("()(())"))
	fmt.Println(longestValidParentheses("))))"))

}
