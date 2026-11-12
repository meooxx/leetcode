/**
 Do not return anything, modify board in-place instead.
 */
function solve(board: string[][]): void {
	if (board.length == 0) return
	let rows = board.length
	let cols = board[0].length
	const checkRound = (b: string[][], row: number, col: number) => {
		if (row < 0 || col < 0) return
		if (row >= b.length || col >= b[0].length) return
		if (b[row][col] === 'O') {
			b[row][col] = '*'
			checkRound(b, row - 1, col)
			checkRound(b, row + 1, col)
			checkRound(b, row, col + 1)
			checkRound(b, row, col - 1)
		}

	}

	for (let i = 0; i < rows; i++) {
		if (board[i][0] == 'O') checkRound(board, i, 0)
		if (board[i][cols - 1] == 'O') checkRound(board, i, cols - 1)
	}
	for (let i = 0; i < cols; i++) {
		if (board[0][i] == 'O') checkRound(board, 0, i)
		if (board[rows - 1][i] == 'O') checkRound(board, rows - 1, i)
	}

	for (let row = 0; row < board.length; row++) {
		for (let col = 0; col < board[0].length; col++) {

			if (board[row][col] === 'O') {
				board[row][col] = 'X'
			}
			if (board[row][col] === '*') {
				board[row][col] = 'O'
			}
		}
	}
};
