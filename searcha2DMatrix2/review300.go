package main

func searchMatrix(matrix [][]int, target int) bool {
	row := 0
	col := len(matrix[row]) - 1
	for row < len(matrix) && col >= 0 {
		for col >= 0 {
			if matrix[row][col] == target {
				return true
			} else if matrix[row][col] < target {
				row++
				break
			} else {
				col--
			}

		}

	}
	return false
}
