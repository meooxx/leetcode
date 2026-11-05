


// transition equaltation


function numDistinct(s: string, t: string): number {
	const dp = Array.from({
		length: s.length + 1,
	}, () => {
		return Array(t.length + 1).fill(0)
	})

	for (let i = 0; i < s.length; i++) {
		dp[i][0] = 1
	}

	for (let row = 1; row <= s.length; row++) {
		for (let col = 1; col <= t.length; col++) {
			if (s[row - 1] == t[col - 1]) {
				dp[row][col] = dp[row - 1][col - 1] + dp[row - 1][col]
			} else {
				dp[row][col] = dp[row - 1][col]
			}
		}
	}
	return dp[s.length][t.length]
};