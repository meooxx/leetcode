package main

func majorityElement(nums []int) int {
	count := 0
	result := -1
	for i := range nums {
		if count == 0 {
			result = nums[i]
		}
		if result == nums[i] {
			count++
		} else {
			count--
		}
	}
	return result
}
