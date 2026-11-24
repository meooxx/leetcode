function titleToNumber(columnTitle: string): number {
	let base = 1
	let result = 0
	const codeA = 'A'.charCodeAt(0) - 1
	for (let i = columnTitle.length - 1; i >= 0; i--) {
		const idx = columnTitle[i].charCodeAt(0) - codeA
		result += idx * base
		base *= 26
	}
	return result
};