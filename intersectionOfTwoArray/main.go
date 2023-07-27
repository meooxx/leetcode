package main

func intersection(nums1 []int, nums2 []int) []int {
	mp := [1000]int{}

	for i := range nums1 {
		mp[nums1[i]] = 1
	}
	result := []int{}
	for i := range nums2 {
		if mp[nums2[i]] == 1 {
			result = append(result, nums2[i])
			mp[nums2[i]] = 0
		}
	}
	return result
}
