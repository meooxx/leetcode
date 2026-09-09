function myAtoi(s: string): number {
	let answer = 0
	// 0014   -014
	let sign = 1
	const int0 = '0'.charCodeAt(0)
	// const int9 = '9'.charCodeAt(0)
	let signFound = false
	let invalid = false
	let intFound = false
	for (const char of s) {
		switch (true) {
			case char === '_':
			case char === ' ':
				if (signFound || intFound) {
					invalid = true

				}
				break
			// .charCodeAt(0)
			case char === '-':
				if (signFound || intFound) {
					invalid = true
					break
				}
				signFound = true
				sign = -1
				break
			case char === "+":
				if (signFound || intFound) {
					invalid = true
					break
				}
				signFound = true
				sign = 1
				break
			case char >= '0' && char <= '9':
				answer = answer * 10 + char.charCodeAt(0) - int0
				intFound = true
				break
			default:
				invalid = true
				break
		}
		if (invalid) {
			break
		}
	}

	answer *= sign
	if (answer > Math.pow(2, 31) - 1) {
		return Math.pow(2, 31) - 1
	}
	if (answer < -Math.pow(2, 31)) {
		return -Math.pow(2, 31)
	}
	return answer
};

