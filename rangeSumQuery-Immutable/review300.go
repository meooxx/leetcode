package main

type NumArray struct {
	dp []int
}

func Constructor(nums []int) NumArray {
	dp := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		dp[i+1] = dp[i] + nums[i]
	}
	return NumArray{
		dp,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.dp[right+1] - this.dp[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
