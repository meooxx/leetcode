package main

func findKthLargest(nums []int, k int) int {

	start := 0

	pivot := len(nums) - 1
	end := pivot - 1

	// 3 2 1 5 6 4,  2
	for start <= end {
		for start <= end &&  nums[start] < nums[pivot] {
			start++
		}
		for start <= end &&  nums[end] > nums[pivot] {
			end--
		}
		if start <= end {
			nums[start], nums[end] = nums[end], nums[start]
			start++
			end--
		}
	}
	nums[start], nums[pivot] = nums[pivot], nums[start]

	// 1 2 3
	if len(nums)-start == k {
		return nums[start]
	} else if len(nums)-start > k {
		return findKthLargest(nums[start+1:], k)
	} else {
		// case: the number of values which are great than midVal is less than k
		return findKthLargest(nums[:start], k-(len(nums)-start))
	}
}
