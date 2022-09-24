package main

import "fmt"

func main() {
	b := [][]byte{
		{'a', 'b', 'c', 'd'},
		{'e', 'f', 'g', 'h'},
		{'h', 'i', 'j', 'k'},
	}
	fmt.Println(exist(b, "abc"))
	c := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	fmt.Println(exist(c, "ABCEF"))
	d := [][]byte{
		{'A', 'B'},
		{'C', 'D'},
	}
	fmt.Println(exist(d, "ACDB"))

}

func exist(board [][]byte, word string) bool {
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[row]); col++ {
			if nextByte(&board, row, col, word) {
				return true
			}
		}
	}
	return false
}

func nextByte(b *[][]byte, row, col int, word string) bool {
	board := *b
	if len(word) == 0 {
		return true
	}
	if row < len(board) && row >= 0 && col < len(board[row]) && col >= 0 {
		if board[row][col] == word[0] {
			board[row][col] = '#'
			pass := nextByte(b, row, col+1, word[1:]) || nextByte(b, row, col-1, word[1:]) || nextByte(b, row+1, col, word[1:]) || nextByte(b, row-1, col, word[1:])
			if pass {
				return pass
			}
			board[row][col] = word[0]
			return false
		}
		return false
	}
	return false

}
