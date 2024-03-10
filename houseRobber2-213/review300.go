package main

// 0 1 | 1 4 5  2
//
//	max[i] = max(maxVals[i-2] + nums[i], maxVals[i-1])
func rob(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}

	preMax1 := 0
	preMax2 := 0

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := 0; i < len(nums)-1; i++ {
		temp := preMax2
		preMax2 = max(preMax1+nums[i], preMax2)
		preMax1 = temp
	}

	maxV := preMax2
	preMax1 = 0
	preMax2 = 0
	for i := 1; i < len(nums); i++ {
		temp := preMax2
		preMax2 = max(preMax1+nums[i], preMax2)
		preMax1 = temp
	}
	return max(preMax2, maxV)
}
