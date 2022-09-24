package main

import "fmt"

func main() {
	a := []int{
		0, 2, 1, 1, 2, 2, 1, 1, 0, 0, 0,
	}
	sortColors(a)
	fmt.Println(a)
	aa := []int{
		2, 0, 2, 1, 1, 0,
	}
	sortColors2(aa)
	fmt.Println(aa)
}

func sortColors(nums []int) {
	start := 0
	end := len(nums) - 1
	for i := 0; i <= end; {
		if nums[i] == 0 {
			nums[start], nums[i] = nums[i], nums[start]
			start++
			i++
		} else if nums[i] == 2 {
			nums[i], nums[end] = nums[end], nums[i]
			end--
		} else {
			i++
		}

	}
}
func sortColors2(nums []int) {
	left := 0
	right := len(nums) - 1
	// case: nums[right] == 0
	// i <= right, not i < right
	for i := 0; i <= right;{
		if nums[i] == 0 {
			nums[i], nums[left] = nums[left], nums[i]
			left++
            i++
		} else if nums[i] == 2 {
			nums[i], nums[right] = nums[right], nums[i]
			right--
        } else {
            i++
        }
	}
}

