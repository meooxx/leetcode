
// dp[i][j]
// i, 0~i substring of s
// j, 0-j substring of p

//     p
//        a * b *
// s    t f t f t 
//      a
//      b

function isMatch(s: string, p: string): boolean {
	const dp = Array.from({ length: s.length + 1 }, _ => Array(p.length + 1).fill(false))
	// '' match '', which is true
	dp[0][0] = true

	// using a char+* (e.g. a*) to match ''
	for (let pIndex = 1; pIndex < p.length; pIndex++) {
		if (p[pIndex] === "*") {
			dp[0][pIndex + 1] = dp[0][pIndex - 1]
		}
	}
	// using substring of pattern to match substring of s
	for (let pIndex = 0; pIndex < p.length; pIndex++) {
		for (let sIndex = 0; sIndex < s.length; sIndex++) {
			// e.g. 'ab' match 'a.'  only need to check a vs a
			if (s[sIndex] === p[pIndex] || p[pIndex] === '.') {
				dp[sIndex + 1][pIndex + 1] = dp[sIndex][pIndex]
			} else if (p[pIndex] === '*') {
				// 
				// * match 0(e.g. ab vs abc*, only check ab vs ab)
				//  or manytimes
				//  only when '.' or p[pIndex] == s[sIndex]
				//  aa vs a* is true because a* and a are matched
				dp[sIndex + 1][pIndex + 1] = dp[sIndex + 1][pIndex - 1] || (
					(s[sIndex] === p[pIndex - 1] || p[pIndex-1] === '.') && dp[sIndex][pIndex + 1]
				)
			}
		}
	}
	return dp[s.length][p.length]
};

console.log(isMatch('aa',
	'.*'
))