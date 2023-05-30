package main

type NumArray struct {
	val []int
}

// nums  0        1    4
//
//	nums[i]
//	  	2         3    7

// sum[i] = num[i-1] + num[i]
func Constructor(nums []int) NumArray {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	return NumArray{
		val: nums,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return this.val[right]
	}
	// nums[right] = sum(0..right)
	// nums[left...right] = sum(0..right) - sum(0..left-1)
	return this.val[right] - this.val[left-1]
}

/**
* Your NumArray object will be instantiated and called as such:
* obj := Constructor(nums);
* param_1 := obj.SumRange(left,right);
 */
