package main

func intersect(nums1 []int, nums2 []int) []int {
	mp := [1001]int{}

	for i := range nums1 {
		mp[nums1[i]]++
	}

	result := []int{}
	for i := range nums2 {
		if mp[nums2[i]] > 0 {
			result = append(result, nums2[i])
			mp[nums2[i]]--
		}
	}

	return result

}
