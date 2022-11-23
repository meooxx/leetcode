package main

// require: marjority > n/2
// https://leetcode.com/problems/majority-element/solutions/51613/o-n-time-o-1-space-fastest-solution/?orderBy=most_votes
func majorityElement(nums []int) int {
	marjority := -1
	count := 0
	for i := range nums {
		if count == 0 {
			marjority = nums[i]
		}
		if marjority == nums[i] {
			count++
		} else {
			count--
		}
	}
	return marjority
}
