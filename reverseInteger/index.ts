
// 123 ==> 321
// 12300 ==> 321
// -123 ==> -321
function reverse(x: number): number {
	const maxInt31 = Math.pow(2, 31) - 1

	if (x > maxInt31 || x < -maxInt31) {
		return 0
	}
	const s = x > 0 ? x.toString() : x.toString().slice(1)

	const digitals = s.split('')
	// x == 0 return 0
	if(digitals.length <2) return x
	let r = ''
	for (let i = 0; i < digitals.length; i++) {
		if (r === "" && digitals[i] === "0") {
			continue
		}
		r = digitals[i] + r

	}

	const result = x > 0 ? parseInt(r, 10) : -parseInt(r, 10)
	if (result > maxInt31 || result < -maxInt31) {
		return 0
	}
	return result
}