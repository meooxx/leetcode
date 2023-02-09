package main

func findPeakElement(nums []int) int {
	// length == 1
	if len(nums) == 1 {
		return 0
	}
	// length == 2
	left, right := 0, len(nums)-1
	if nums[0] > nums[1] {
		return 0
	}
	if nums[right] > nums[right-1] {
		return right
	}

	left++
	right--
	// [1,5,3]
	for left <= right {
		mid := (right-left)/2 + left
		if nums[mid] > nums[mid-1] && nums[mid] > nums[mid+1] {
			return mid
		} else if nums[mid] < nums[mid+1] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
