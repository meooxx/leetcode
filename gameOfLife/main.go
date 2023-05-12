package main

// import "fmt"

func main() {
	gameOfLife([][]int{
		{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0},
	})
}

func gameOfLife1(board [][]int) {
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

// board[i][j] == 0, if live,  0 + 2 , before value = 2 % 2 == 0
// board[i][j] == 1, if live,  1 + 2,  before value = 3 % 2 == 1

func gameOfLife(board [][]int) {
	m := len(board)
	n := len(board[0])
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
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			count := 0
			for i := max(row-1, 0); i < min(row + 2, m); i++ {
				for j := max(col-1, 0); j < min(col+2 , n); j++ {
					if i != row || j != col {
						count += board[i][j] & 1
					}
				}
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