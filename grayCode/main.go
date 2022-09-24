package main

import "fmt"

func main() {
	fmt.Println(grayCode(2))
}

// dp
// 2 
// 000---------
// 010-------
// 011-----
// 001--     倒叙前缀+1
// 3
// 101--        
// 111-----
// 110-------
// 100---------
//  1   0 	1 	1    => grayCode
//      1   0   1    == 1011 >>1
//  1高位不变
//     1+1 0+1 1+1 上一位加本位
func grayCode(n int) []int {
	result := []int{}

	for b := 0; b < 1<<n; b++ {
		result = append(result, b^b>>1)
	}
	return result
}
