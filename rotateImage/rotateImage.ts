/**
 Do not return anything, modify matrix in-place instead.
 */
function rotate(matrix: number[][]): void {
	// up-side down
	let i = 0
	let j = matrix.length
	while (i < j) {
		for (let col = 0; col < matrix[0].length; col++) {
			[matrix[i][col], matrix[j][col]] = [matrix[j][col], matrix[i][col]]
		}
		i++
		j--
	}

	for (let i = 0; i < matrix.length; i++) {
		for (let j = i + 1; j < matrix[0].length; j++) {
			[matrix[i][j], matrix[j][i]] = [matrix[j][i], matrix[i][j]]

		}
	}
};


