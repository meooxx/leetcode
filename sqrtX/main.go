package main

import "fmt"

func main() {
	fmt.Println(mySqrt(9))
	fmt.Println(mySqrt(8))
	fmt.Println(mySqrt(7))
	fmt.Println(mySqrt(6))
}

func mySqrt(x int) int {
	if x == 0 {
		return x
	}
	left := 1
	right := x
	for left <= right {
		mid := (-left+right)/2 + left
		if mid <= x/mid {
			left = mid + 1
		} else {
			right = mid - 1
		}

	}
	return right
}
