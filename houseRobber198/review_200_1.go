package main

//                 3  [4]
//             2  [3]
//       2    [2]
// pre1 pre0
// 0     2     1   1   2
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	pre1 := 0
	pre0 := nums[0]
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := 1; i < len(nums); i++ {
		pre := pre0
		pre0 = max(pre0, pre1+nums[i])
		pre1 = pre
	}
	return pre0
}
