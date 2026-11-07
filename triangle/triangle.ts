function minimumTotal(triangle: number[][]): number {
	if (triangle.length == 1) return triangle[0][0]
	for (let bottom = triangle.length - 2; bottom >= 0; bottom--) {
		for (let i = 0; i < triangle[bottom].length; i++) {
			triangle[bottom][i] = triangle[bottom][i] + Math.min(triangle[bottom + 1][i], triangle[bottom + 1][i + 1])
		}
	}
	return triangle[0][0]
};