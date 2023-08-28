package main

//				1   17   5   10  13  15   10   5  16    8
//
// up    1    2    2   4   4   4     4   4   6
// down  1   1     3   3   3   3     5   5   5    7
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
			down[i] = up[i-1] + 1
			up[i] = up[i-1]
		} else {
			up[i] = up[i-1]
			down[i] = down[i-1]
		}
	}
	if up[len(nums)-1] > down[len(nums)-1] {
		return up[len(nums)-1]
	}
	return down[len(nums)-1]

}
