package main

import (
	"fmt"
)

// 1, 5,2,3,7, 3
// i           j   1
//   i         j   15
//   i      j      35
// left, right 俩个指针, 移动较小值的指针
// 保存每次移动俩个乘积
func maxArea(height []int) int {
	i := 0
	j := len(height) - 1
	max := 0
	sum := 0
	for i < j {
		if height[i] < height[j] {
			sum = height[i] * (j - i)
			if sum > max {
				max = sum
			}
			i++

		} else {
			sum = height[j] * (j - i)
			if sum > max {
				max = sum
			}
			j--

		}
	}
	return max
}

func main() {
	a := [4]int{3, 71, 71, 3}
	fmt.Println(maxArea(a[:]))
}
