function totalNQueens(n: number): number {
	const answer: { count: number } = { count: 0 }
	const board = Array.from(Array(n), _ => Array(n).fill('.'))
	placeQueen(board, 0, answer)
	return answer.count
};


function placeQueen(board: string[][], row: number, result: { count: number }) {
	if (row > board.length) return
	if (row === board.length) {
		result.count++
		return
	}
	for (let i = 0; i < board[0].length; i++) {
		board[row][i] = 'Q'
		if (isValid(board, row, i)) {
			placeQueen(board, row + 1, result)
		}
		board[row][i] = '.'
	}
}

function isValid(board: string[][], row: number, col: number): boolean {

	for (let c = col - 1; c >= 0; c--) {
		if (board[row][c] === 'Q') return false
	}
	for (let r = row - 1; r >= 0; r--) {
		if (board[r][col] === 'Q') return false
	}
	// 135'
	let r = row - 1
	let c = col - 1
	while (r >= 0 && c >= 0) {
		if (board[r][c] === 'Q') return false
		r--
		c--
	}
	r = row - 1
	c = col + 1
	while (r >= 0 && c < board.length) {
		if (board[r][c] === 'Q') return false
		r--
		c++
	}
	return true
}