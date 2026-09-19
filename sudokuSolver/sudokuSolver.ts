/**
 Do not return anything, modify board in-place instead.
 */
function solveSudoku(board: string[][]): void {
	solve(board)
};

function solve(board: string[][]): boolean {
	for (let row = 0; row < 9; row++) {
		for (let col = 0; col < 9; col++) {
			if (board[row][col] === '.') {
				// try 1-9
				for (let n = 1; n <= 9; n++) {
					if (isValid(board, row, col, String(n))) {
						board[row][col] = String(n)
						if (solve(board)) {
							return true
						} else {
							board[row][col] = '.'
						}
					}


				}
				return false
			}
		}
	}
	return true
}

function isValid(board: string[][], row: number, col: number, n: string): boolean {
	for (let i = 0; i < 9; i++) {
		if (board[row][i] === n) {
			return false
		}
		if (board[i][col] === n) {
			return false
		}

		// row /3 *3 find the start row of the block
		if (board[~~(row / 3) * 3 + ~~(i / 3)][~~(col / 3) * 3 + i % 3] === n) {
			return false
		}
	}
	return true

}

solveSudoku([["5", "3", ".", ".", "7", ".", ".", ".", "."], ["6", ".", ".", "1", "9", "5", ".", ".", "."], [".", "9", "8", ".", ".", ".", ".", "6", "."], ["8", ".", ".", ".", "6", ".", ".", ".", "3"], ["4", ".", ".", "8", ".", "3", ".", ".", "1"], ["7", ".", ".", ".", "2", ".", ".", ".", "6"], [".", "6", ".", ".", ".", ".", "2", "8", "."], [".", ".", ".", "4", "1", "9", ".", ".", "5"], [".", ".", ".", ".", "8", ".", ".", "7", "9"]])