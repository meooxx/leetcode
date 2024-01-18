package main

// 1  1
// 2  dp[1] + 1
// 3  dp[2] + 1
// 4  dp[4- 1 * 1] + 1, dp[4- 2 * 2] + 1
// 10  dp[5- 1*1] + 1, dp[5-2*2] + 1, dp[10-3*3] + 1

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = i
	}
	for i := 4; i <= n; i++ {
		for j := 2; j*j <= n; j++ {
			dp[i] = min(dp[i-j*j]+1, dp[i])
		}
	}
	return dp[n]
}
