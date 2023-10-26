package main

func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 {
		return 0
	}
	cache := make([][]int, len(matrix))
	for i := range cache {
		cache[i] = make([]int, len(matrix[0]))
	}
	maxV := 1
	for row := range matrix {
		for col := range matrix[row] {
			v := next(matrix, row, col, cache, -1)
			if v > maxV {
				maxV = v
			}
		}
	}
	return maxV
}

func next(matrix [][]int, row, col int, cache [][]int, pre int) int {
	if row < 0 || col < 0 || row >= len(matrix) || col >= len(matrix[row]) {
		return 0
	}
	if matrix[row][col] <= pre {
		return 0
	}
	if cache[row][col] > 0 {
		return cache[row][col]
	}
	left := next(matrix, row, col-1, cache, matrix[row][col])
	right := next(matrix, row, col+1, cache, matrix[row][col])
	top := next(matrix, row+1, col, cache, matrix[row][col])
	bottom := next(matrix, row-1, col, cache, matrix[row][col])
	maxV := 1 + max(max(left, right), max(top, bottom))
	cache[row][col] = maxV
	return maxV
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
