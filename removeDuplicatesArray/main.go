package main

import (
	"fmt"
)

func main() {
	// a1 := []int{1, 2, 3}
	// fmt.Println(removeDuplicates(a1), a1)

	a2 := []int{1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(a2), a2)

}

func removeDuplicates(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	n := 1
	j := 1	
	// 1, 1, 1, 2, 3
	// 1, 2, 3
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] != nums[i+1] {
			nums[j] = nums[i+1]
			j++
			n++
		}
	}
	return n

}
