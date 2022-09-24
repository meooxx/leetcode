package main

import "fmt"

func main() {
	b1 := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	}
	solve(b1)
	fmt.Println(b1)

	b2 := [][]byte{
		{'X', 'O', 'X', 'O', 'X', 'O', 'O', 'O', 'X', 'O'},
		{'X', 'O', 'O', 'X', 'X', 'X', 'O', 'O', 'O', 'X'},
		{'O', 'O', 'O', 'O', 'O', 'O', 'O', 'O', 'X', 'X'},
		{'O', 'O', 'O', 'O', 'O', 'O', 'X', 'O', 'O', 'X'},
		{'O', 'O', 'X', 'X', 'O', 'X', 'X', 'O', 'O', 'O'},
		{'X', 'O', 'O', 'X', 'X', 'X', 'O', 'X', 'X', 'O'},
		{'X', 'O', 'X', 'O', 'O', 'X', 'X', 'O', 'X', 'O'},
		{'X', 'X', 'O', 'X', 'X', 'O', 'X', 'O', 'O', 'X'},
		{'O', 'O', 'O', 'O', 'X', 'O', 'X', 'O', 'X', 'O'},
		{'X', 'X', 'O', 'X', 'X', 'X', 'X', 'O', 'O', 'O'},
	}

	solve(b2)
	for _, l := range b2 {
		fmt.Println(string(l))
	}

}

//  边界的 O 可能被修改, 所以找所有边界O相连的 O,先替换成其他字符#
// 最后遍历board, 其他字符#修改成O, 本来O ->  X
func solve(board [][]byte) {
	if len(board) <= 2 {
		return
	}
	if len(board[0]) <= 2 {
		return
	}

	var checkAround func(row, col int)
	checkAround = func(row, col int) {
		if row >= len(board) || col >= len(board[0]) || row < 0 || col < 0 || board[row][col] != 'O' {
			return
		}
		board[row][col] = '#'
		checkAround(row+1, col)
		checkAround(row-1, col)
		checkAround(row, col-1)
		checkAround(row, col+1)
	}
	rowSize := len(board)
	colSize := len(board[0])
	for i := 0; i < rowSize; i++ {
		if board[i][0] == 'O' {
			checkAround(i, 0)
		}

		if board[i][colSize-1] == 'O' {
			checkAround(i, colSize-1)
		}
	}
	for i := 1; i < colSize-1; i++ {
		if board[0][i] == 'O' {
			checkAround(0, i)
		}
		if board[rowSize-1][i] == 'O' {
			checkAround(rowSize-1, i)
		}
	}

	for row := 0; row < rowSize; row++ {
		for col := 0; col < colSize; col++ {
			if board[row][col] == '#' {
				board[row][col] = 'O'
			} else if board[row][col] == 'O' {
				board[row][col] = 'X'
			}
		}
	}

}
func solve3(board [][]byte) {
	if len(board) <= 2 {
		return
	}
	if len(board[0]) <= 2 {
		return
	}
	for row := 1; row < len(board); row++ {
		for col := 1; col < len(board[row]); col++ {
			if board[row][col] == 'O' {
				if board[row][col-1] == 'O' || board[row-1][col] == 'O' {
					continue
				}
				path := [][]int{}
				checkAround(&board, row, col, &path)
				for _, cor := range path {
					board[cor[0]][cor[1]] = 'X'
				}
			}
		}
	}

}

func checkAround(b *[][]byte, row, col int, path *[][]int) bool {
	board := *b

	if row <= 0 || col <= 0 || row >= len(board)-1 || col >= len(board[0])-1 {
		return false
	}
	board[row][col] = 'X'
	if board[row][col+1] != 'O' && board[row][col-1] != 'O' && board[row-1][col] != 'O' && board[row+1][col] != 'O' {
		board[row][col] = 'O'
		*path = append(*path, []int{row, col})
		return true
	}
	p := [][]int{{row, col}}

	if board[row][col+1] == 'O' {
		if !checkAround(b, row, col+1, &p) {
			board[row][col] = 'O'
			return false
		}
	}
	if board[row][col-1] == 'O' {
		if !checkAround(b, row, col-1, &p) {
			board[row][col] = 'O'
			return false
		}
	}
	if board[row-1][col] == 'O' {
		if !checkAround(b, row-1, col, &p) {
			board[row][col] = 'O'
			return false
		}
	}
	if board[row+1][col] == 'O' {
		if !checkAround(b, row+1, col, &p) {
			board[row][col] = 'O'
			return false
		}
	}
	*path = append(*path, p...)
	board[row][col] = 'O'
	return true
}
