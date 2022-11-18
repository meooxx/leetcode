package main

func findMin(nums []int) int {
	left := 0
	right := len(nums) - 1
	// 1 1 1 2 3
	// 3 3 3 1 3
	for left < right {
		midIndex := left + (right-left)/2

		if nums[midIndex] < nums[right] {
			right = midIndex
		} else if nums[midIndex] > nums[right] {
			left = midIndex + 1
		} else {
			right--
		}

	}
	return nums[left]

}
