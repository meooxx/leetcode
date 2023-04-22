package main

func missingNumber1(nums []int) int {

	for i := range nums {
		for i != nums[i] && nums[i] < len(nums) {
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}
	for i := range nums {
		if nums[i] != i {
			return i
		}
	}
	return len(nums)
}
/*
* 
* 1   3 4  
* 1 2 3 4
* ^
* 0 2 0 0 =>2 
*/
func missingNumber(nums []int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		result ^= i ^ nums[i]
	}
	return result ^ len(nums)
}
