function longestPalindrome(s: string): string {
	const dp: boolean[][] = []
	let ml = s[0]
	//  a === a check bab which is dp[i-1][j+1] 
	//   a b a b a
	// a T F T F
	// b   T F  
	// a     T
	for (let i = 0; i < s.length; i++) {
		dp.push([] as boolean[])
		dp[i][i] = true
	}
	for (let j = 1; j < s.length; j++) {
		for (let i = 0; i < j; i++) {
			if (s[i] == s[j]) {
				if (j - i + 1 === 2) {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i + 1][j - 1]
				}
			} else {
				dp[i][j] = false
			}
			if (dp[i][j] && j - i + 1 > ml.length) {
				ml = s.slice(i, j + 1)
			}
		}

	}
	return ml


}