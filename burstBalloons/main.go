package main

func maxCoins(nums []int) int {

	dp := make([][]int, len(nums)+2)
	for i := range dp {
		dp[i] = make([]int, len(nums)+2)
	}

	n := []int{1}
	n = append(n, nums...)
	n = append(n, 1)

	return solve(n, 1, len(n)-2, dp)

}

func solve(nums []int, i, j int, mp [][]int) int {
	if i > j {
		return 0
	}
	if mp[i][j] != 0 {
		return mp[i][j]
	}
	maxV := 0
	for last := i; last <= j; last++ {
		leftCoins := solve(nums, i, last-1, mp)
		cur := nums[i-1] * nums[last] * nums[j+1]
		rightCoins := solve(nums, last+1, j, mp)
		maxV = max(maxV, leftCoins+cur+rightCoins)
	}
	mp[i][j] = maxV
	return maxV
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
