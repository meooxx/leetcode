package main

func rangeBitwiseAnd(left int, right int) int {
	step := 0
	for left != right {
		step++
		left >>= 1
		right >>= 1
	}
	return left << step
}
