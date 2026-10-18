

// optimation: use 1d array as we only use dp[i-1][j] or dp[i][j-1]
function isInterleave(s1: string, s2: string, s3: string): boolean {

	const l1 = s1.length
	const l2 = s2.length
	if (l1 + l2 !== s3.length) return false

	const dp = Array.from(Array(l1 + 1), _ => Array(l2 + 1))
	dp[0][0] = true

	for (let i = 1; i <= l1; i++) {
		dp[i][0] = s3[i-1] == s1[i - 1] && dp[i - 1][0]

	}
	for (let i = 1; i <= l2; i++) {
		dp[0][i] = s3[i-1] == s2[i - 1] && dp[0][i - 1]
	}
	// aa b     aa  b

	// try every possibilities
	// s1, pick substring which length from 1 to l1
	// s2, pick substring which length from 1 to l2

	// i=>1   try j = 1...l2
	//  =>2   try j = 1...l2
	// ...   
	for (let i = 1; i <= l1; i++) {
		for (let j = 1; j <= l2; j++) {
			dp[i][j] = s3[i + j - 1] == s1[i - 1] && dp[i - 1][j] || s3[i + j - 1] == s2[j - 1] && dp[i][j - 1]
		}
	}

	return dp[l1][l2]

};