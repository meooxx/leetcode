package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{
		3, 2, 5, 8, 1, 9,
	}))
	fmt.Println(maxProfit1([]int{
		3, 2, 5, 8, 1, 9,
	}))
}
func maxProfit1(prices []int) int {
	hold1 := -100000
	hold2 := hold1
	profit1 := 0
	profit2 := 0
	max := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}
	for i := range prices {
		hold1 = max(hold1, -prices[i])
		profit1 = max(profit1, prices[i]+hold1)
		hold2 = max(hold2, -prices[i]+profit1)
		profit2 = max(profit2, prices[i]+hold2)

	}
	return profit2
}
func maxProfit(prices []int) int {
	dp := make([][]int, 3)
	dp[0] = make([]int, len(prices))
	max := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}
	// 每一个节点使用最小价格做卖出运算, 利用变量maxProfit或者minPirce保存
	// 用dp[:]保存利润或者最低价格
	
	// 第二次同理, 同时利用第一次利润或者最低价格
	// ==>
	//              i-1   i     dp[k-1][i] == dp[k-1][i] 原因是i是降低点,
	//                          要获取最大利润一点是在i的前一个时间, 且i不能卖出继承前一天利润
 	//  3,  2,  5,   8,  1, 9,
	// -3  -2  -2   -2  -1 -1    maxProfit  loop1
	//  0   0   3   6   6   8     dp
	// -3  -2  -2  -2   5   5    maxProfit  loop2
	//  0   0   3   6   5  14
	for k := 1; k <= 2; k++ {
		maxProfit := -prices[0]
		dp[k] = make([]int, len(prices))
		for i := 1; i < len(prices); i++ {
			maxProfit = max(maxProfit, -prices[i]+dp[k-1][i])
			dp[k][i] = max(dp[k][i-1], prices[i]+maxProfit)
		}
	}
	return dp[2][len(prices)-1]
}
