package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{4, 5, 6, 0, 1, 2}
	fmt.Println(search(nums, 0))
	nums1 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(search(nums1, 1))

	nums2 := []int{5, 1, 3}
	fmt.Println(search(nums2, 3))
}

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	midNum := 0
	for left <= right {
		mid := (right-left)/2 + left
		// 左侧升序或者右侧升序
		if nums[mid] >= nums[0] && target >= nums[0] || nums[mid] < nums[0] && target < nums[0] {
			midNum = nums[mid]
			// 在右侧时候不可能等于num[0]
			// 56 7 1 2   7,搜1 不可能等于 5
			// 56 1 23   1 搜6
		} else {
			if target < nums[0] {
				midNum = math.MinInt32
			} else {
				midNum = math.MaxInt32
			}
		}
		if target == midNum {
			return mid
		} else if target > midNum {
			left = mid + 1
		} else {
			right = mid - 1
		}

	}
	return -1

}
