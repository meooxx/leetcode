package main

// n = d0 + d1 * 10^1 + d2 * 10^2 ... + dk * 10^k
// n = d0 + d1 * (9 * 1 + 1) + d2*(9*11+1))...dk*(9*11...11 + 1)
// n = (d0 + d1 + d2...+dk) + 9*(d1*1 + d2 * 11 + d3 * 11..11)
//
//	|
//	V
//
// n mod 9 = (d0 + d1 + d2... + dk) mod 9
// n == 0 => 0
// 1 + (n-1)mod9
func addDigits(num int) int {
	if num == 0 {
		return num
	}
	return 1 + (num-1)%9
}
