package main

// 1 1
//   1 1
// 1 2 1
//   1 2 1
// 1 3 3 1
func getRow(rowIndex int) []int {
	dp := []int{1}
	if rowIndex == 0 {
		return dp
	}
	for i := 1; i <= rowIndex; i++ {
		for index := len(dp) - 1; index >= 1; index-- {
			dp[index] += dp[index-1]
		}
		dp = append(dp, 1)
	}
	return dp

}