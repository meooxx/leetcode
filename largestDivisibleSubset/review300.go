package main
import "sort"

//  2  3  4  9  24
//  1  1  2  2  3    								 dp 
// -1 -1  0  1  2(max of 1 and 2)    cIndex
func largestDivisibleSubset(nums []int) []int {
	dp := make([]int, len(nums))
	cIndex := make([]int, len(nums))
    sort.Slice(nums, func(a, b int) bool {
        return nums[a] < nums[b]
    })
	for i := range dp {
		dp[i] = 1
		cIndex[i] = -1
	}

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
				cIndex[i] = j
			}
		}
	}
	maxArr := []int{}
	maxCIndex := 0
	for i := range dp {
		if dp[i] > dp[maxCIndex] {
			maxCIndex = i
		}
	}
	for maxCIndex != -1 {
		maxArr = append(maxArr, nums[maxCIndex])
        maxCIndex = cIndex[maxCIndex]
	}
	return maxArr

}