

function uniquePathsWithObstacles(obstacleGrid: number[][]): number {

	const dp = Array(obstacleGrid[0].length).fill(0)

	for (let i = 0; i < dp.length; i++) {
		if (obstacleGrid[0][i] !== 1) {
			dp[i] = 1
		} else {
			break
		}
	}

	for (let row = 1; row < obstacleGrid.length; row++) {
		dp[0] = obstacleGrid[row][0] === 1 ? 0 : dp[0]
		for (let col = 1; col < obstacleGrid[0].length; col++) {
			if (obstacleGrid[row][col] === 0) {
				dp[col] = dp[col] + dp[col - 1]
			} else {
				dp[col] = 0
			}
		}
	}
	return dp[obstacleGrid[0].length - 1]


};