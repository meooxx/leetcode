package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 3}
	fmt.Println(searchRange(nums, 3))
}

func searchRange(nums []int, target int) []int {
	first := -1
	last := -1
	l, r := 0, len(nums)-1

	for l <= r {
		mid := (r-l)/2 + l
		if nums[mid] == target {
			first, last = mid, mid
			for first > 0 && nums[first-1] == target {
				first--
			}
			for last < len(nums)-1 && nums[last+1] == target {
				last++
			}
			break
		} else if target < nums[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return []int{first, last}
}
