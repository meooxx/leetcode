function wordBreak(s: string, wordDict: string[]): boolean {
	const wordSet = new Set(wordDict)
	const dp: boolean[] = Array(s.length + 1).fill(false)
	dp[0] = true
	//   leetcode
	// 0 j
	for (let i = 0; i < s.length; i++) {
		for (let j = i + 1; j <= s.length; j++) {
			const w = s.slice(i, j)
			// && dp[i] to avoid ture value overided by a false value
			// aaa [aaaa] is true
			// aaa(a) [aaa] turn it to false
			//     i
			if (wordSet.has(w) && dp[i]) {
				dp[j] = dp[i]
			}
		}
	}
	return dp[s.length]
};