package main

import "math"

func maxSumSubmatrix(matrix [][]int, k int) int {
	rows := len(matrix)
	cols := len(matrix[0])
	maxSum := math.MinInt

	for i := 0; i < cols; i++ {
		sum := make([]int, rows)
		for j := i; j < cols; j++ {
			for k := 0; k < rows; k++ {
				sum[k] += matrix[k][j]
			}
			for l := 0; l < len(sum); l++ {
				sumArr := 0
				for r := l; r < len(sum); r++ {
					sumArr += sum[r]
					if sumArr <= k && sumArr > maxSum {
						maxSum = sumArr
					}
				}
			}
		}
	}

	return maxSum

}
//                     sumArr   
//  l       r           r =0     r = 1   r =....
//  5  -4  -3   4       5          1 (5 + -4)
// -3  -4   4   5      -3         -7 (-3 + -4 )
//  5   1   5  -4       5          6

// and 
// 5 + -3 == rectangle(5, -3)
// 1 + -7 == rectangle(5, -4, -3, -4)

// maxSum no larger than k = sum[i,j] == sum[j] - sum[i]

