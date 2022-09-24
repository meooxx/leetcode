package main

import "fmt"

func main() {
	// sudoKu := [][]byte{
	// 	[]byte{'5', '3', '4', '6', '7', '8', '9', '1', '.'},
	// 	[]byte{'6', '7', '2', '1', '9', '5', '3', '4', '8'},
	// 	[]byte{'1', '9', '8', '.', '.', '.', '.', '6', '.'},
	// 	[]byte{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
	// 	[]byte{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
	// 	[]byte{'7', '1', '3', '9', '2', '4', '8', '5', '6'},
	// 	[]byte{'9', '6', '1', '5', '3', '7', '2', '8', '4'},
	// 	[]byte{'2', '8', '7', '4', '1', '9', '6', '3', '5'},
	// 	[]byte{'3', '4', '5', '2', '8', '6', '1', '7', '9'},
	// }
	sudoKu := [][]byte{
		[]byte{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		[]byte{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		[]byte{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		[]byte{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		[]byte{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		[]byte{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		[]byte{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		[]byte{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		[]byte{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
	}
	solveSudoku(sudoKu)
	for i := range sudoKu {
		for _, v := range sudoKu[i] {
			fmt.Printf("%s", string(v))
		}
		fmt.Println()
	}
}

func solveSudoku(board [][]byte) {
	solve(board)
}

func solve(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				for c := byte('1'); c <= '9'; c++ {
					if isValid(board, i, j, c) {
						board[i][j] = c
						if solve(board) {
							return true
						} else {
							board[i][j] = '.'
						}
					}
				}

				return false
			}
		}
	}
	return true

}

func isValid(board [][]byte, i, j int, c byte) bool {
	for n := 0; n < 9; n++ {
		if board[n][j] == c {
			return false
		}
		if board[i][n] == c {
			return false
		}
		// 1 2 3  456  7
		// 4 5 6
		// 7 8 9

		// 5
		if board[(i/3)*3+n/3][(j/3)*3+n%3] != '.' && board[(i/3)*3+n/3][(j/3)*3+n%3] == c {
			return false
		}
	}
	return true
}
