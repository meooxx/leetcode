package main

func isPowerOfThree(n int) bool {
	// 3 9 27 81 243
	// 11
	// 1001
	if n < 1 {
		return false
	}

	for n%3 == 0 {
		n /= 3
	}
	return n == 1
}
