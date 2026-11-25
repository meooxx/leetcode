function maxProfit(k: number, prices: number[]): number {
	const dp = Array.from({
		length: k + 1,
	}, () => Array(prices.length).fill(0))
	let i = 1
	while (i <= k) {
		// keep track of the maxProfit(pre- minPrice)
		let profit = -prices[0]
		for (let pIndex = 1; pIndex < prices.length; pIndex++) {
			profit = Math.max(profit, -prices[pIndex] + dp[i - 1][pIndex - 1])
			// update dp = lastAccumulative + price or don't trade
			dp[i][pIndex] = Math.max(dp[i][pIndex - 1], profit + prices[pIndex])
		}
		i++
	}
	return dp[k][prices.length - 1]
};