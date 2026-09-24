//  1   2   3     4   
// 13   14  15    5
// 12   17  16    6
// 11   10  9    8  
function spiralOrder(matrix: number[][]): number[] {

	let left = 0
	let right = matrix[0].length - 1
	let top = 0
	let bottom = matrix.length - 1
	const answer: number[] = []
	while (top <= bottom && left <= right) {
		let l = left
		let r = right
		// >>>>
		while (l <= r) {
			answer.push(matrix[top][l])
			l++
		}
		// >>>>> 	 |
		//				 V
		let t = top + 1
		let b = bottom
		while (t <= b) {
			answer.push(matrix[t][right])
			t++
		}
		l = left
		r = right - 1
		// <<<<
		while (r >= l &&  bottom > top) {
			answer.push(matrix[bottom][r])
			r--
		}
		t = top
		b = bottom - 1
		while (b > t && left < right) {
			answer.push(matrix[b][left])
			b--
		}

		top++
		left++
		right--
		bottom--
	}

	return answer
};


spiralOrder([[1, 2, 3, 4], [5, 6, 7, 8], [9, 10, 11, 12]])