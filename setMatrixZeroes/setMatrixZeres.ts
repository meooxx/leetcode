/**
 Do not return anything, modify matrix in-place instead.
 */

// if there is zero in any col or any rows
// set first  row and col  zero
// and another flag for any zero in first col
// and matrix[0][0] for any zero in first row
function setZeroes(matrix: number[][]): void {
	let col0 = false

	for (let i = 0; i < matrix.length; i++) {
		if (matrix[i][0] === 0) col0 = true
		for (let j = 1; j < matrix[0].length; j++) {
			if (matrix[i][j] === 0) {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for (let row = matrix.length - 1; row >= 0; row--) {
		for (let col = matrix[0].length - 1; col >= 1; col--) {
			if (matrix[row][0] === 0 || matrix[0][col] === 0) {
				matrix[row][col] = 0
			}
		}
		if (col0) {
			matrix[row][0] = 0
		}

	}


};