function convert(s: string, numRows: number): string {
	let index = 0
	const mp: string[][] = []
	let anwser = ""
	let col = 0

	while (index < s.length) {
		let row = 0
		for (; row < numRows; row++) {
			if (mp[row] === undefined) {
				mp[row] = []
			}
			mp[row][col] = s[index]
			index++

		}

		row = numRows-2
		col++
		while (row > 0) {
			mp[row][col] = s[index]
			row--
			index++
		}
		col++
	}
	mp.forEach((arr) => {
		anwser += arr.filter(Boolean).join('')
	})
	return anwser
};


console.log(convert('PAYPALISHIRING', 3))
console.log(convert('PAYPALISHIRING', 4))