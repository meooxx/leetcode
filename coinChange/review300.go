package main

import "math"

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt
		for _, c := range coins {
			if  i - c < 0 {
				continue
			}
			if dp[i-c] != math.MaxInt {
				dp[i] = min(dp[i], 1 + dp[i-c])
			}
		}
	}
	if dp[amount] == math.MaxInt {
		return -1
	}
	return dp[amount]

}

func min(a, b int) int {
	if a < b {
		return a 
	}
	return b
}