


//     112
//  i means length of substring
// 11 2 is valid dp[3]+=dp[2]
// 1 12 is valid dp[3] += dp[1]

function numDecodings(s: string): number {
	if (s[0] == '0') return 0
	if (s.length == 1) return 1
	// 2 before current
	let dp2 = 1
	// 1 before current
	let dp1 = 1

	for (let i = 2; i <= s.length; i++) {
		let curr = 0
		if (isValid(s[i - 1])) {
			curr = dp1
		}
		if (isValid(s.slice(i - 2, i))) {
			curr += dp2
		}
		dp2 = dp1
		dp1 = curr
	}
	return dp1




};

function isValid(s: string): boolean {
	if (s.length > 1 && s[0] === '0') return false
	return Number(s) >= 1 && Number(s) <= 26
}
