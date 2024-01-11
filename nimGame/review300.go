package main

// (4n + k)  1 <= k <= 3
// 1 2 3    5 6  7     9
//
//				4                 8
func canWinNim(n int) bool {
	return n%4 != 0
}
