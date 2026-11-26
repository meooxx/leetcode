function reverseBits(n: number): number {
	let answer = 0
	for (let i = 0; i < 32; i++) {
		const bit = n & 1
		answer <<= 1
		answer |= bit
		n >>= 1

	}
	return answer
};