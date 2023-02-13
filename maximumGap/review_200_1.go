package main

import "math"

func maximumGap(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	if len(nums) == 2 {
		return nums[1] - nums[0]
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	minV := nums[0]
	maxV := nums[0]
	for i := range nums {
		minV = min(minV, nums[i])
		maxV = max(maxV, nums[i])
	}

	bucketSize := int(math.Ceil(float64(maxV-minV) / float64(len(nums)-1)))

	if bucketSize == 0 {
		return 0
	}
	// n - 1 buckets
	minBucket := make([]int, len(nums)-1)
	maxBucket := make([]int, len(nums)-1)
	for i := range maxBucket {
		minBucket[i] = math.MaxInt
		maxBucket[i] = math.MinInt
	}

	// 1 3 6
	for i := range nums {
		if nums[i] == maxV || nums[i] == minV {
			continue
		}
		index := (nums[i] - minV) / bucketSize
		minBucket[index] = min(nums[i], minBucket[index])
		maxBucket[index] = max(nums[i], maxBucket[index])
	}
	maxGap := 0
	premax := minV
	for i := 0; i < len(nums)-1; i++ {
		if minBucket[i] == math.MaxInt && maxBucket[i] == math.MinInt {
			continue
		}
		maxGap = max(maxGap, minBucket[i]-premax)
		premax = maxBucket[i]
	}

	maxGap = max(maxGap, maxV-premax)
	return maxGap
}
