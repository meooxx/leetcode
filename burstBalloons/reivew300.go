package main

// [ 1... n ],
// we only need 1 ~ len(nums) -2 space to store coins
func maxCoins(nums []int) int {
	nums = append([]int{1}, nums...)
	nums = append(nums, 1)
	dp := make([][]int, len(nums))
	for i := range dp {
		dp[i] = make([]int, len(nums))
	}

	return solve(nums, 1, len(nums)-2, dp)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func solve(nums []int, i, j int, dp [][]int) int {
	if dp[i][j] != 0 {
		return dp[i][j]
	}
	if i > j {
		return 0
	}
	maxV := 0
	for last := i; last <= j; last++ {
		left := solve(nums, i, last-1, dp)
		// if we pick last, then ballon of i~last-1 will burst before the last
		// so the total coins of cur is nums[i-1] * nums[last] * nums[j+1]
		cur := nums[i-1] * nums[last] * nums[j+1]
		right := solve(nums, last+1, j, dp)
		maxV = max(maxV, left+cur+right)
	}
	dp[i][j] = maxV
	return maxV
}
