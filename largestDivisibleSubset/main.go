package main

import "sort"


/**
origin      3  8   4   16
sorted      3  4   8   16

dp          1  1   2    3
childIndex -1 -1   1(4) 2(8)
*/

func largestDivisibleSubset(nums []int) []int {
	sort.Slice(nums, func(a, b int) bool {
		return nums[a] < nums[b]
	})
	dp := make([]int, len(nums))
	childIndex := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
		childIndex[i] = -1
	}
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
				childIndex[i] = j
			}
		}
	}
	maxIndex := 0
	for i := range dp {
		if dp[i] > dp[maxIndex] {
			maxIndex = i
		}
	}

	result := []int{}
	for maxIndex != -1 {
		result = append(result, nums[maxIndex])
		maxIndex = childIndex[maxIndex]
	}
	return result
}
