


function minPathSum(grid: number[][]): number {
	if (grid.length === 0) return 0
	const dp = Array(grid[0].length)
	dp[0] = grid[0][0]
	for (let i = 1; i < dp.length; i++) {
		dp[i] = dp[i - 1] + grid[0][i]
	}


	for (let i = 1; i < dp.length; i++) {
		dp[i] = grid[0][i] + dp[i - 1]
	}
	for (let row = 1; row < grid.length; row++) {
		dp[0] = dp[0] + grid[row][0]
		for (let col = 1; col < dp.length; col++) {
			dp[col] = grid[row][col] + Math.min(dp[col], dp[col - 1])
		}
	}
	return dp[dp.length - 1]
};