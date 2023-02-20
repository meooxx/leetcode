package main

import "math"

func maxProfit(k int, prices []int) int {

	dpHold := make([]int, len(prices)+1)
	dpProfit := make([]int, len(prices)+1)
	dpHold[0] = math.MinInt
	dpProfit[0] = 0

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for k > 0 {
		for i := range prices {
			dpHold[i+1] = max(dpHold[i], dpProfit[i+1]-prices[i])
			dpProfit[i+1] = max(dpProfit[i], dpHold[i+1]+prices[i])
		}
		k--
	}
	return dpProfit[len(prices)]

}
