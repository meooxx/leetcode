package main

import "fmt"

func main() {
	// fmt.Println(singleNumber([]int{1, 1, 1, 2}))
	// fmt.Println(singleNumber([]int{1, 1, 1, 2, 2, 2, 3}))

	// fmt.Println(singleNumber([]int{1, 2, 2, 2, 6, 1, 1, 3, 3, 3, 5, 5, 5}))
	fmt.Println(singleNumber([]int{2, 2, 2, 6}))

	fmt.Println(singleNumber1([]int{1, 1, 1, 2}))
	fmt.Println(singleNumber1([]int{1, 1, 1, 2, 2, 2, 3}))

	fmt.Println(singleNumber1([]int{1, 2, 2, 2, 6, 1, 1, 3, 3, 3, 5, 5, 5}))
	fmt.Println(singleNumber1([]int{-1, 2, 2, 2}))

}

func singleNumber(nums []int) int {

	// 2^2 = 4 > 3
	//俩位覆盖重复数量
	//x2, x1
	x1, x2 := 0, 0
	for i := range nums {
		// x1 == 0, 加1, x2不变,x2是高位
		// 当 x1是1, 0 + 1 = 1, 1+1 = 0 所以
		// nums[i] & x1 XOR x2
		x2 ^= nums[i] & x1
		// 0+1 == 1, 1+1 = 0所有x1位必定是XOR
		x1 ^= nums[i]
		
		// 0bx11时 则取0, 所以
		mask := ^(x1 & x2)
		x2 &= mask
		x1 &= mask
	}
	// Suppose to single number is 6
	//  2 2 2 6     222  66
	// 010          010
	// 010          010
	// 010          010
	// 110          110, 110
	//
 	//  q3           q2			q1	
	// 1*1%3 == 1  | 4%3 == 1 | 0  == q'
	// 1               1        0  == x1
	// 0               0        1  == p'
	// every 1-bit of x1 is as same as single number's
	// 不是来自singleNumber的1-bit 全被削掉(%3)了
	// evert 0-bit of x1 equals single number's 因为不相同的话满足不了三个重复一个不重复条件
	// 所以当 p'j == 1, xj == single number

	// 公式 xj_r = s_r & p'j_r
	return x1

}

// 32-bit
func singleNumber1(nums []int) int {
	result := 0

	for bit := 0; bit < 32; bit++ {
		bitCount := 0
		for i := range nums {
			bitVal := (nums[i] >> bit) & 1
			bitCount += bitVal
		}
		// 0 || 1
		// bitCount equals x times three +1
		if bitCount%3 > 0 {
			result += 1 << bit
		}
	}
	if result > 1<<31 {
		return result - 1<<32
	}
	return result
}
