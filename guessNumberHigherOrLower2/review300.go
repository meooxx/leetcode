package main
import "math"
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func getMoneyAmount(n int) int {
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	return getMoney(1, n, &dp)
}

func getMoney(start, end int, cache *[][]int) int {
	dp := *cache
	if start >= end {
		return 0
	}
	if dp[start][end] != 0 {
		return dp[start][end]
	}
	cost := math.MaxInt
	for k := start; k <= end; k++ {
		cost = min(cost, k+max(getMoney(start, k-1, cache), getMoney(k+1, end, cache)))

	}
	dp[start][end] = cost
	return cost
}
