package main

import (
	"fmt"
)

func main() {
	fmt.Println(minimumTotal([][]int{
		{2},
		{3, 4},
		{6, 5, 7}, // min(6+4,6+1), min(5+4, 5+8), min(7+8, 7+3)
		{4, 1, 8, 3},
	}))
}

// 从倒数第二排倒着计算
func minimumTotal(triangle [][]int) int {
	levels := len(triangle) - 1
	dp := triangle[levels]
	for l := levels - 1; l >= 0; l-- {
		nums := triangle[l]
		for i, n := range nums {
			if dp[i] <= dp[i+1] {
				dp[i] = dp[i] + n
			} else {
				dp[i] = dp[i+1] + n
			}
		}
	}
	return dp[0]
}
func minimumTotal1(triangle [][]int) int {
	dp := []int{0}
	for _, nums := range triangle {
		for index, n := range nums {
			if index > 0 && index < len(dp) {
				if dp[index] <= dp[index-1] {
					nums[index] = dp[index] + n
				} else {
					nums[index] = dp[index-1] + n
				}
			} else if index > 0 {
				nums[index] = dp[index-1] + n
			} else {
				nums[index] = dp[index] + n
			}
		}

		dp = nums
	}
	min := dp[0]

	for _, v := range dp {
		if v < min {
			min = v
		}
	}
	return min
}
func minimumTotal2(triangle [][]int) int {
	if len(triangle) == 1 {
		return triangle[0][0]
	}
	var addNum func(acc, level, index int, t [][]int) int

	addNum = func(acc, level, index int, t [][]int) int {
		if level == len(t) {
			return acc
		}
		nums := t[level]
		sum1 := addNum(acc+nums[index], level+1, index, t)
		sum2 := addNum(acc+nums[index+1], level+1, index+1, t)
		if sum1 <= sum2 {
			return sum1
		}
		return sum2
	}
	return addNum(triangle[0][0], 1, 0, triangle)

}
