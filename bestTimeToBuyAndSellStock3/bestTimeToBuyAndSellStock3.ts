function maxProfit(prices: number[]): number {
	if (prices.length == 0) return 0
	const dp = Array.from({ length: 3 }, () => Array(prices.length).fill(0))
	let k = 1
	while (k <= 2) {
		let profit = -prices[0]
		for (let i = 1; i < prices.length; i++) {
			// keypoint: dp[k-1][i] is accumulative profit from pre days
			profit = Math.max(profit, -prices[i] + dp[k - 1][i - 1])
			dp[k][i] = Math.max(dp[k][i - 1], prices[i] + profit)
		}
		k++
	}
	return dp[2][prices.length - 1]

};  