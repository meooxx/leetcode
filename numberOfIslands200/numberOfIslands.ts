function numIslands(grid: string[][]): number {
	let count = 0
	for (let row = 0; row < grid.length; row++) {
		for (let col = 0; col < grid[0].length; col++) {
			if (grid[row][col] == '1') {
				count++
				mark(grid, row, col)
			}
		}
	}
	return count
};

function mark(grid: string[][], row: number, col: number) {
	if (row < 0 || row >= grid.length) {
		return
	}
	if (col < 0 || col >= grid[0].length) {
		return
	}
	if (grid[row][col] == '1') {
		grid[row][col] = '0'
		mark(grid, row + 1, col)
		mark(grid, row - 1, col)
		mark(grid, row, col + 1)
		mark(grid, row, col - 1)
	}
}