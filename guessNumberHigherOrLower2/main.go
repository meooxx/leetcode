package main

import "math"

func getMoneyAmount0(n int) int {
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



//  1  2  3  4   5   
//  1-2  -> 1
//  2-3  -> 2
//  3-4  -> 3
//  4-5  -> 4
// 
// dp[i][i] == 0, e.g. 2-2, 1-1, 3-3 的范围,only 2
// 1-5 =>
// min
//		1+Max(dp[1][0], d[2][5])
//    2+Max(dp[1][1], d[3][5])     ✅
//  	3 + Max(dp[1][2], dp[4][5] )
// 

func getMoneyAmount(n int) int {
	dp := make([][]int, n+2)
	for i := range dp {
		dp[i] = make([]int, n+2)
	}

	for l := 1; l < n; l++ {
		for i := 1; i <= n-l; i++ {
			j := i + l
			dp[i][j] = math.MaxInt
			for k := i; k <= j; k++ {
				dp[i][j] = min(dp[i][j], max(dp[i][k-1], dp[k+1][j])+k)
			}
		}
	}
	return dp[1][n]
}
