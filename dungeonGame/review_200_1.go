package main

import "math"

/*
*

	   2    -3     3    inf
		 -5   -10    1    inf
	   10   30    -5    1
		 inf  inf    1

		1 last line inital with 1, 1, make sure at least 1 health
		2 line which is not last line, gei it to inf
*/
func calculateMinimumHP(dungeon [][]int) int {
	dp := make([]int, len(dungeon[0])+1)

	for i := range dp {
		dp[i] = math.MaxInt
	}
	// at -5 position
	dp[len(dungeon[0])] = 1
	dp[len(dungeon[0])-1] = 1

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for i := len(dungeon) - 1; i >= 0; i-- {
		if i != len(dungeon)-1 {
			dp[len(dungeon[0])] = math.MaxInt
		}
		for j := len(dungeon[i]) - 1; j >= 0; j-- {
			need := min(dp[j], dp[j+1]) - dungeon[i][j]
			if need <= 0 {
				dp[j] = 1
			} else {
				dp[j] = need
			}
		}

	}
	return dp[0]
}
