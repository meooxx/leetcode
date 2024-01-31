package main

func maxSlidingWindow(nums []int, k int) []int {
	// store index of max val which is in window
	stack := []int{}
	result := []int{}
	for i := range nums {
		for s := len(stack) - 1; s >= 0 && nums[stack[len(stack)-1]] < nums[i]; s-- {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
		//   1     2       4   6
		// 				i-k  i-k+1   i      k==2
		if i-k+1 > stack[0] {
			stack = stack[1:]
		}
		//   i-3+1(0)         i(2)
		//  [ 1         2      4 ] 6
		if i-k+1 >= 0 {
			result = append(result, nums[stack[0]])
		}
	}
	return result
}
