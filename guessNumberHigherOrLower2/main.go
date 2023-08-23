package main
import "math"

func getMoneyAmount(n int) int {
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	return getMoney(1, n, dp)

}

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

func getMoney(left, right int, dp [][]int) int {
	if left >= right {
		return 0
	}
	if dp[left][right] != 0 {
		return dp[left][right]
	}
	maximum := math.MaxInt
	for i := left; i <= right; i++ {
		maxCost := max(getMoney(left, i-1, dp), getMoney(i+1, right, dp)) + i
		maximum = min(maximum, maxCost)
	}
	dp[left][right] = maximum
	return maximum
}
