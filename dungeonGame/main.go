package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println(calculateMinimumHP([][]int{
	// 	{-2, -3, 3},
	// 	{-5, -10, 1},
	// 	{10, 30, -5},
	// }))
	// fmt.Println(calculateMinimumHP([][]int{
	// 	{100},
	// }))
	fmt.Println(calculateMinimumHP([][]int{
		{1, -3, 3},
		{0, -2, 0},
		{-3, -3, -3},
	}))
}

func calculateMinimumHP1(dungeon [][]int) int {
	rowLen := len(dungeon)
	colLen := len(dungeon[0])
	dp := make([][]int, rowLen+1)
	for i := range dp {
		dp[i] = make([]int, colLen+1)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt
		}
	}
	dp[rowLen][colLen-1] = 1
	dp[rowLen-1][colLen] = 1
	min := func(a, b int) int {
		if a <= b {
			return a
		}
		return b
	}
	for i := rowLen - 1; i >= 0; i-- {
		for j := colLen - 1; j >= 0; j-- {
			need := min(dp[i][j+1], dp[i+1][j]) - dungeon[i][j]
			if need <= 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = need
			}
		}
	}
	return dp[0][0]

}

func calculateMinimumHP(dungeon [][]int) int {
	rowLen := len(dungeon)
	colLen := len(dungeon[0])
	dp := make([]int, colLen+1)
	for i := range dp {
		dp[i] = math.MaxInt
	}

	dp[colLen] = 1
	dp[colLen-1] = 1
	min := func(a, b int) int {
		if a <= b {
			return a
		}
		return b
	}
	for i := rowLen - 1; i >= 0; i-- {
		if i != rowLen-1 {
			dp[colLen] = math.MaxInt
		}
		for j := colLen - 1; j >= 0; j-- {
			minHelth := min(dp[j], dp[j+1]) - dungeon[i][j]
			if minHelth <= 0 {
				dp[j] = 1
			} else {
				dp[j] = minHelth
			}
		}
	}
	return dp[0]
}
