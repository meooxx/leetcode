function plusOne(digits: number[]): number[] {
	let remainder = 0
	let index = digits.length - 1
	digits[index] += 1
	do {
		const sum = digits[index] + remainder
		remainder = remainder > 0 ? 1 : 0
		digits[index] = sum % 10
		remainder = ~~(sum / 10)
		index--
	} while (remainder > 0 && index >= 0)
	if (remainder > 0) {
		digits.unshift(1)
	}
	return digits

};