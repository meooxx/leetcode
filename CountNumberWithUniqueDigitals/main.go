package main

//  0      1
//  1      10
//  2      9 * 9
//  3      9 * 9 * 8

func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	result := 10
	base := 9
	factor := 9
	for i := 2; i <= 10 && i <= n; i++ {
		base *= factor
		factor--
		result += base
	}
	return result
}
