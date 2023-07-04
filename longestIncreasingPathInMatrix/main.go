package main

import "fmt"

func main() {
	fmt.Println(longestIncreasingPath([][]int{
		{3, 4, 5},
		{3, 2, 6},
		{2, 2, 1},
	}))
}

func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 {
		return 0
	}
	maxVal := 1
	cache := make([][]int, len(matrix))
	for i := range cache {
		cache[i] = make([]int, len(matrix[0]))
	}
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			maxVal = max(maxVal, findLongest(matrix, row, col, -1, cache))
		}
	}
	return maxVal
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func findLongest(matrix [][]int, row, col, pre int, cache [][]int) int {
	maxVal := 1
	if row < 0 || row >= len(matrix) || col < 0 || col >= len(matrix[0]) || matrix[row][col] <= pre {
		return 0
	}
	if cache[row][col] > 0 {
		return cache[row][col]
	}
	right := findLongest(matrix, row, col+1, matrix[row][col], cache)
	left := findLongest(matrix, row, col-1, matrix[row][col], cache)
	bottom := findLongest(matrix, row+1, col, matrix[row][col], cache)
	top := findLongest(matrix, row-1, col, matrix[row][col], cache)
	maxVal = max(max(max(left, right), max(top, bottom))+1, maxVal)
	cache[row][col] = maxVal
	return maxVal

}
