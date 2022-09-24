package main

import "fmt"

func main() {
	fmt.Println(uniquePathsWithObstacles([][]int{
		{0, 0, 0}, {0, 0, 0},
	}))
	fmt.Println(uniquePathsWithObstacles([][]int{
		{0, 1}, {0, 0},
	}))
	fmt.Println(uniquePathsWithObstacles([][]int{
		{1, 0},
	}))
	fmt.Println(uniquePathsWithObstacles([][]int{
		{0},
		{1},
	}))
	fmt.Println(uniquePathsWithObstacles([][]int{
		{1},
		{0},
		{0},
	}))
	fmt.Println(uniquePathsWithObstacles([][]int{
		{1},
		{1},
	}))

}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {

	r := make([]int, len(obstacleGrid[0]))
	// if there is a square at first row, break at the square position
	// 0   1   0     square
	// x   x   x
	// 1   0   0
	for i := range r {
		if obstacleGrid[0][i] == 0 {
			r[i] = 1
		} else {

			break
		}
	}
	for row := 1; row < len(obstacleGrid); row++ {

		// 1^1 0, 0^1 1
		// == r[0] == 0 || obstacleGrid[row][0] == 1
		r[0] = (obstacleGrid[row][0] ^ 1) & r[0]
		for col := 1; col < len(obstacleGrid[row]); col++ {
			if obstacleGrid[row][col] == 0 {
				r[col] = r[col] + r[col-1]
			} else {
				r[col] = 0
			}
		}

	}
	return r[len(r)-1]
}
