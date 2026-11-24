

//
//              infinity
//    10,30,-5  1
//          infinity

function calculateMinimumHP(dungeon: number[][]): number {
	const dp = Array(dungeon[0].length + 1).fill(Infinity)
	dp[dp.length - 1] = 1
	for (let row = dungeon.length - 1; row >= 0; row--) {
		if (row < dungeon.length - 1) dp[dp.length - 1] = Infinity
		for (let col = dungeon[0].length - 1; col >= 0; col--) {
			dp[col] = Math.min(-dungeon[row][col] + dp[col + 1], dp[col])
			dp[col] = Math.max(1, dp[col])
		}
	}
	return dp[0]
};