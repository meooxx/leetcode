function minCut(s: string): number {
	//	     l -1  		[i]
	// 0  xxx [j] a b  a 
	//            l    r 
	//     [ str[l] == str[r] is palindrome ]
	//  r+1 r indicates index i in s, cuase Array[s.length+1]
	//   dp[i] = min(dp[j], dp[i])
	// == dp[r+1] = min(dp[l-1+1], dp[r+1])
	if (s.length <= 1) return 0
	const dp = Array.from({ length: s.length + 1 }, (_, index) => index - 1)

	for (let i = 0; i < s.length; i++) {
		// odd
		let l = i
		let r = i
		while (l >= 0 && r < s.length && s[l] == s[r]) {

			dp[r + 1] = Math.min(dp[r + 1], 1 + dp[l])
			l--
			r++
		}
		// even
		l = i
		r = i + 1
		// abba
		while (l >= 0 && r < s.length && s[l] == s[r]) {

			dp[r + 1] = Math.min(dp[r + 1], 1 + dp[l])
			l--
			r++
		}
	}
	return dp[s.length]
};