

package main

func countBits(n int) []int {
	dp := make([]int, n+1)
	for i:=1;i<=n;i++ {
		dp[i] = dp[i/2] + i&1
	}
	return dp

}