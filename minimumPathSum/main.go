package main

import "fmt"

func main() {
	fmt.Println(minPathSum([][]int{
		{1, 3, 1}, {1, 5, 1}, {4, 2, 1},
	}))
	fmt.Println(minPathSum([][]int{
		{1, 2, 3}, {4, 5, 6},
	}))
}

// 第 x,y 个位置可以从 x-1, y 或者 x, y-1处得出
// 改变原来数组
func minPathSum(grid [][]int) int {

	firstRow := grid[0]
	m := len(grid)
	n := len(firstRow)
	for i := 1; i < n; i++ {
		firstRow[i] += firstRow[i-1]
	}
	for row := 1; row < m; row++ {
		grid[row][0] += grid[row-1][0]
		for col := 1; col < n; col++ {
			if grid[row-1][col] > grid[row][col-1] {
				grid[row][col] += grid[row][col-1]
			} else {
				grid[row][col] += grid[row-1][col]
			}
		}
	}
	return grid[m-1][n-1]
}

// 新建数组存储结果
func minPathSum2(grid [][]int) int {

	firstRow := grid[0]
	r := make([]int, len(firstRow))
	for i := range r {
		if i == 0 {
			r[i] = firstRow[i]
		} else {
			r[i] = firstRow[i] + r[i-1]
		}
	}
	for row := 1; row < len(grid); row++ {
		r[0] = r[0] + grid[row][0]
		for col := 1; col < len(firstRow); col++ {
			if r[col] > r[col-1] {
				r[col] = r[col-1] + grid[row][col]
			} else {
				r[col] = r[col] + grid[row][col]
			}
		}
	}
	return r[len(r)-1]
}
