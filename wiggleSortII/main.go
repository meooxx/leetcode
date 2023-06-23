package main

import (
	"fmt"
	"sort"
)

func wiggleSort(nums []int) {
	copied := make([]int, len(nums))
	copy(copied, nums)
	fmt.Println(copied)

	sort.Slice(copied, func(a, b int) bool {
		return copied[a] < copied[b]
	})
	fmt.Println(copied)
	index := len(nums) - 1
	for i := 1; i < len(nums); i += 2 {
		nums[i] = copied[index]
		index--
	}
	for i := 0; i < len(nums); i += 2 {
		nums[i] = copied[index]
		index--
	}
}
