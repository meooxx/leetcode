package main

import "fmt"

func main() {
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
	fmt.Println(canJump([]int{0}))
	fmt.Println(canJump([]int{1, 2}))
	fmt.Println(canJump([]int{0, 1, 2}))
	fmt.Println(canJump([]int{3, 0, 8, 2, 0, 0, 1}))
}

func canJump(nums []int) bool {
	// {0}
	if len(nums) == 1 {
		return true
	}
	maxRight := 0
	// end := 0
	for i := 0; i < len(nums)-1; i++ {
		if i > maxRight {
			return false
		}
		if nums[i]+i > maxRight {
			maxRight = i + nums[i]
		}
		if maxRight >= len(nums)-1 {
			return true
		}
		// if i == end {
		// 	end = maxRight
		// }

	}

	return false
}
