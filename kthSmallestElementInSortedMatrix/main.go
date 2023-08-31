package main


func kthSmallest(matrix [][]int, k int) int {

	rows := len(matrix) - 1
	cols := len(matrix[0]) - 1

	left := matrix[0][0]
	right := matrix[rows][cols]
	anwser := 0
	for left <= right {
		mid := (left + right) / 2
		// 1
		// find the mid as small as possible
		// k = 3
		// 1 6 7 7 15, mid == 8, but the anwser is 7
		// 2 
		// why >= k, 
		// 1 8 8 8 15, k = 2, count(m, 8) == 4, we must save the 8
		// because the next loop mid == 1, and count(m, 1) ==  1
		if count(matrix, mid) >= k {
			anwser = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return anwser
}

func count(matrix [][]int, mid int) int {
	count := 0
	for r := 0; r < len(matrix); r++ {
		c := len(matrix[r]) - 1
		for c >= 0 && matrix[r][c] > mid {
			c--
		}
		count += c + 1
	}
	return count

}