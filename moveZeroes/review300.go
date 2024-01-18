package main

func moveZeroes(nums []int) {
	size := 0

	for i := range nums {
		if nums[i] == 0 {
			size++
		} else if size > 0 {
			// shift left size spaces
			// e.g.	0 1 2 4
			// 		 		1, 2, 4 shift 1 space
			// e.g. 0 1 0 2
			// 				1 shift left 1 space, 2 shift left 2 spaces 
			nums[i], nums[i-size] = nums[i-size], nums[i]
		}
	}

}
