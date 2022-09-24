function longestPalindromicSub(s: string): string {
	const dp: boolean[][] = [];
	let maxLen = s[0];
	for (let i = 0; i < s.length; i++) {
		dp.push([] as boolean[]);
		dp[i][i] = true;
	}
	// dp[i][j] = dp[i+1][j-1] && dp[i]==dp[j]
	// left i   i+1  j-1   j right
	//     	a    b   b     a
	// a   	T    t1  t2
	// b         T   t3
	// b             T
	// a                   T
	//  竖着填写才能复用已经计算出来的值, abba参考bb, a[0][3] 参考 a[1][2]
	//  当j确定时, 填写j那一列如 j == 2
	//  i(0-2)填写t2, t3 即a[i0][j], a[i1][j]
	//  dp[i][j]
	for (let j = 1; j < s.length; j++) {
		for (let i = 0; i < j; i++) {
			if (s[i] == s[j]) {
				if (j - i + 1 <= 3) {
					dp[i][j] = true;
				} else {
					dp[i][j] = dp[i + 1][j - 1];
				}
			}
			if (dp[i][j] && j - i + 1 > maxLen.length) {
				maxLen = s.slice(i, j + 1);
			}
		}
	}

	return maxLen;
}
