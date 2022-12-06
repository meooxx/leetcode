package main

// it only require number of island
// so, we can easily modify '1' to '0'
func numIslands(grid [][]byte) int {
	rows := len(grid)
	cols := len(grid[0])
	count := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				if checkNext(grid, i, j) {
					count++
				}
			}
		}
	}
	return count
}

func checkNext(grid [][]byte, row, col int) bool {
	if row >= len(grid) || col >= len(grid[0]) || row < 0 || col < 0 {
		return true
	} else {
		// equals !grid[row][col]
		// may move to if statement
		if grid[row][col] == '0' || grid[row][col] == '#' {
			return true
		}
		if grid[row][col] == '1' {
			grid[row][col] = '#'
			checked := checkNext(grid, row+1, col) && checkNext(grid, row-1, col) && checkNext(grid, row, col+1) && checkNext(grid, row, col-1)
			return checked
		}
		return false

	}
}
