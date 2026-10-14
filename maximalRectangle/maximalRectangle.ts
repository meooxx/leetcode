
function getMaxArea(arr: number[]): number {

	let maxArea = 0
	const leftLess = [-1]
	const rightLess = []
	rightLess[arr.length - 1] = arr.length

	for (let index = 1; index < arr.length; index++) {
		let p = index - 1
		while (p >= 0 && arr[p] >= arr[index]) {
			p = leftLess[p]
		}
		leftLess[index] = p
	}
	for (let index = arr.length - 2; index >= 0; index--) {
		let p = index + 1
		while (p < arr.length && arr[p] >= arr[index]) {
			p = rightLess[p]
		}
		rightLess[index] = p
	}
	for (let index = 0; index < arr.length; index++) {
		maxArea = Math.max(maxArea, (-leftLess[index] + rightLess[index] - 1) * arr[index])
	}
	return maxArea


}

//  0  1  0  0   cal  each (0, 1, 0, 0)  row cal max area in histogram
//  0  1  1  0             (0, 2, 1, 0)  accumlative     
//  1  1  1  0             (0 ,3, 2, 0)


function maximalRectangle(matrix: string[][]): number {

	const mp: number[] = Array(matrix[0].length).fill(0)
	let maxArea = 0
	for (let row = 0; row < matrix.length; row++) {

		for (let col = 0; col < matrix[0].length; col++) {
			mp[col] = matrix[row][col] === '1' ? mp[col] + 1 : 0
		}
		maxArea = Math.max(getMaxArea(mp), maxArea)

	}
	return maxArea


};

maximalRectangle([["1", "0", "1", "0", "0"], ["1", "0", "1", "1", "1"], ["1", "1", "1", "1", "1"], ["1", "0", "0", "1", "0"]])