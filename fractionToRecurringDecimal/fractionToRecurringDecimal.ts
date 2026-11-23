function fractionToDecimal(numerator: number, denominator: number): string {

	if (numerator == 0) return '0'
	let sign = ''
	if (numerator < 0) {

		sign = '-'
		numerator = Math.abs(numerator)


	}
	if (denominator < 0) {
		sign = sign == '-' ? '' : '-'
		denominator = Math.abs(denominator)
	}
	const n1 = Math.floor(numerator / denominator)
	if (numerator % denominator === 0) return sign + n1
	let f = ''
	const map = new Map()
	let r = numerator % denominator
	while (r > 0) {

		if (map.has(r)) {
			const str = n1 + '.' + f.slice(0, map.get(r)) + '(' + f.slice(map.get(r)) + ')'

			return sign + str
		}
		map.set(r, f.length)
		r *= 10
		while (r < denominator) {
			f += 0
			r *= 10
		}
		f += Math.floor(r / denominator)
		r = r % denominator
	}
	const str = f === '' ? String(n1) : n1 + '.' + f
	return sign + str
}