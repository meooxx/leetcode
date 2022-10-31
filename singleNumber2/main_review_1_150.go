package main

import "fmt"

func main() {
	fmt.Println(singleNumber1([]int{
		2, 2, 2, 1,
	}))
	fmt.Println(singleNumber1([]int{
		2, 2, 2, 8, 3, 3, 3,
	}))
	fmt.Println(singleNumber1([]int{
		-2, -2, -2, -1, 3, 3, 3,
	}))
	fmt.Println(1 << 31)
}

func singleNumber(nums []int) int {
	x0 := 0
	x1 := 0
	mask := 0
	//   001    x1  x0                       '1' 
	//   001        0 -> 1                    1 
	//   001    0   1 -> 0    0^(1&1)==1      1       
	//   010    1   0 -> 1    1^(0&1)==1 
	//   010    0   1 -> 0    0^(1&1)==0 
	//   010     
	
	//   101
	for i := range nums {
		x1 ^= nums[i] & x0
		x0 ^= nums[i]
		mask = ^(x0 & x1)
		x0 &= mask
		x1 &= mask
	}
	return x0
}

func singleNumber1(nums []int) int {
	result := 0

	for bit := 0; bit < 32; bit++ {
		count := 0
		// 101
		for i := range nums {
			if (nums[i]>>bit)&1 == 1 {
				count++
			}

		}
		if count%3 == 1 {
			result += 1 << bit
		}
	}
	if result >= 1<<31 {
		result = result - 1<<32
	}
	return result
}
