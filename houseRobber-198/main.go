package main
//           1用当前值和前一个不相邻的值和不使用当前值, 取较大的一个           
// max[i] == nums[i] + max[i-2]    >?  max[i-1]
//   1  3  2  5
// 0             max[0]
//   1           max[1]
//     [?]== 3 + max[i-2] >? max[i-1] ==> 3 + max[i-2] ==3

func rob1(nums []int) int {

	maxRob := make([]int, len(nums)+1)
	maxRob[0] = 0
	maxRob[1] = nums[0]
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := 1; i < len(nums); i++ {
		maxRob[i+1] = max(nums[i]+maxRob[i-1], maxRob[i])
	}
	return maxRob[len(nums)]
}
// 可用俩个变量缓存前面的计算结果
func rob(nums []int) int {
	pre0 := 0
	pre1 := nums[0]
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	// 1 2 3
	for i := 1; i < len(nums); i++ {
		temp := pre1
		pre1 = max(nums[i]+pre0, pre1)
		pre0 = temp
	}
	return pre1
}
