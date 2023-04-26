package main

// dp[i] = min(dp[i-square] + dp[square], dp[i])

// 1 * 1
// dp[0] 1  2  3 4 5...8 9 ... 12 .....dp[n]
//  0    1  2  3 4 5   8 9     12         n

// 2 * 2                       12 dp([12-4] + 1, dp[12])
//
//	...     1 2 3 2 3      3
//
// 3*3
//
//	1     min(9_1_1_1, 4_4_4)
func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = i
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for s := 2; s*s <= n; s++ {
		sq := s * s
		for i := sq; i < n+1; i++ {
			dp[i] = min(dp[i-sq]+1, dp[i])
		}
	}
	return dp[n]
}

// 1 2    3     4      5    16    7   8    9  10   11    12  13 14
// 1 11  111    4     41   411  4111 44    9  91   911  9111 94 941
// 1 2   3      1
// 15    16    17     18     19   20    21    22    23     24  25 26
// 9411  16   161     1611  16111 164   1641 16411 164111 1644 25 251
