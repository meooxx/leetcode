package main

func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	//  8
	// 1 * 7, 2 * 6, 3 * 5, 4 *4
	for i := 2; i <= n; i++ {
		for j := 1; j <= i/2; j++ {
			a := max(j, dp[j])
			b := max(i-j, dp[i-j])
			dp[i] = max(dp[i], a*b)
		}
	}
	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
