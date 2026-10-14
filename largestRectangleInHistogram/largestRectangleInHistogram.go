package main

func largestRectangleArea1(heights []int) int {

	stack := []int{}
	heights = append(heights, 0)
	maxArea := -1
	leftIndex := 0
	for index, v := range heights {

		if index > 0 && v < heights[index-1] {
			top := 0
			for len(stack) > 0 && v <= heights[stack[len(stack)-1]] {
				top = heights[stack[len(stack)-1]]

				leftIndex = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				// 5 8 9, 4
				if len(stack) == 0 {
					leftIndex = 0
				}
				maxArea = max(maxArea, (index-leftIndex-1)*top)
			}

		}
		stack = append(stack, index)

	}
	return maxArea
}


func largestRectangleArea(heights []int) int {
	leftLess := make([]int, len(heights))
	leftLess[0] = -1
	rightLess := make([]int, len(heights))
	rightLess[len(heights)-1] = len(heights)
	maxArea := 0
	//  4  5  3 2
	for index := 1; index < len(heights); index++ {
		p := index - 1

		for p >= 0 && heights[index] <= heights[p] {
			p = leftLess[p]
		}
		leftLess[index] = p
	}

	for index := len(heights) - 2; index >= 0; index-- {
		p := index + 1
		for p < len(heights) && heights[index] <= heights[p] {
			p = rightLess[p]
		}
		rightLess[index] = p

	}
	//  2  5 6 4
	for index, v := range heights {
		maxArea = max(maxArea, v*(rightLess[index]-leftLess[index]-1))
	}
	return maxArea

}
