package main

func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}
	// 1
	// 10
	// 100
	// 1000
	// 10000
	// 5 1001
	// 4 1000  5 & 4 == 1
	// and  pow of two n only have a '1'
	return n&(n-1) == 0

}
