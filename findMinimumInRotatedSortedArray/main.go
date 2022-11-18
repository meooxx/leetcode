package main

// 4  5  1  2  3, mid < 3, 4-5-1
// 4  5  1
// 1
func findMin(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		midIndex := left + (right-left)/2
		if nums[midIndex] > nums[right] {
			left = midIndex + 1
		} else {
			right = midIndex
		}
	}
	return nums[left]
}
