package main

import (
	"fmt"
)

func main() {
	//fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 2}, 3))

	fmt.Println(search2([]int{1, 0, 1, 1, 1}, 0))
	fmt.Println(search2([]int{1, 3}, 3))
}

func search(nums []int, target int) bool {
	left := 0
	right := len(nums) - 1

	//2,5,6,0,0,1,2
	// 2222 222
	// 4445444
	for left <= right {
		for left < right && nums[left] == nums[left+1] {
			left++
		}
		for left < right && nums[right] == nums[right-1] {
			right--
		}
		mid := (right-left)/2 + left
		if nums[mid] == target {
			return true
		} else if nums[mid] >= nums[left] {
			if target <= nums[mid] && target >= nums[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] <= target && nums[right] >= target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}

	}
	return false
}

func search2(nums []int, target int) bool {
	left := 0
	right := len(nums) - 1

	for left <= right {
		for left < right && nums[left] == nums[left+1] {
			left++
		}
		for right > left && nums[right] == nums[right-1] {
			right--
		}

		mid := (right-left)/2 + left
		if nums[mid] == target {
			return true
		} else if nums[mid] >= nums[left] {
			if target >= nums[mid] && target < nums[mid] {
				right = mid - 1

			} else {
				left = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}

		}

	}
	return false
}
