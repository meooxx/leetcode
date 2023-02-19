package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maximumGap([]int{
		3, 9, 6, 1,
	}))
	fmt.Println(maximumGap([]int{1, 1, 1, 1}))
	fmt.Println(maximumGap([]int{
		100, 3, 2, 1,
	}))
}

/*
*
// 	https://leetcode.com/problems/maximum-gap/solutions/50643/bucket-sort-java-solution-with-explanation-o-n-time-and-space/
// 	totalGap = max - min
//  buckets = n - 1  (appointed)
//  bucketsSize = (max - min) / min
//  bucketIndex = (num - minNum) / bucketsSize
//                       1-2 2-3  3-5  5-10
//  1 2 3 5 10 avgGap = [ 1,  1,   2,   5 ] / 4 = x
//  totalGap = Max-min(10 - 1)  = x * gap nums(n-1)
//  x = Max - min / n - 1
// If every gap is smaller or equal to Floor of (max-min)/(n-1),
// then average number won't be (max-min)/(n-1), 
// so there must be a gap larger than Ceil of (max-min)/(n-1)
*/
func maximumGap1(nums []int) int {

	if len(nums) <= 1 {
		return 0
	}
	max := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}
	if len(nums) == 2 {
		return max(nums[1]-nums[0], nums[0]-nums[1])
	}
	minV := nums[0]
	maxV := nums[0]

	min := func(a, b int) int {
		if a <= b {
			return a
		}
		return b
	}
	for i := range nums {
		minV = min(nums[i], minV)
		maxV = max(nums[i], maxV)
	}
	bucketSize := int(math.Ceil(float64(maxV-minV) / float64((len(nums) - 1))))
	buckets := make([][]int, len(nums))
	if bucketSize == 0 {
		return 0
	}
	for i := range nums {
		bIndex := (nums[i] - minV) / bucketSize
		vals := buckets[bIndex]
		if vals == nil {
			vals = []int{math.MaxInt, math.MinInt}
		}
		// 0 < nums[i] < 10**9
		vals[0] = min(vals[0], nums[i])
		vals[1] = max(vals[1], nums[i])
		buckets[bIndex] = vals
	}
	var pre []int
	anwser := 0
	for i := 0; i < len(buckets); i++ {
		if buckets[i] != nil {
			if pre == nil {
				pre = buckets[i]
			} else {
				anwser = max(buckets[i][0]-pre[1], anwser)
				pre = buckets[i]
			}
		}
	}
	return anwser
}

func maximumGap(nums []int) int {

	if len(nums) <= 1 {
		return 0
	}
	max := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}
	if len(nums) == 2 {
		return max(nums[1]-nums[0], nums[0]-nums[1])
	}
	minV := nums[0]
	maxV := nums[0]

	min := func(a, b int) int {
		if a <= b {
			return a
		}
		return b
	}
	for i := range nums {
		minV = min(nums[i], minV)
		maxV = max(nums[i], maxV)
	}
	bucketSize := int(math.Ceil(float64(maxV-minV) / float64((len(nums) - 1))))
	if bucketSize == 0 {
		return bucketSize
	}
	// buckets:    [][min, max]
	//        拆成 []bucketsMin
	//            []bucketsMax
	// buckets := make([][]int, len(nums)-1)
	bucketsMin := make([]int, len(nums))
	bucketsMax := make([]int, len(nums))
	for i := range bucketsMin {
		bucketsMin[i] = math.MaxInt
		bucketsMax[i] = math.MinInt
	}
	for i := range nums {
		// 可以不放置 maxVal 和  minV
		// 就可以使用 n-1个 筒子了
		bIndex := (nums[i] - minV) / bucketSize
		bucketsMin[bIndex] = min(bucketsMin[bIndex], nums[i])
		bucketsMax[bIndex] = max(bucketsMax[bIndex], nums[i])
	}
	preMax := minV
	anwser := 0
	for i := 0; i < len(bucketsMin); i++ {
		if bucketsMin[i] == math.MaxInt || bucketsMax[i] == math.MinInt {
			continue
		}
		anwser = max(bucketsMin[i]-preMax, anwser)
		preMax = bucketsMax[i]
	}
	return anwser
}
