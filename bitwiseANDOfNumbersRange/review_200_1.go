package main
// 1101
// 1102
// 1110
// 1111
// for left != right
//   then left >> 1, right >> 1
// 11 -> 1100
func rangeBitwiseAnd(left int, right int) int {
	if left == 0 {
		return 0
	}
	step := 0
	for left != right {
		step++
		left >>= 1
		right >>= 1
	}
	return left << step

}
