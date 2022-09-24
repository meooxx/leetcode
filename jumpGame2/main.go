package main

import "fmt"

func main() {
	fmt.Println(jump([]int{2, 3, 0, 1, 4}))
	fmt.Println(jump([]int{1, 2}))
	fmt.Println(jump([]int{2, 1}))
	fmt.Println(jump([]int{1}))
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
	fmt.Println(jump([]int{1, 2, 1, 1, 1}))
	fmt.Println(jump([]int{3, 4, 3, 2, 5, 4, 3}))
	fmt.Println(jump([]int{2, 3, 1, 2, 1, 8, 3, 1, 1}))

}
// 1 每次本次内循环内选范围内可到达最右边值, 
// 2 设为下次的边界

// 2 1 3 1 2 4
// 2  -> 3 -> 4
func jump(nums []int) int {
	end := 0
	maxEnd := 0
	steps := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i]+i > maxEnd {
			maxEnd = nums[i] + i
		}
		if i == end {
			end = maxEnd
			steps++
		}
	}
	return steps
}
