package main

import "fmt"

func main() {
	fmt.Println(maximalRectangle2([][]byte{
		{
			'1', '0', '0', '1',
		},
		{
			'1', '1', '1', '1',
		},
		{
			'0', '1', '1', '1',
		},
	}))
}

func maximalRectangle(matrix [][]byte) int {
	heights := make([]int, len(matrix[0]))
	max := 0
	calcArea := func(h []int) int {
		if len(h) == 1 {
			return h[0]
		}
		m := 0
		stack := []int{0}

		for index := 1; index <= len(h); index++ {
			cur := 0
			if index < len(h) {
				cur = h[index]
			}
			if cur < h[index-1] {
				for len(stack) > 0 && h[stack[len(stack)-1]] >= cur {
					top := h[stack[len(stack)-1]]
					stack = stack[:len(stack)-1]
					width := 0
					if len(stack) == 0 {
						width = index
					} else {
						width = index - stack[len(stack)-1] - 1
					}
					area := top * width
					if area > m {
						m = area
					}
				}
			}
			stack = append(stack, index)
		}
		return m
	}
	for row := range matrix {
		for col := range matrix[row] {
			if matrix[row][col] == '1' {
				heights[col] += 1
			} else {
				heights[col] = 0
			}

		}
		v := calcArea(heights)
		if v > max {
			max = v
		}
	}
	return max
}

func maximalRectangle2(matrix [][]byte) int {

	heights := make([]int, len(matrix[0]))
	lefts := make([]int, len(matrix[0]))
	rights := make([]int, len(matrix[0]))
	for index := range rights {
		rights[index] = len(matrix[0])
	}
	max := 0

	maxInt := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	minInt := func(a, b int) int {
		if a < b {
			return a
		} else {
			return b
		}
	}
	for row := range matrix {
		for col := range matrix[row] {
			if matrix[row][col] == '1' {
				heights[col] += 1
			} else {
				heights[col] = 0
			}
		}
		leftIndex := 0
		for col := range matrix[row] {

			if matrix[row][col] == '1' {
				lefts[col] = maxInt(lefts[col], leftIndex)
			} else {
				lefts[col] = 0
				leftIndex = col + 1
			}
		}
		rightIndex := len(matrix[row])
		for col := len(matrix[row]) - 1; col >= 0; col-- {

			if matrix[row][col] == '1' {
				rights[col] = minInt(rights[col], rightIndex)
			} else {
				rights[col] = len(matrix[row])
				rightIndex = col
			}
		}
		// 0 1 2
		// 1 2 3
		for index := range heights {
			max = maxInt((rights[index]-lefts[index])*heights[index], max)
		}
	}

	return max
}
