package main

func singleNumber(nums []int) int {
	anwser := 0
	for i := range nums {
		anwser ^= nums[i]
	}
	return anwser
}
