package main

// if target == 0, only 0 or nums[j] == target
// so we got base case dp[0] = 1, only if target == nums[j]
// nums = [1, 2], target 3
// dp[i]
// target         1,    2,    3 
//      dp[1] = 1, dp[2] = 1
// 			dp[3] = dp[3-1] + dp[3-2] = dp[2] + dp[1]
func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for j := range nums {
			if i-nums[j] >= 0 {
				dp[i] += dp[i-nums[j]]
			}
		}
	}
	return dp[target]
}
