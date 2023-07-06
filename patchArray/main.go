package main

// 1        missing [0,1)
// 1 2  16 ,    32 
// nums[0] <= 1, missing += 1, [0, 2)
// nums[1] <= 2, missing += 2  [0, 4)
//                                add4,   1, 2, 4
// nums[2] >= 4, missing+=missing [0, 8)
//                                add 8, 1,2,4,8
// nums[2] >= 8, missing += missing [0, 16)
func minPatches(nums []int, n int) int {
	added := 0
	missing := 1
	i := 0
	for missing <= n {
		if i < len(nums) && nums[i] <= missing {
			missing += nums[i]
			i++
		} else {
			missing += missing
			added++
		}
	}
	return added
}
