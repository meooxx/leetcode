package main

func minSubArrayLen(target int, nums []int) int {
	l, r := 0, 0
	sum := 0
	count := len(nums) + 1
	for ; r < len(nums); r++ {
		sum += nums[r]
		for sum >= target {
			c := r - l + 1
			sum -= nums[l]
			l++
			if c < count {
				count = c
			}
		}
	}
	// Not found
	if count == len(nums)+1 {
		return 0
	}
	return count
}
