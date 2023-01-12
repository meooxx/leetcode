package main

func containsDuplicate(nums []int) bool {
	mp := map[int]bool{}
	for i := range nums {
		if mp[nums[i]] {
			return true
		}
		mp[nums[i]] = true
	}
	return false
}
