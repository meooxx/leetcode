package main

func maximalRectangle(matrix [][]byte) int {
	heights := make([]int, len(matrix[0]))
	left := make([]int, len(matrix[0]))
	right := make([]int, len(matrix[0]))
	// at the farest rigth by default
	for i := range right {
		right[i] = len(matrix[0])
	}
	maxArea := 0
	for _, rowVal := range matrix {
		leftCol := 0
		for col := range rowVal {
			if rowVal[col] == '1' {

				// 0  1   1
				// 0  0  [1]  
				// alway use the greater left index
				// to narrow the actual area size
				left[col] = max(left[col], leftCol)
			} else {
				left[col] = 0
				leftCol = col + 1
			}
			// accumlative heights
			//  		0 1 1
			//      0 1 1
			// ==>  0 2 2
			if rowVal[col] == '1' {
				heights[col] += 1
			} else {
				heights[col] = 0
			}
		}
		rightCol := len(rowVal)
		for col := len(rowVal) - 1; col >= 0; col-- {
			if rowVal[col] == '1' {

				//   0  1  0
				//   0  1  1
				// use smaller right index
				right[col] = min(right[col], rightCol)
			} else {

				right[col] = len(rowVal)
				rightCol = col
			}

		}

		for col := range rowVal {
			maxArea = max(maxArea, heights[col]*(right[col]-left[col]))
		}
	}
	return maxArea
}
func main() {
	maximalRectangle([][]byte{[]byte{'1'}})
}
