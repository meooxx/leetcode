package main

func maxProfit(prices []int) int {

	buy := -prices[0]
	sell := 0
	cooldown := 0
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := range prices {
		maxbuy := max(buy, cooldown-prices[i])
		maxsell := max(sell, buy+prices[i])
		cooldown = max(cooldown, sell)

		buy = maxbuy
		sell = maxsell
	}
	return max(cooldown, sell)
}
