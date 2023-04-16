package main

func productExceptSelf(nums []int) []int {

	//        2            5        4         3
	// left   1           1*2     1*2*5     1*2*5 *4
	// right  5*4*3*1     4*3*1      3*1         1
	result := make([]int, len(nums))
	left := 1
	right := 1
	for i := 0; i < len(nums); i++ {
		if i > 0 {
			left *= nums[i-1]
		}
		result[i] = left
	}

	for i := len(nums) - 1; i >= 0; i-- {
		if i < len(nums)-1 {
			right *= nums[i+1]
		}
		result[i] *= right
	}
	return result

}
