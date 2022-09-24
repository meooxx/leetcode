package main

import "fmt"

func main() {
	a := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroes(a)
	fmt.Println(a)
	b := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	setZeroes(b)
	fmt.Println(b)
}

// 第一行,第一列存该行该列是否有0
// tip:从第二列开始 因为如果第一排三个任意有个0, 则 (0,0)为0
// 此时区分不了 第一列是否应该是0
//  1   1   0
//  1   1   1
//  1   1   1
//  1   1   1
func setZeroes(matrix [][]int) {
	col0 := false
	for row := 0; row < len(matrix); row++ {
		if matrix[row][0] == 0 {
			col0 = true
		}
		for col := 1; col < len(matrix[row]); col++ {
			if matrix[row][col] == 0 {
				matrix[row][0] = 0
				matrix[0][col] = 0
			}
		}
	}
	for row := len(matrix) - 1; row >= 0; row-- {
		for col := len(matrix[row]) - 1; col >= 1; col-- {
			if matrix[row][0] == 0 || matrix[0][col] == 0 {
				matrix[row][col] = 0
			}
		}
		if col0 {
			matrix[row][0] = 0
		}
	}
}

func setZeroes2(matrix [][]int) {
	p := [][]int{}
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			if matrix[row][col] == 0 {
				p = append(p, []int{row, col})
			}
		}
	}
	for _, position := range p {
		row := position[0]
		col := position[1]
		for i := 0; i < len(matrix[row]); i++ {
			matrix[row][i] = 0
		}
		for i := 0; i < len(matrix); i++ {
			matrix[i][col] = 0
		}
	}

}
