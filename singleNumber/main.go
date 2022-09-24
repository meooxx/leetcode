package main

import "fmt"

func main() {
	fmt.Println(singleNumber([]int{1, 2, 2, 1, 5}))
}

func singleNumber(nums []int) int {
	result := 0
	for i := range nums {
		result ^= nums[i]
	}
	return result
}
