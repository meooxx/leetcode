package main

// max[i] = max(max[i-2], max[i-1])
func rob(nums []int) int {

	premax1 := 0
	premax2 := 0

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := range nums {
		temp := premax2
		premax2 = max(premax1+nums[i], premax2)
		premax1 = temp
	}
	return premax2

}
