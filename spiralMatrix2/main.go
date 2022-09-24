package main

import "fmt"

func main() {
	fmt.Println(generateMatrix(4))
}

func generateMatrix(n int) [][]int {
	result := make([][]int, n)
	for index := range result {
		result[index] = make([]int, n)
	}

	left := 0
	right := n - 1
	top := 0
	bottom := n - 1
	count := 1
	for left <= right && top <= bottom {
		for i := left; i <= right; i++ {
			result[top][i] = count
			count++
		}
		top++
		for i := top; i <= bottom; i++ {
			result[i][right] = count
			count++
		}
		right--
		if top < bottom {
			for i := right; i >= left; i-- {
				result[bottom][i] = count
				count++
			}
			bottom--
		}
		if right > left {
			for i := bottom; i >= top; i-- {
				result[i][left] = count
				count++
			}
			left++
		}
	}
	return result

}
