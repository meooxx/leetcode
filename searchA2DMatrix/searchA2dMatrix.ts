


function searchMatrix1(matrix: number[][], target: number): boolean {
	let i = 0
	let j = matrix.length - 1
	while (i <= j) {
		const mid = ~~((j - i) / 2) + i
		if (matrix[mid][0] == target) return true
		else if (matrix[mid][0] < target) {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	if (j < 0 || j == matrix.length) return false
	let left = 0
	let right = matrix[j].length - 1

	while (left <= right) {
		const mid = ~~((right - left) / 2) + left
		if (matrix[j][mid] == target) return true
		else if (matrix[j][mid] < target) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false

};

// apply binary search in entire matrix
//  1 2  3
//  4 5  6
//  7 8  9
function searchMatrix(matrix: number[][], target: number): boolean {
	let low = 0
	let high = matrix.length * matrix[0].length - 1

	while (low <= high) {
		const mid = ~~((high - low) / 2) + low
		const row = ~~(mid / matrix[0].length)
		const col = mid % matrix[0].length
		if (matrix[row][col] === target) return true
		else if (matrix[row][col] < target) {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false


}
