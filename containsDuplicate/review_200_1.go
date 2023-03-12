package main

func containsDuplicate(nums []int) bool {
	mp := map[int]int{}
	for _, v := range nums {
		if mp[v] > 0 {
			return true
		}
		mp[v] = 1
	}
	return false

}
