package main

func isPowerOfFour(n int) bool {
	for i := 0; i < 32; i += 2 {
		if 1<<i == n {
			return true
		}
	}
	return false
}
