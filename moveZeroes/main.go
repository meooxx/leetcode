package main

func moveZeroes(nums []int) {
	size := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			size++
		} else if size > 0 {
			nums[i], nums[i-size] = nums[i-size], nums[i]
		}
	}
}
