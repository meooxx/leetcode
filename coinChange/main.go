package main

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	// sort.Slice(coins, func(a, b int) bool {
	//     if coins[a] < coins[b] {
	//         return true
	//     }
	//     return false
	// })
	dp[0] = 0
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt
		for j := range coins {
			if i-coins[j] < 0 {
				continue
			}
			if dp[i-coins[j]] != math.MaxInt {
				dp[i] = min(dp[i], 1+dp[i-coins[j]])
			}
		}
	}
	if dp[amount] == math.MaxInt {
		return -1
	}
	return dp[amount]
}
