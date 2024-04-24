package main

import "math"

func maxProfit(k int, prices []int) int {
	dpHold := make([]int, len(prices)+1)
	dpProfit := make([]int, len(prices)+1)
	dpHold[0] = math.MinInt32
	dpProfit[0] = 0
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for k > 0 {
		for i := range prices {
			// dpProfit[i+1], 只能使用k-1轮的盈利
			// 不能使用dpProfit[i], 前面交易盈利, 单次交易
			dpHold[i+1] = max(dpHold[i], dpProfit[i+1]-prices[i])
			dpProfit[i+1] = max(dpProfit[i], prices[i]+dpHold[i])
		}

		k--
	}
	return dpProfit[len(prices)]
}
