function generateMatrix(n: number): number[][] {
	const answer: number[][] = Array.from(Array(n), _ => Array(n))

	let left = 0
	let top = 0
	let right = n - 1
	let bottom = n - 1

	let count = 1;
	while (left <= right && top <= bottom) {
		let l = left
		let r = right
		while (l <= r) {
			answer[top][l] = count
			count++
			l++
		}
		let t = top + 1
		while (t <= bottom) {
			answer[t][right] = count
			count++
			t++
		}
		r = right - 1
		l = left
		while (r >= l && bottom > top) {
			answer[bottom][r] = count
			count++
			r--
		}
		let b = bottom - 1
		while (b > top && right > left) {
			answer[b][left] = count
			count++
			b--
		}
		top++
		left++
		right--
		bottom--
	}

	return answer

};