package main

import "fmt"

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5}, 3))
}

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (right-left)/2 + left
		midNum := nums[mid]
		if midNum == target {
			return mid
		} else if target < midNum {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left

}
