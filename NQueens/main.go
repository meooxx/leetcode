package main

import (
	"bytes"
	"fmt"
)

func main() {
	v := solveNQueens(4)

	for index := range v {
		fmt.Println(v[index])
	}

	fmt.Printf("%#v", solveNQueens(1))
}

// //  每行的开始尝试放Q
// //
func solveNQueens(n int) [][]string {
	result := [][]string{}
	if n < 4 && n != 1 {
		return result
	}
	current := make([][]byte, n)
	for index := range current {
		current[index] = bytes.Repeat([]byte{'.'}, n)

	}
	placeQueen(&result, current, 0, n)
	return result
}

func placeQueen(result *[][]string, current [][]byte, row int, n int) {
	if row == n {
		ans := []string{}
		for i := 0; i < n; i++ {
			ans = append(ans, string(current[i]))
		}
		*result = append(*result, ans)

		return
	}
	rowData := current[row]
	for i := 0; i < n; i++ {
		rowData[i] = 'Q'
		if isValid(current, row, i) {
			placeQueen(result, current, row+1, n)
		}
		rowData[i] = '.'
	}

}

func isValid(current [][]byte, row, col int) bool {
	//  一行只放了一个, 只检查列
	// 从当前往上检查
	for i := row - 1; i >= 0; i-- {
		if current[i][col] == 'Q' {
			return false
		}
	}
	// 135'
	// 只检查前面, 后面还没放
	for i, j := row-1, col-1; i >= 0 && j >= 0; {
		if current[i][j] == 'Q' {
			return false
		}
		i--
		j--
	}
	// 45' 方向
	for i, j := row-1, col+1; i >= 0 && j < len(current); {
		if current[i][j] == 'Q' {
			return false
		}
		i--
		j++
	}
	return true

}
