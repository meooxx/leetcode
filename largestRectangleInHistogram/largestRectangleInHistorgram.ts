function largestRectangleArea(heights: number[]): number {
	const leftLessHeights: number[] = [-1]
	const rightLessHeights: number[] = Array(heights.length)
	rightLessHeights[heights.length - 1] = heights.length
	// create array in which, each index recored the first element < array[index]
	//  3     7                6
	// -1     0(index of 3)    0
	// 找比6小的, 先看前一位, 如果前一位 > 6, 直接找比7小的
	// 然后跟6比较, 相当于缓存了
	for (let index = 1; index < heights.length; index++) {
		let i = index - 1
		while (i >= 0 && heights[index] <= heights[i]) {
			i = leftLessHeights[i]
		}
		leftLessHeights[index] = i
	}
	// 7  8         6     
	//    current   len(array)
	// same as previous
	// look forward to find the first element which is less than current
	// element
	for (let index = heights.length - 2; index >= 0; index--) {
		let i = index + 1
		while (i < heights.length && heights[index] <= heights[i]) {
			i = rightLessHeights[i]
		}
		rightLessHeights[index] = i
	}
	let maxArea = -1
	for (let i = 0; i < heights.length; i++) {
		maxArea = Math.max(maxArea, (rightLessHeights[i] - leftLessHeights[i] - 1) * heights[i])

	}
	return maxArea

};