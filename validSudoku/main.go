package main

import "fmt"

func main() {

	sudo := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'1', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	fmt.Println(isValidSudoku(sudo))
}

// 对数独内数字按 row, col, block 三个维度做编码处理, 储存三个值
// 即一个数字对应三个带位置信息的值, 所在行, 列, block 进行判断重复
//  1 2 3
//  4 5 6
func isValidSudoku(nums [][]byte) bool {
	seen := map[string]int{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {

			num := string(nums[i][j])
			if nums[i][j] == '.' {
				continue
			}
			rowChar := num + fmt.Sprint(i) + "row"
			colChar := num + fmt.Sprint(j) + "col"
			blockChar := num + fmt.Sprint(i/3, j/3)
			// row
			seen[rowChar]++
			// column
			seen[colChar]++
			// block
			seen[blockChar]++
			if seen[rowChar] > 1 || seen[colChar] > 1 || seen[blockChar] > 1 {
				return false
			}
		}
	}
	return true
}
