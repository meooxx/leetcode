function countPrimes(n: number): number {
	const primes = Array(n).fill(false)
	let answer = 0
	// 1,2,3
	for (let i = 2; i < n; i++) {
		if (!primes[i]) {
			answer++
		}
		for (let j = 2; i * j < n; j++) {
			primes[i * j] = true
		}
	}
	return answer
};