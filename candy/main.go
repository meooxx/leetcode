package main

import "fmt"

func main() {
	fmt.Println(candy([]int{
		1, 0, 2,
	}))
	fmt.Println(candy([]int{
		1, 2, 2,
	}))
	fmt.Println(candy([]int{
		1, 0, 2, 3, 4, 1,
	}))
	fmt.Println(candy([]int{
		2, 3, 1, 0, 4, 1,
	}))
}

func candy(ratings []int) int {
	candys := make([]int, len(ratings))
	for i := range candys {
		candys[i] = 1
	}
	max := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			candys[i] = candys[i-1] + 1
		} 
	}
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candys[i] = max(candys[i+1]+1, candys[i])
		}
	}
	anwser := 0

	for i := range candys {
		anwser += candys[i]
	}
	return anwser
}

// passed , but it costs alot time
func candy1(ratings []int) int {
	dp := make([]int, len(ratings))
	for i := range dp {
		dp[i] = 1
	}
	anwser := 0

	// case1: 2 1
	// case2: 1 2
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			dp[i] = dp[i-1] + 1
		} else if ratings[i] < ratings[i-1] {
			for j := i - 1; j >= 0; j-- {
				if ratings[j] > ratings[j+1] && dp[j] <= dp[j+1] {
					dp[j]++
				} else {
					break
				}
			}
		}
	}
	for _, v := range dp {
		anwser += v
	}
	return anwser
}
