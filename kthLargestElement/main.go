package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println(findKthLargest([]int{1, 2, 3, 4}, 2))
	// fmt.Println(findKthLargest([]int{1, 4, 5, 6}, 3))
	// fmt.Println(findKthLargest([]int{1, 4, 5, 4}, 3))
	// fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
	// fmt.Println(findKthLargest([]int{2, 1}, 1))
	fmt.Println(findKthLargest([]int{99, 99}, 1))

}

func findKthLargest1(nums []int, k int) int {
	maxN := map[string]bool{}
	m := nums[0]
	for k > 0 {
		m = math.MinInt
		maxKey := ""
		for i := range nums {
			key := fmt.Sprint(nums[i], "#", i)
			if nums[i] > m && !maxN[key] {
				m = nums[i]
				maxKey = key
			}
		}
		maxN[maxKey] = true
		k--
	}
	return m
}

// quick select
func findKthLargest(nums []int, k int) int {
	var quickSelect func(nums []int, left, right, count int) int
	// 1 2 3
	quickSelect = func(nums []int, left, right int, count int) int {
		pivot := nums[right]
		i, j := left, right-1
		// i == j?
		// 1, 1 => j-- => i <= j == false
		// size == 2
		for i <= j {
			// case: i == j, e.g. 1,2
			for i <= j && nums[i] < pivot {
				i++
			}
			for j >= i && nums[j] >= pivot {
				j--
			}
			if i <= j {
				nums[i], nums[j] = nums[j], nums[i]
				i++
			} else {
				break
			}
		}
		nums[right], nums[i] = nums[i], nums[right]
		// 1 6 3 5
		// 1 3 5 6
		//     2    len==4
		size := right - i + 1

		if size == count {
			return nums[i]
		}
		if size > count {
			return quickSelect(nums, i+1, right, count)
		} else {
			return quickSelect(nums, left, j, count-size)
		}
	}

	return quickSelect(nums, 0, len(nums)-1, k)
}
