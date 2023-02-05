package main

//  			1   3    5  7  2     6
//  start ^     	 ^ end    count = end - left+1 =3
//  1. end shift right, 	sum+= nums[end],
//	2. for sum >= target 
//		 		count = min(end-start+1, count)
//  	 		start shift right, sum-=nums[start]    	
// 
func minSubArrayLen(target int, nums []int) int {
	start := 0
	end := 0
	sum := 0
	count := len(nums) + 1
	for ; end < len(nums); end++ {
		sum += nums[end]
		for sum >= target {
			if end-start+1 < count {
				count = end - start + 1
			}
			sum -= nums[start]
			start++
		}
	}
	if count == len(nums)+1 {
		return 0
	}
	return count
}
