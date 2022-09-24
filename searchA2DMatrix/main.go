package main

import "fmt"

func main() {
	fmt.Println(searchMatrix([][]int{
		{-8, -7, -5, -3, -3, -1, 1}, {2, 2, 2, 3, 3, 5, 7},
	}, -5))
}

func searchMatrix(matrix [][]int, target int) bool {
	start := 0
	end := len(matrix) - 1
	// 1 4 10
	for start <= end {
		mid := (end-start)/2 + start

		if matrix[mid][0] == target {
			return true
		} else if matrix[mid][0] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	start -= 1
	if start < 0 || start == len(matrix) {
		return false
	}
	rowData := matrix[start]
	start = 0
	end = len(rowData) - 1
	// len() == 1
	for start <= end {
		mid := (end-start)/2 + start
		if rowData[mid] == target {
			return true
		} else if rowData[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return false
}
