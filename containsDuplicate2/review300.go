package main

func containsNearbyDuplicate(nums []int, k int) bool {
	
	cache := map[int]int{}

	for i, v := range nums {
		if cache[v] > 0  {
			if  i + 1 - cache[v] <= k {
				return true
			}
		}
		// use 0 as empty value
		cache[v] = i + 1
	}
	return false


}
