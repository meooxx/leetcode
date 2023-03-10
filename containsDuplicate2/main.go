package main

func containsNearbyDuplicate(nums []int, k int) bool {
	mp := map[int]int{}

	for i, v := range nums {
		if mp[v] > 0 {
			distance := i + 1 - mp[v]
			if distance <= k {
				return true
			}
		}
		mp[v] = i + 1
	}
	return false
}
