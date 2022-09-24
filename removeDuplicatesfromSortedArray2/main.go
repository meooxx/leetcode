package main

import "fmt"

func main() {
	// fmt.Println(removeDuplicates([]int{1, 1, 1, 2, 2, 2, 3}))
fmt.Println(removeDuplicates([]int{1, 1, 1, 2, 2, 3}))
}

func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	left := 2
	count := 2
	// 1 1 1 3
	// 1 1 1 4 4 5
	for i := 2; i < len(nums); i++ {
		if nums[i] > nums[left-2] {
			// nums[left] = nums[i]
			left++
			count++
		}
	}
	return count
}
func removeDuplicates2(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	left := 2
	count := 2
	// 1 1 1 3
	// 1 1 1 4 4 5
	for i := 2; i < len(nums); i++ {
		if nums[i] != nums[left-1] || nums[left-1] != nums[left-2] {
			nums[left] = nums[i]
			left++
			count++
		}
	}
	return count
}
