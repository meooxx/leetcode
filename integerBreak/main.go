package main

func integerBreak0(n int) int {
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	product := 1

	for n > 4 {
		product *= 3
		n -= 3
	}
	product *= n
	return product
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	// n = 8,
	// n can be broken into 2 parts
	// 1 * 7, 2 * 6, 3 * 5, 4 * 4, in the form of n = a * b,
	// but a and b can be further broken if a, b >= 4
	// so we can get maxProduct = max(j, dp[4]) * max(i-j, dp[i-j])
	for i := 2; i <= n; i++ {
		for j := 1; j <= i/2; j++ {
			// eg. j == 3, dp[3] < 3
			a := max(j, dp[j])
			// j = 1, dp[i-1] == dp[7] == 0, b = 7, max(7, dp[7])
			b := max(i-j, dp[i-j])
			dp[i] = max(dp[i], a*b)
		}
	}
	return dp[n]
}
