package main

type NumMatrix struct {
	sum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	sum := make([][]int, len(matrix)+1)
	for i := range sum {
		sum[i] = make([]int, len(matrix[0])+1)
	}
	// if range the matrix,  it starts from row 1 and col 1 and will miss matrix[0][0].
	// if it starts row-1, col-1 and it will miss matrix[end][end]
	// so we range the sum
	for row := 1; row < len(sum); row++ {
		for col := 1; col < len(sum[row]); col++ {
			sum[row][col] = matrix[row-1][col-1] + sum[row][col-1] + sum[row-1][col] - sum[row-1][col-1]
		}
	}
	return NumMatrix{
		sum,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.sum[row2+1][col2+1] + this.sum[row1][col1] - this.sum[row2+1][col1] - this.sum[row1][col2+1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
