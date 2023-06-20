package main

import "math"

// amount   M for MaxInt
//
//	     1 				 2 								5    conins
//	0    0 				 0 								0
//	1    1 				 M 								M
//	2    2 				 1 								M
//	3    3 				 2 								M
//	4    4 				 2 								M
//	5    5 				-1 								1
//	6    6  3(1+ dp(6-2))   2 == 1 + dp(6-5)
//
// ..
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	// sort.Slice(coins, func(a, b int) bool {
	//     if coins[a] < coins[b] {
	//         return true
	//     }
	//     return false
	// })
	dp[0] = 0
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt
		for j := range coins {
			if i-coins[j] < 0 {
				continue
			}
			if dp[i-coins[j]] != math.MaxInt {
				dp[i] = min(dp[i], 1+dp[i-coins[j]])
			}
		}
	}
	if dp[amount] == math.MaxInt {
		return -1
	}
	return dp[amount]
}
