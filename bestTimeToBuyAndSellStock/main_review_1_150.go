package main

func maxProfit(prices []int) int {
	m := prices[0]
	anwser := 0
	for i:=1;i<len(prices);i++ {
		if prices[i] < m {
			m = prices[i]
		}
		if prices[i] - m > anwser {
			anwser = prices[i]  - m
		}
	}
	return anwser
}