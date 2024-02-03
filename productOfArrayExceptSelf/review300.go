package main

func productExceptSelf(nums []int) []int {

	left := 1
	right := 1
	result := make([]int, len(nums))
	result[0] = 1
	//  4     2     3     5
	
	//        i         ->  first loop from 1
  //  0     4  	 4*2	 4*2*3							result
	//                    i  <- second loop from right-1
  //             4*2*5  4*2*3*1
	// result[i] = Acc[0:i-1] * Acc[i+1:right-1]
	for i := 1; i < len(nums); i++ {
		left *= nums[i-1]
		result[i] = left
	}

	for i := len(nums)-1;i>=0;i--{
		
		result[i] *= right
		right *= nums[i]
	}
	return result

}
