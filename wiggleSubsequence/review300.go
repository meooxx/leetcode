package main

func wiggleMaxLength(nums []int) int {
	up := make([]int, len(nums))
	down := make([]int, len(nums))
	up[0] = 1
	down[0] = 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			up[i] = down[i-1] + 1
			down[i] = down[i-1]
		} else if nums[i] < nums[i-1] {
			up[i] = up[i-1]
			down[i] = up[i-1] + 1
		} else {
			up[i] = up[i-1]
			down[i] = down[i-1]
		}
	}
	return max(up[len(nums)-1], down[len(nums)-1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
