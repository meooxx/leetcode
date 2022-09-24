package main

import "fmt"

func main() {
	a := [][]int{
		{1, 2, 3, 4, 5},
		{14, 15, 16, 17, 6},
		{13, 20, 19, 18, 7},
		{12, 11, 10, 9, 8},
	}
	fmt.Println(spiralOrder(a))
	b := [][]int{
		{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12},
	}
	fmt.Println(spiralOrder(b))

	c := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	fmt.Println(spiralOrder(c))
}

// 按四周顺序循环, 注意右往左, 下往上时重复
// 同时要求right > left, bottom > top
// 因为如果不成立说明不换行, 会重复
// 1 2 3 5 x
// 8 5 4 6
// 7 6 5 7
//     y
func spiralOrder(matrix [][]int) []int {
	answer := []int{}

	top := 0
	bottom := len(matrix)
	left := 0
	right := len(matrix[0])

	for left < right && top < bottom {
		row := top
		col := left

		for col < right {
			answer = append(answer, matrix[row][col])
			col++
		}
		top++
		col--
		row++
		for row < bottom {
			answer = append(answer, matrix[row][col])
			row++
		}
		right--
		row--
		col--

		for top < bottom && col >= left {
			answer = append(answer, matrix[row][col])
			col--

		}
		bottom--
		col++
		row--
		for right > left && row >= top {
			answer = append(answer, matrix[row][col])
			row--
		}
		left++
	}
	return answer
}
