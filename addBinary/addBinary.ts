


function addBinary(a: string, b: string): string {
	if (a === '0') return b
	if (b === '0') return a
	let aIndex = a.length - 1
	let bIndex = b.length - 1
	let answer = ''
	let reminder = 0
	const zero = '0'.charCodeAt(0)
	while (aIndex >= 0 || bIndex >= 0) {
		const n1 = aIndex >= 0 ? a[aIndex].charCodeAt(0) - zero : 0
		const n2 = bIndex >= 0 ? b[bIndex].charCodeAt(0) - zero : 0
		const result = n1 + n2 + reminder
		const digit = result % 2
		reminder = ~~(result / 2)
		answer = digit + answer
		aIndex--
		bIndex--
	}
	answer = reminder > 0 ? '1' + answer : answer
	return answer

};