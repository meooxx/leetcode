package main

func searchMatrix(matrix [][]int, target int) bool {

	for i := 0; i < len(matrix[0]); i++ {
		if target >= matrix[0][i] {
			top := 0
			bottom := len(matrix) - 1
			for top <= bottom {
				mid := (bottom-top)/2 + top
				midVal := matrix[mid][i]
				if target == midVal {
					return true
				} else if target > midVal {
					top = mid + 1
				} else {
					bottom = mid - 1
				}
			}
		}
	}
	return false
}
