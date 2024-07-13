package main

func majorityElement(nums []int) int {
	count := 0
	pre := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] != pre {
			count--
		} else {
			count++
		}
		if count == -1 {
			pre = nums[i]
			count++
		}

	}
	return pre
}
