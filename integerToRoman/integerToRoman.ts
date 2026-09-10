

function intToRoman1(num: number): string {
	if (num < 1 || num > 3999) {
		return ''
	}
	let answer = ''
	const map = new Map([
		[1000, 'M'],
		[900, 'CM'],
		[500, "D"],
		[400, 'CD'],
		[100, 'C'],
		[90, 'XC'],
		[50, 'L'],
		[40, "XL"],
		[10, "X"],
		[9, 'IX'],
		[5, 'V'],
		[4, 'IV'],
		[1, 'I']
	])
	let n = 0
	let base = 1000
	while (num > 0) {
		n = ~~(num / base) * base
		for (let i of [9, 5, 4]) {
			if (n >= i * base) {
				answer += map.get(i * base)
				n -= i * base
			}
		}
		if (n >= base) {
			answer += map.get(base)!.repeat(~~(n / base))

		}
		num %= base

		base /= 10
	}
	return answer


};

console.log(intToRoman(3749))