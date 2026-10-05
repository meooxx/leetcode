function minDistance2(word1: string, word2: string): number {
	return next(word1, 0, word2, 0, 0)
};

function next(word1: string, i: number, word2: string, j: number, step: number): number {

	if (i > word1.length && j > word2.length) {
		return step
	}
	if (word1[i] === word2[j]) {
		return next(word1, i++, word2, j++, step++)
	}
	return Math.min(
		next(word1, i++, word2, j, step++),
		next(word1, i++, word2, j++, step++),
		next(word1, i, word2, j++, step++)
	)
}
//      c   b  d
//   0  [1] 2 
// a 1  1   ?
// b 2  2    
// c				
//     
function minDistance1(word1: string, word2: string): number {
	const dp = Array.from(Array(word1.length + 1), () => Array(word2.length + 1))

	dp[0][0] = 0
	for (let i = 1; i <= word1.length; i++) {
		dp[i][0] = i
	}
	for (let i = 1; i <= word2.length; i++) {
		dp[0][i] = i
	}
	for (let i = 1; i <= word1.length; i++) {
		for (let j = 1; j <= word2.length; j++) {
			if (word1[i - 1] === word2[j - 1]) {
				dp[i][j] = dp[i - 1][j - 1]
			} else {
				dp[i][j] = Math.min(dp[i - 1][j], dp[i - 1][j - 1], dp[i][j - 1]) + 1
			}
		}
	}
	return dp[word1.length][word2.length]

};
//      c              b            d
//   0  1(i-1,j-1)     2 (i-1, j)
// a 1  1 (i, j-1)     ?
// b 2  2 
// d   
// we need to save (i-1, j-1) before we actually change it
//
//   abd vs abb     find the minmum of there cases + 1
//try  1 delete  dp[...] = dp[ab][ab]
//     2 replace dp[...] = dp[ab][ab]
//     3 add one dp[...] = d[ab][abb]
// https://leetcode.com/problems/edit-distance/solutions/25846/c-o-n-space-dp
function minDistance(word1: string, word2: string): number {
	const dp = Array(word2.length + 1)
	for (let i = 0; i <= word2.length; i++) {
		dp[i] = i
	}
	for (let i = 1; i <= word1.length; i++) {
		let pre = dp[0]
		dp[0] = i
		for (let j = 1; j <= word2.length; j++) {
			// save original dp[j] value in pre,
			// as it is dp[i-1][j-1] for the next loop
			// d[j-1] === dp[i][j-1]
			/// d[j]  === dp[i-1][j]
			const tmp = dp[j]
			if (word1[i - 1] === word2[j - 1]) {
				dp[j] = pre
			} else {
				dp[j] = Math.min(dp[j - 1], dp[j], pre) + 1
			}
			pre = tmp
		}
	}
	return dp[word2.length]

}
