package main

// base on rob 1
// we just divide the this problem to 2 sub problem
// 1, 0-len-1   rob1
// 2, 1-len     rob2
// return max(rob1, rob2)
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	pre := 0
	prepre := 0
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	// 1   2   3
	// pp pre
	for i := 0; i < len(nums)-1; i++ {
		tmp := pre
		pre = max(prepre+nums[i], pre)
		prepre = tmp
	}
	maxV := pre
	pre = 0
	prepre = 0
	for i := 1; i < len(nums); i++ {
		tmp := pre
		pre = max(prepre+nums[i], pre)
		prepre = tmp
	}
	return max(maxV, pre)
}