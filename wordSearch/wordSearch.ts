

function exist(board: string[][], word: string): boolean {

	for (let row = 0; row < board.length; row++) {
		for (let col = 0; col < board[0].length; col++) {
			if (board[row][col] == word[0]) {
				if (next(board, row, col, word)) {
					return true
				}
			}
		}
	}
	return false
};

function next(board: string[][], row: number, col: number, word: string): boolean {
	if (word === '') {
		return true
	}
	if (row < 0 || row >= board.length || col < 0 || col >= board[0].length) {
		return false

	}


	if (board[row][col] === word[0]) {
		board[row][col] = '#'
		if (next(board, row, col + 1, word.slice(1)) || next(board, row, col - 1, word.slice(1)) || next(board, row + 1, col, word.slice(1)) || next(board, row - 1, col, word.slice(1))) {
			return true
		}
		board[row][col] = word[0]
	}

	return false


}