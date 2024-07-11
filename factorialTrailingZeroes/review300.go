package main

//
func trailingZeroes(n int) int {
	total := 0
	// at the fist loop, it means n / 5
	// at the second loop, it means n / 25
	// so forth
	for n >= 5 {
		total += n / 5
		n /= 5

	}
	return total

}
