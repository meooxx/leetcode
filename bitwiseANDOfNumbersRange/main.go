package main

func rangeBitwiseAnd1(left int, right int) int {
	anwser := 0
	start := 0
	count := right - left + 1
	if right == left {
		return left
	}
	for start < 31 {
		maxCount := 1 << start
		maxVal := (maxCount << 1) - 1
		if (left&maxVal) >= maxCount && (right&maxVal) <= maxVal && (right&maxVal) >= maxCount && count <= maxCount {
			anwser += maxCount
		}
		start++
	}

	return anwser
}
//      
//  11010
// 	11011
//	11100　　
//	11101　　
//	11110
//      ^  left != right   step++
//     ^   left != right   step++
//    ^    left != right   step++
//  11<<step => 11 000
func rangeBitwiseAnd(left int, right int) int {
	if left == 0 {
		return 0
	}
	base := 0
	for left != right {
		left >>= 1
		right >>= 1
		base++
	}
	return left << base
}

//  100  4
//  101. 5
//  110. 6
//  111. 7
// 1000. 8
// 1001  9
// 1010
// 1011
// 1100
// 1101
// 1110 14
// 1111 15
