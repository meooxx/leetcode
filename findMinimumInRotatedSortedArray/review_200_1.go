package main

func findMin(nums []int) int {
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := (right-left)/2 + left
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[right]
}
