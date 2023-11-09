package main

func lengthOfLIS(nums []int) int {
	result := []int{nums[0]}
	for i := range nums {
		if nums[i] > result[len(result)-1] {
			result = append(result, nums[i])
		} else {
			// 1 4  7 8,  5
			l, r := 0, len(result)-1
			for l < r {
        midIndex := (r-l)/2 + l
				if nums[i] <= result[midIndex] {
					r = midIndex
				} else {
					l = midIndex+1
				}
			}
			result[r] = nums[i]

		}
	}
	return len(result)
}