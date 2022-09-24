package main

import (
	"fmt"
)

func main() {
	// fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
	fmt.Println(largestRectangleArea2([]int{4, 2, 0, 3, 2, 5}))

}

// 按位推入栈中
// [2]
//下降趋势: 1  1< [2] // 最大值可能存在 下降前, 或者包含下降值
//  如 561, 最大5*2 或者 564, 最大4*3
// []  max=2 * 1
// [1,5,6]
// 2 2<6, max = 6 *1 => 5*2
// [1,4,5]
// 0 =>  (3*1, 2*3, 1 * 1, )
func largestRectangleArea(heights []int) int {
	stack := []int{}
	max := 0
	for index := 0; index <= len(heights); index++ {
		n := 0
		if index < len(heights) {
			n = heights[index]
		}

		if index > 0 && n < heights[index-1] {

			for len(stack) > 0 && heights[stack[len(stack)-1]] >= n {
				top := heights[stack[len(stack)-1]]
				stack = stack[:len(stack)-1]
				width := 0
				// case: 2 1 2 stack [1,2]
				if len(stack) == 0 {
					width = index // index - (-1) -1
				} else {
					width = index - stack[len(stack)-1] - 1
				}
				area := top * width

				if area > max {
					max = area
				}
			}
		}
		stack = append(stack, index)
	}
	return max
}

func largestRectangleArea2(heights []int) int {
	stack := []int{}
	max := 0

	for index := 0; index <= len(heights); index++ {
		cur := 0
		if index < len(heights) {
			cur = heights[index]
		}
		if index > 0 && cur < heights[index-1] {
			for len(stack) > 0 && heights[stack[len(stack)-1]] >= cur {
				heightIndex := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				width := 0
				
				if len(stack) == 0 {
					width = index
				} else {
					width = index - stack[len(stack)-1] - 1
				}
				m := heights[heightIndex] * width
				if m > max {
					max = m
				}
			}
		}

		stack = append(stack, index)

	}

	return max
}
