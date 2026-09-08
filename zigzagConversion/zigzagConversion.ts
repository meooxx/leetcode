function convert(s: string, numRows: number): string {
	if (numRows === 1) {
		return s
	}
	let index = 0
	const mp: string[] = Array(numRows).fill('')
	let direction = 1
	let row = 0
	while (index < s.length) {
		mp[row] += s[index]
		index++
		if (row === numRows - 1) {
			direction = -1
		}
		if (row === 0) {
			direction = 1
		}
		row += direction
	}
	return mp.join('')

};


console.log(convert('PAYPALISHIRING', 3))
console.log(convert('PAYPALISHIRING', 4))