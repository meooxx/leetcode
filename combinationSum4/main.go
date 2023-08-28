package main

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	for i := range dp {
		dp[i] = -1
	}
	dp[0] = 0
	return sum(nums, target, dp)
}

func sum(nums []int, target int, dp []int) int {
	if target < 0 {
		return 0
	}
	if target == 0 {
		return 1
	}
	if dp[target] != -1 {
		return dp[target]
	}

	anwser := 0

	for i := range nums {
		anwser += sum(nums, target-nums[i], dp)
	}
	dp[target] = anwser
	return anwser
}
