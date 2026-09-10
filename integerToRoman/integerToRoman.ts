

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

function intToRoman(num: number): string {
	const map = new Map([
		[3000, 'MMM'],
		[2000, 'MM'],
		[1000, 'M'],
		[900, 'CM'],
		[800, 'DCCC'],
		[700, 'DCC'],
		[600, 'DC'],
		[500, "D"],
		[400, 'CD'],
		[300, 'CCC'],
		[200, 'CC'],
		[100, 'C'],
		[90, 'XC'],
		[80, 'LXXX'],
		[70, 'LXX'],
		[60, 'LX'],
		[50, 'L'],
		[40, "XL"],
		[30, 'XXX'],
		[20, 'XX'],
		[10, "X"],
		[9, 'IX'],
		[8, 'VIII'],
		[7, 'VII'],
		[6, 'VI'],
		[5, 'V'],
		[4, 'IV'],
		[3, 'III'],
		[2, 'II'],
		[1, 'I']])
	let base = 1
	let answer = ''
	// 74
	while (num > 0) {
		const digit = num % 10
		answer = (map.get(digit * base) ?? '') + answer
		num = ~~(num / 10)
		base *= 10
	}
	return answer
};

console.log(intToRoman(3749))