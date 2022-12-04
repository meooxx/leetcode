package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxProfit(3, []int{ 3, 5, 0, 3, 1, 4, }))
}
func maxProfit(k int, prices []int) int {
	if k >= len(prices)/2 {
		profit := 0
		for i := 1; i < len(prices); i++ {
			if prices[i] > prices[i-1] {
				profit += prices[i] - prices[i-1]
			}
		}
		return profit
	}

	dpHold := make([]int, len(prices)+1)
	dpProfit := make([]int, len(prices)+1)

	max := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}
	dpHold[0] = math.MinInt
	dpProfit[0] = 0
	//  3     5    0  3    1  4
	// -3    -3    0  0    0  0   hold1
	//  0     2    2  3    3  4   profit
	
	// -3   -3     2  2    2  2  hold2
    // 0    2      2  5    5  6  profit = holde2+price[i]
	// 							 如果k=3
	// -3   -3     2  2    4  4 
	// 0     2     2  5    5  8
	for ; k > 0; k-- {
		for i := range prices {
			dpHold[i+1] = max(dpProfit[i+1]-prices[i], dpHold[i])
			dpProfit[i+1] = max(dpHold[i+1]+prices[i], dpProfit[i])
		}
	}
	return dpProfit[len(prices)]
}
