package main

import (
	"fmt"
)

func main() {
	a := [][]int{
		[]int{1, 2, 3, 5},
		[]int{4, 5, 6, 3},
		[]int{7, 8, 9, 4},
		[]int{7, 8, 9, 6},
	}
	rotate(a)
	for i := range a {
		fmt.Println(a[i])
	}

}

// clock wise
// 上下翻转, 再对称翻转
// unclock wise
// 左右翻转, 再对称翻转
func rotate(matrix [][]int) {
	for i, j := 0, len(matrix)-1; i < len(matrix) && j >= 0 && i < j; {
		matrix[i], matrix[j] = matrix[j], matrix[i]
		i++
		j--
	}
	for i := 0; i < len(matrix); i++ {
		for j := i + 1; j < len(matrix[i]); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

func reverseUpTodown(matrix [][]int) {
	lens := len(matrix)
	for i := 0; i < lens-i-1; i++ {
		for k := 0; k < lens; k++ {
			fmt.Println(i, k)
			matrix[i][k], matrix[lens-i-1][k] = swap(matrix[i][k], matrix[lens-i-1][k])
		}
	}

}

func reverseLToR(matrix [][]int) {
	lens := len(matrix)
	for i := 0; i < lens; i++ {
		for j := 0; j < lens-j-1; j++ {
			matrix[i][j], matrix[i][lens-j-1] = swap(matrix[i][j], matrix[i][lens-j-1])
		}
	}
}

func swapSymmetry(matrix [][]int) {
	lens := len(matrix)
	for i := 0; i < lens; i++ {
		for k := i + 1; k < lens; k++ {
			matrix[i][k], matrix[k][i] = swap(matrix[i][k], matrix[k][i])
		}
	}
}

func swap(a, b int) (int, int) {
	a, b = b, a
	return a, b
}
