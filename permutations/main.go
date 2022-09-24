package main

import "fmt"

func main() {
	fmt.Println(permute([]int{1, 2, 0}))
}

func permute(nums []int) [][]int {
	result := [][]int{}
	place(&result, []int{}, nums)
	return result
}

func place(result *[][]int, current []int, nums []int) {
	if len(nums) == 0 {
		*result = append(*result, current)
		return
	}
	for i := 0; i < len(nums); i++ {
		remainder := append([]int{}, nums[0:i]...)
		remainder = append(remainder, nums[i+1:]...)

		c := make([]int, len(current))
		copy(c, current)
		c = append(c, nums[i])
		place(result, c, remainder)

	}
}
