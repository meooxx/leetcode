package main

func canWinNim(n int) bool {
	// 123    5 6 7    9 10 11
	//      4       8
	if n < 4 {
		return true
	}
	return n%4 != 0

}
