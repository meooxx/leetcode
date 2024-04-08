package main

func numIslands(grid [][]byte) int {
	count := 0
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == '1' {
				if checkNext(grid, row, col) {
					count++
				}
			}
		}
	}
	return count

}
func checkNext(grid [][]byte, row, col int) bool {
	if row >= len(grid) || row < 0 || col >= len(grid[row]) || col < 0 {
		return true
	}

	if grid[row][col] == '0' || grid[row][col] == '#' {
		return true
	}
	if grid[row][col] == '1' {
		grid[row][col] = '#'
		return checkNext(grid, row, col+1) && checkNext(grid, row, col-1) && checkNext(grid, row-1, col) && checkNext(grid, row+1, col)
	}

	return false

}
