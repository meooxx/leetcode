
/// take this example

// 57/5 = (10) * 5 +2
//      = 0x(1011) * 5 + 2
//      =    (2^3+2^1+2^0) * 5
//      = 2^3 *5 + 2^1 *5 + 2^0 *5
function divide(dividend: number, divisor: number): number {
	let answer = 0n
	let negative = false

	if (dividend < 0) {
		negative = !negative
		dividend = Math.abs(dividend)
	}
	if (divisor < 0) {
		negative = !negative
		divisor = Math.abs(divisor)
	}
	let div = BigInt(dividend)
	let divs = BigInt(divisor)
	// 9/4
	while (div >= divs) {
		let exponent = 0n
		while (div >= divs << exponent) {
			exponent++
		}

		answer += 1n << (exponent - 1n)
		div -= divs << (exponent - 1n)
	}
	answer = negative ? -answer : answer
	if (answer > Math.pow(2, 31) - 1) {
		return Math.pow(2, 31) - 1
	}
	if (answer < -Math.pow(2, 31)) {
		return -Math.pow(2, 31)
	}
	return Number(answer)
};

divide(2147483647, 2)