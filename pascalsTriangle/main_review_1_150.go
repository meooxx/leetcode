package main

func generate(numRows int) [][]int {
	dp := make([][]int, numRows)
	for n := 0; n < numRows; n++ {
		dp[n] = make([]int, n+1)
		dp[n][0] = 1
		dp[n][n] = 1
		for i := 1; i < n; i++ {
			dp[n][i] = dp[n-1][i-1] + dp[n-1][i]
		}
	}
	return dp
}
