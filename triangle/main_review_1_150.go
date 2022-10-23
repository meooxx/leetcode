package main


func minimumTotal(triangle [][]int) int {
	min := func(a, b int) int {
		if a <= b {
			return a
		}
		return b
	}
	for row := len(triangle) - 2; row >= 0; row-- {
		for i := 0; i < len(triangle[row]); i++ {
			triangle[row][i] += min(triangle[row+1][i], triangle[row][i+1])
		}
	}
	return triangle[0][0]
}
