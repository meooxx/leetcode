package main

func minSubArrayLen(target int, nums []int) int {
	right := 0
	left := 0
	sum := 0
	count := len(nums) + 1
	for ; right < len(nums); right++ {
		sum += nums[right]
		for sum >= target {
			c := right - left + 1
			sum -= nums[left]
			left++
			if c < count {
				count = c
			}

		}
	}
	if count == len(nums)+1 {
		return 0
	}
	return count
}
