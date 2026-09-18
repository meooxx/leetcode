
// each block includes 9 digits
// block00  01      02
//  .       .        .  
//  .(10)   .(11)    .(12)  
//  .       .        .  
//  
function isValidSudoku(board: string[][]): boolean {
	const rowSeen = new Set<string>()
	const colSeen = new Set<string>()
	const blockSeen = new Set<string>()
	for (let r = 0; r < 9; r++) {
		for (let c = 0; c < 9; c++) {
			if (board[r][c] === '.') continue
			const rowChar = `${board[r][c]}-${r}-row`
			const colChar = `${board[r][c]}-${c}-col`
			const blockChar = `${board[r][c]}-${~~(r / 3)}/${~~(c / 3)}-block`
			if (rowSeen.has(rowChar) || colSeen.has(colChar) || blockSeen.has(blockChar)) {
				return false
			}
			rowSeen.add(rowChar)
			colSeen.add(colChar)
			blockSeen.add(blockChar)

		}

	}


	return true
}

isValidSudoku([
	["5", "3", ".", ".", "7", ".", ".", ".", "."],
	["6", ".", ".", "1", "9", "5", ".", ".", "."],
	[".", "9", "8", ".", ".", ".", ".", "6", "."],
	["8", ".", ".", ".", "6", ".", ".", ".", "3"],
	["4", ".", ".", "8", ".", "3", ".", ".", "1"],
	["7", ".", ".", ".", "2", ".", ".", ".", "6"]
	, [".", "6", ".", ".", ".", ".", "2", "8", "."]
	, [".", ".", ".", "4", "1", "9", ".", ".", "5"]
	, [".", ".", ".", ".", "8", ".", ".", "7", "9"]])