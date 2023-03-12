package main

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	pre, prepre := 0, 0
	maxV := nums[0]
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	// 0 0 | 1 2 3
	for i := 0; i < len(nums)-1; i++ {
		tmp := pre
		pre = max(pre, prepre+nums[i])
		prepre = tmp
	}
	maxV = pre
	pre, prepre = 0, 0
	for i := 1; i < len(nums); i++ {
		tmp := pre
		pre = max(pre, prepre+nums[i])
		prepre = tmp
	}
	maxV = max(maxV, pre)
	return maxV
}
