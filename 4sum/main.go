package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	var result [][]int
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			// case: 2,2,2,2, 2,2,2
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			for l, r := j+1, len(nums)-1; l < r; {
				sum := nums[i] + nums[j] + nums[l] + nums[r]
				if sum == target {
					result = append(result, []int{nums[i], nums[j], nums[l], nums[r]})
					l++
					r--
					for l < r && nums[l] == nums[l-1] {
						l++
					}
					for l < r && nums[r] == nums[r+1] {
						r--
					}
				} else if sum > target {
					r--
				} else {
					l++
				}
			}
		}
	}
	return result

}

func main() {

	// n := fourSum([]int{1, 0, -1, 0, -2, 2}, 0)
	// fmt.Println(n)

	// n1 := fourSum([]int{2, 2, 2, 2}, 8)
	// fmt.Println(n1)

	n2 := fourSum([]int{-2, -1, -1, 1, 1, 2, 2}, 0)
	fmt.Println(n2)
}
