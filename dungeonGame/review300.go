package main

import (
	"math"
)

//	right boundry
//
// -2  -3   3     m      m -> denotes max
// -5  10   1     m
// 10  30  -5(?)  1
//	m   m   1     m  bottom boundry

func calculateMinimumHP(dungeon [][]int) int {
	dp := make([]int, len(dungeon[0])+1)
	for i := range dp {
		dp[i] = math.MaxInt
	}
	//          ?     1
	//  m   m   1     m    two dimensions
	//  [m, m,  1,    1]   1 dimension
	dp[len(dungeon[0])] = 1
	dp[len(dungeon[0])-1] = 1
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	//               right boundry
	// -2  -3   3     m      m -> denotes max
	// -5  10   1     m
	// 10  30  -5(?)  1
	//  m   m   1     m  bottom boundry
	for row := len(dungeon) - 1; row >= 0; row-- {
		if row != len(dungeon)-1 {
			dp[len(dungeon[0])] = math.MaxInt
		}
		for col := len(dungeon[0]) - 1; col >= 0; col-- {
			health := min(dp[col], dp[col+1]) - dungeon[row][col]
			if health <= 0 {
				dp[col] = 1
			} else {
				dp[col] = health
			}
		}
	}
	return dp[0]

}
