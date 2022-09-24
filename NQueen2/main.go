package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(totalNQueens(4))
}

func totalNQueens(n int) int {
	if n == 1 {
		return n
	}
	if n < 4 {
		return 0
	}
	count := 0
	current := make([][]byte, n)
	for index := range current {
		current[index] = bytes.Repeat([]byte{'.'}, n)
	}
	placeQueen(&count, current, 0, n)
	return count
}

func placeQueen(count *int, current [][]byte, row int, n int) {
	if row == n {
		*count++
		return
	}
	rowData := current[row]
	for i := 0; i < n; i++ {
		rowData[i] = 'Q'
		if isValid(current, row, i) {
			placeQueen(count, current, row+1, n)
		}
		rowData[i] = '.'
	}

}

func isValid(current [][]byte, row, col int) bool {

	for i := row - 1; i >= 0; i-- {
		if current[i][col] == 'Q' {
			return false
		}
	}
	for i, j := row-1, col-1; i >= 0 && j >= 0; {
		if current[i][j] == 'Q' {
			return false
		}
		j--
		i--
	}
	for i, j := row-1, col+1; i >= 0 && j < len(current); {
		if current[i][j] == 'Q' {
			return false
		}
		i--
		j++

	}

	return true
}
