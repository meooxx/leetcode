function convertToTitle(columnNumber: number): string {
	let answer = ''
	while (columnNumber > 0) {
		// n = (A+1)*26^3 +(B+1)*26^2 ... (Z+1)*26^0
		// n-1 = ........                 Z
		columnNumber--
		const idx = columnNumber % 26
		columnNumber = Math.floor(columnNumber / 26)
		answer = String.fromCharCode(65 + idx) + answer
	}
	return answer
};