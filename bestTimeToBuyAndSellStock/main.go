package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{
		7, 1, 2, 4, 5, 6,
	}))
}
func maxProfit(prices []int) int {

	buy := prices[0]
	max := 0
	for _, v := range prices {
		if v < buy {
			buy = v
		}

		if v-buy > max {
			max = v - buy
		}

	}
	return max
}

func maxProfit1(prices []int) int {
	left, right := 0, 1
	max := 0
	for right < len(prices) {
		if prices[left] <= prices[right] {
			if prices[right]-prices[left] > max {
				max = prices[right] - prices[left]
			}
		} else {
			left = right
		}
		right++
	}
	return max
}
func maxProfit2(prices []int) int {
	max := 0
	for i := 0; i < len(prices)-1; i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j]-prices[i] > max {
				max = prices[j] - prices[i]
			}
		}
	}
	return max
}
