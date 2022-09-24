package main

import "fmt"

func main() {
	// fmt.Println(firstMissingPositive2([]int{
	// 	1, 2, 4,
	// }))
	fmt.Println(firstMissingPositive2([]int{
		3, 4, -1, 1,
	}))
	fmt.Println(firstMissingPositive2([]int{
		-1,
	}))
	fmt.Println(firstMissingPositive2([]int{
		0, 1, 2,
	}))
	fmt.Println(firstMissingPositive2([]int{
		1,
	}))
}

// 1 len(nums) = 4, 1-5
// 2 将范围内的数字放到自己位置上去
//  3 , 1, 4,  5,555==>  0 1 2
// 4, 1, 3, 5
// 5, 1,3,4
// 1, 5,3,4
func firstMissingPositive(nums []int) int {
	n := len(nums)

	for i := 0; i < n; i++ {
		for nums[i] > 0 && nums[i] < n && nums[nums[i]-1] != nums[i] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}

// 1,2,4 0
func firstMissingPositive2(nums []int) int {
	nums = append(nums, 0)
	n := len(nums)

	for i := 0; i < n; i++ {
		if nums[i] < 0 || nums[i] >= n {
			nums[i] = 0
		}
	}
	for i := 0; i < n; i++ {
		nums[nums[i]%n] += n
	}
	for i := 1; i < n; i++ {
		if nums[i] < n {
			return i
		}
	}
	return n
}
