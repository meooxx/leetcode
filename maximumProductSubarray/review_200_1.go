package main

func maxProduct(nums []int) int {
	maxV := nums[0]
	pre := 1
	for i := 0; i < len(nums); i++ {
		if pre*nums[i] > maxV {
			maxV = pre * nums[i]
		}
		pre *= nums[i]
		if pre == 0 {
			pre = 1
		}
	}
	pre = 1
	for i := len(nums) - 1; i >= 0; i-- {
		if pre*nums[i] > maxV {
			maxV = pre * nums[i]
		}
		pre *= nums[i]
		if pre == 0 {
			pre = 1
		}
	}
	return maxV
}
