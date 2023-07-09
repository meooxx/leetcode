package main

import "math"

// c1 for the best candidate for one-cell subsequence
// c2 for the beset candidate for two-cell subsequence
// 1 5 0 6
// results in 0 5 6
// it have  meaning of
//  0 so far the best candidate for 1-cell subsequence,   0
//  5 so far the best candidate for 2-cell subsequence,   1, 5 

// if 1 5 0 2 6 => 0, 2, 6
func increasingTriplet(nums []int) bool {
	c1 := math.MaxInt
	c2 := math.MaxInt

	for i := range nums {
		if nums[i] <= c1 {
			c1 = nums[i]
		} else if nums[i] <= c2 {
			c2 = nums[i]
		} else {
			return true
		}
	}
	return false

}
