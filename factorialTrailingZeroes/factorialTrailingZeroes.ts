function trailingZeroes(n: number): number {

	// 0 comes from 10
	// which can written 2*5
	// number of 2 is more than that of 5
	// we need count for 5s
	// while 25 contribute to a extra 5
	// so does 125


	let sum = 0
	while (n > 0) {
		n = Math.floor(n / 5)
		sum += n

	}
	return sum
}