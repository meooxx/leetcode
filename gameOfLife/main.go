package main

func gameOfLife(board [][]int) {
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[row]); col++ {
			count := 0
			if row > 0 {
				count += board[row-1][col] % 2
				if col > 0 {
					count += board[row-1][col-1] % 2
				}
				if col < len(board[row])-1 {
					count += board[row-1][col+1] % 2
				}
			}
			if row < len(board)-1 {
				count += board[row+1][col] % 2
				if col > 0 {
					count += board[row+1][col-1] % 2
				}
				if col < len(board[row])-1 {
					count += board[row+1][col+1] % 2
				}
			}
			if col > 0 {
				count += board[row][col-1] % 2
			}
			if col < len(board[row])-1 {
				count += board[row][col+1] % 2
			}

			if count == 3 {
				board[row][col] += 2
			} else if count == 2 {
				if (board[row][col] % 2) == 1 {
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
