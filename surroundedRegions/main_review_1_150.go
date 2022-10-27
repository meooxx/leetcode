package main

// X O X
// O X O
// X O X
func solve(board [][]byte) {
	rowSize := len(board)
	colSize := len(board[0])

	var mark func(row, col int)
	mark = func(row, col int) {
		if row < 0 || row >= rowSize || col < 0 || col >= colSize {
			return
		}
		if board[row][col] == 'O' {
			board[row][col] = '#'
			mark(row, col+1)
			mark(row, col-1)
			mark(row+1, col)
			mark(row-1, col)
		}
	}

	for i := 0; i < rowSize; i++ {
		if board[i][0] == 'O' {
			mark(i, 0)
		}
		if board[i][colSize-1] == 'O' {
			mark(i, colSize-1)
		}
	}
	for i := 1; i < colSize; i++ {
		if board[0][i] == 'O' {
			mark(0, i)
		}
		if board[rowSize-1][i] == 'O' {
			mark(rowSize-1, i)
		}
	}
	for i := 0; i < rowSize; i++ {
		for j := 0; j < colSize; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == '#' {
				board[i][j] = 'O'
			}
		}
	}

}
