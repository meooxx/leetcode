package main

import "fmt"

func main() {
	fmt.Println(maxSubArray([]int{-2, -1}))
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	// 5,4,-1,7,8
	fmt.Println(maxSubArray([]int{5, 4, -1, 7, 8}))
	fmt.Println(maxSubArray([]int{-3, -2, 0, -1}))

}

func maxSubArray(nums []int) int {
	sum := 0
	max := nums[0]
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum > max {
			max = sum
		}
		if sum <= 0 {
			sum = 0
		}

	}

	return max
}
