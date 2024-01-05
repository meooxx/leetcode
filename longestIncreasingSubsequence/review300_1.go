package main

func lengthOfLIS(nums []int) int {
	stack := []int{nums[0]}

	for i := range nums {

		if nums[i] > stack[len(stack)-1] {
			stack = append(stack, nums[i])
		} else {
			left, right := 0, len(stack)-1
			// 1 10 , 2
			//
			for left < right {
				midIndex := (right-left)/2 + left
				if nums[i] <= stack[midIndex] {
					right = midIndex
				} else {
					left = midIndex + 1
				}
			}
			stack[left] = nums[i]
		}

	}

	return len(stack)

}
