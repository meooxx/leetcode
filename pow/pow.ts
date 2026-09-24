function myPow(x: number, n: number): number {
	if (n < 0) return myPow(1 / x, -n)
	if (x === 1 || n === 0) return 1
	let answer = 1
	let exp = BigInt(n)

	while (exp > 0) {
		if ((exp & 1n) === 1n) {
			answer *= x
		}
		x *= x
		exp >>= 1n
	}

	return answer
};

