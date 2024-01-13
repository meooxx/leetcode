package main

func gameOfLife(board [][]int) {

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[row]); col++ {
			count := 0
			for i := max(row-1, 0); i < min(row+2, len(board)); i++ {
				for j := max(col-1, 0); j < min(col+2, len(board[0])); j++ {
					if i != row || j != col {
						count += board[i][j] & 1
					}
				}
			}
			if count == 3 {
				board[row][col] += 2
			} else if count == 2 {
				if board[row][col]&1 == 1 {
					board[row][col] += 2
				}
			}
		}

	}
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[row]); col++ {
			board[row][col] /= 2
		}
	}
}
