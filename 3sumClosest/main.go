package main

import (
	"fmt"
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Slice(nums, func(a, b int) bool {
		return nums[a] < nums[b]
	})
	closestNum := math.MaxInt32
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j, k := i+1, len(nums)-1; j < k; {
			sum := nums[i] + nums[j] + nums[k]
			if math.Abs(float64(sum-target)) < math.Abs(float64(closestNum-target)) {
				closestNum = sum
			}
			if sum > target {
				k--
				// next 是否重复
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			} else {
				j++
				for j < k && nums[j] == nums[j-1] {
					j++
				}
			}

		}
	}
	return closestNum

}

func main() {
	n := threeSumClosest([]int{-3, 0, 1, 2}, 1)
	fmt.Println(n)
}
