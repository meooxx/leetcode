package main

func numIslands(grid [][]byte) int {
	rowLen := len(grid)
	colLen := len(grid[0])

	count := 0

	var check func(row, col int) bool
	check = func(row, col int) bool {
		if row >= rowLen || col >= colLen || row < 0 || col < 0 {
			return true
		}
		if grid[row][col] == '0' || grid[row][col] == '#' {
			return true
		}
		grid[row][col] = '#'
		return check(row, col+1) && check(row, col-1) && check(row+1, col) && check(row-1, col)

	}

	for row := 0; row < rowLen; row++ {
		for col := 0; col < colLen; col++ {
			if grid[row][col] == '1' {
				if check(row, col) {
					count++
				}
			}
		}
	}
	return count
}
