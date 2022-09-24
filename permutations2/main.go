package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(permuteUnique([]int{1, 1, 2}))
	fmt.Println(permuteUnique([]int{3, 3, 0, 3}))
}

func permuteUnique(nums []int) [][]int {
	result := [][]int{}
	sort.Ints(nums)
	place(&result, []int{}, nums)
	return result
}

func place(result *[][]int, current []int, nums []int) {
	if len(nums) == 0 {
		*result = append(*result, current)
		return
	}
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		remainder := append([]int{}, nums[0:i]...)
		remainder = append(remainder, nums[i+1:]...)

		c := make([]int, len(current))
		copy(c, current)
		c = append(c, nums[i])
		place(result, c, remainder)

	}
}
