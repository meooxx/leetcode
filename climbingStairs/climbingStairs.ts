function climbStairs(n: number): number {

	if (n <= 3) return n
	let n1 = 3
	let n2 = 2

	for (let i = 4; i <= n; i++) {
		[n2, n1] = [n1, n2 + n1]
	}
	return n1

};