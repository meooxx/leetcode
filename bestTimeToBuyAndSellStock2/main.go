package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{
		3, 2, 5, 8, 1, 9,
	}))
}

// 3 1 2 8 1 9
func maxProfit(prices []int) int {
	i := 0
	end := len(prices) - 1
	max := 0
	for i < end {
		for i < end && prices[i+1] <= prices[i] {
			i++
		}
		buy := prices[i]
		// >= 和 > 一样
		// 1 没有 = 会多一次循环
		// 如  1 8  8 10
		// 1,8 + 8,10
		// 2 有 = 则 1,10
		for i < end && prices[i+1] > prices[i] {
			i++
		}
		sell := prices[i]
		max += sell - buy
	}
	return max
}
