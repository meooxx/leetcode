package main

func maximalSquare(matrix [][]byte) int {
	dp := make([][]int, len(matrix)+1)
	for i := range dp {
		dp[i] = make([]int, len(matrix[0])+1)
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	minOfThree := func(a, b, c int) int {
		return min(min(a, b), c)
	}
	maxV := 0

	for row := range matrix {
		for col := range matrix[row] {
			if matrix[row][col] == '1' {
				dp[row+1][col+1] = minOfThree(dp[row+1][col], dp[row][col+1], dp[row][col]) + 1
				if dp[row+1][col+1] > maxV {
					maxV = dp[row+1][col+1]
				}
			}
		}
	}
	return maxV * maxV
}
