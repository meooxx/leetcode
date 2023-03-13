package main

import "fmt"

func main() {
	fmt.Println(containsNearbyAlmostDuplicate([]int{1, 2, 3, 1}, 3, 0))
	fmt.Println(containsNearbyAlmostDuplicate([]int{-1, 5, 10, -5, -1}, 2, 3))
}

//                   indexDiff     valDiff
//      	     				  2					  1000
// 1           0
// 10000       0
// 20000       0
// 30000       0
// 2           0
// 5           0
// 10          0

func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {
	if len(nums) < 2 {
		return false
	}
	bucketMp := map[int]int{}
	// ??  valueDiff + 1
	width := valueDiff + 1
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	for i := range nums {
		bucket := nums[i] / width
		// 1 2 3    ==>  0
		// -1 -2 -3 should in  -1
		if nums[i] < 0 {
			// -1 / 2 ==> -1
			bucket--
		}
		if _, exists := bucketMp[bucket]; exists {
			return true
		} else {
			if next, ok := bucketMp[bucket+1]; ok && abs(next-nums[i]) <= valueDiff {
				return true
			}
			if pre, ok := bucketMp[bucket-1]; ok && abs(pre-nums[i]) <= valueDiff {
				return true
			}
		}
		bucketMp[bucket] = nums[i]
		// whyyyy not considering nums[i] < 0?
		// if use blow case, it will fail
		// -1 10 20 -1
		if i >= indexDiff {
			bucket := nums[i-indexDiff] / width
			// I think it required 
			// but the leetcode cases can be passed
			if nums[i-indexDiff] < 0 {
				bucket--
			}
			delete(bucketMp, bucket)
		}
	}
	return false
}
