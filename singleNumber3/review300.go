package main

func singleNumber(nums []int) []int {
	mask := 0
	for i := range nums {
		mask ^= nums[i]
	}

	// we got a least one bit mast
	//   11     10
	// ~(10)   ~(1)   ~(n-1)
	//   01   ..10
	//    &
	//   1      10  only has one '1' bit
	//
	mask &= -mask
	result := []int{0, 0}
	for i := range nums {
		//  why == 1 ?
		// cause of mask != 1 or 0
		// if mask == 0b10, 
		// result of nums[i] & mask is either ob10 or 0
		if nums[i]&mask == 0 {
			result[0] ^= nums[i]
		} else {
			result[1] ^= nums[i]
		}

	}
	return result
}
