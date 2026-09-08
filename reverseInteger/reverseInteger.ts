function reverse(x: number): number {
	const maxInt31 = Math.pow(2, 31)
	if (x > maxInt31 - 1 || x < -maxInt31) {
		return 0
	}
	let sign = 1
	if (x < 0) {
		sign = -1
		x = -x
	}
	let anwser = 0
	while (x > 0) {

		anwser = anwser * 10 + x % 10
		// x == ~~x
		x = ~~(x / 10)
	}
	anwser *= sign
	if (anwser <= maxInt31 - 1 && anwser >= -maxInt31) {
		return anwser
	}
	return 0
};

console.log(reverse(-123))
console.log(reverse(123))