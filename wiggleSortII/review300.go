package main

import "sort"

func wiggleSort(nums []int) {
	coped := make([]int, len(nums))
	copy(coped, nums)
	sort.Slice(coped, func(a, b int) bool {
		return coped[a] < coped[b]
	})
	i := len(nums) - 1
	// 1 2 3 4 5
	//   5   4
	// 3   2   1

	for index := 1; index < len(nums); index += 2 {
		nums[index] = coped[i]
		i--
	}
	for index := 0; index < len(nums); index += 2 {
		nums[index] = coped[i]
		i--
	}

}
