package main

import "math"

func maxSumSubmatrix(matrix [][]int, k int) int {
	rows := len(matrix)
	cols := len(matrix[0])
	maxVal := math.MinInt

	for col := 0; col < cols; col++ {
		sum := make([]int, rows)
		for j := col; j < cols; j++ {
			for row := 0; row < rows; row++ {
				sum[row] += matrix[row][j]
			}
			for row := 0; row < len(sum); row++ {
                tangleSum := 0
                for l:=row;l<len(sum);l++ {
                    tangleSum += sum[l]
                    if tangleSum <= k && tangleSum > maxVal {
                        maxVal = tangleSum
                    }
                }
				

			}
		}

	}
	return maxVal
}
