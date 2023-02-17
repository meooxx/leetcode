package main

// CUZ' 10 * x => xx0
//     == 2 * 5 * x
//     2 is much more than 5
// so count num of 5
// 110
// 110 / 5  = 22
//    +
// 110 / 25 = 4
//       ...
// is equal with
//   110 /5 + 22 /5 ...
//  = loop n / 5
func trailingZeroes(n int) int {
	result := 0
	for n >= 5 {
		result += n / 5
		n = n / 5
	}
	return result
}
