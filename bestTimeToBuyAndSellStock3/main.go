package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxProfit([]int{3, 3, 5, 0, 3, 1, 4}))
	fmt.Println(maxProfit([]int{1, 2, 3, 4, 5}))
	fmt.Println(maxProfit2([]int{3, 3, 5, 0, 3, 1, 4}))
	fmt.Println(maxProfit2([]int{1, 2, 3, 4, 5}))
	fmt.Println(maxProfit3([]int{3, 3, 5, 0, 3, 1, 4}))
	fmt.Println(maxProfit3([]int{1, 2, 3, 4, 5}))
}

func maxProfit(prices []int) int {
	// min := func(a, b int) int {
	// 	if a <= b {
	// 		return a
	// 	}
	// 	return b
	// }
	max := func(a, b int) int {
		if a <= b {
			return b
		}
		return a
	}
	hold1 := math.MinInt32
	hold2 := math.MinInt32
	profit1 := 0
	profit2 := 0
	for i := 0; i < len(prices); i++ {
		hold1 = max(hold1, -prices[i])
		profit1 = max(profit1, prices[i]+hold1)
		hold2 = max(hold2, -prices[i]+profit1)
		profit2 = max(profit2, prices[i]+hold2)
	}
	return profit2
}

func maxProfit2(prices []int) int {
	// 3 means K tranctions
	dp := make([][]int, 3)
	for i := range dp {
		dp[i] = make([]int, len(prices))
	}

	// min := func(a, b int) int {
	// 	if a <= b {
	// 		return a
	// 	}
	// 	return b
	// }
	max := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}

	//    3 5 0 3 1 4      3 5 0 3 1 4     3 5 0 3 1 4
	// 0  0 000000000      0 0 0 0 0 0     0 0 0 0 0 0
	// 1  0                0 2 2 3 3 4     0 2 2 3 3 4
	// 2  0                0 2 2 4         0 2 2(no tranction)
	//
	//
	for k := 1; k <= 2; k++ {
		maxProfit := -prices[0]
		for i := 1; i < len(prices); i++ {
			// i-1 与 i 问题(如果限制不能同时卖出和买入,应该不能使用i)
			// 利润最大情况下
			// 因为第二次买入在i点时, 那么i点就是个低点, i点卖出的一定不如不交易的利润大
			// 即dp[k][i]同dp[k][i-1] (i此时必须是下次买入的低点)
			// 所以不管取i-1 和 i 最后结果相同
			maxProfit = max(maxProfit, dp[k-1][i]-prices[i])
			dp[k][i] = max(dp[k][i-1], prices[i]+maxProfit)
		}
	}
	return dp[2][len(prices)-1]
}

func maxProfit3(prices []int) int {
	buy1 := math.MinInt32
	sell1 := 0
	buy2 := math.MinInt32
	sell2 := 0
	max := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}
	for i := range prices {
		// profit
		buy1 = max(buy1, -prices[i])
		sell1 = max(sell1, prices[i]+buy1)
		buy2 = max(buy2, sell1-prices[i])
		sell2 = max(sell2, prices[i]+buy2)
	}
	return sell2
}
