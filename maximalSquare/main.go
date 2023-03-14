package main

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func minThree(a, b, c int) int {
	return min(min(a, b), c)
}
func maximalSquare(matrix [][]byte) int {
	dp := make([][]int, len(matrix)+1)

	for i := range dp {
		dp[i] = make([]int, len(matrix[0])+1)
	}
	maxV := 0
	for row := range matrix {
		for col := range matrix[row] {
			if matrix[row][col] == '1' {
				v := minThree(dp[row][col], dp[row][col+1], dp[row+1][col]) + 1
				dp[row+1][col+1] = v
				if v > maxV {
					maxV = v
				}
			}

		}
	}
	return maxV * maxV

}
