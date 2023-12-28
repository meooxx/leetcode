package main

type NumMatrix struct {
	sum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	sum := make([][]int, len(matrix)+1)


	// matrix        sum 0  0   
	//   1  2  3         0 |              |         |
	//   4  5  6         0 |              |         |
	//   7  8  9         0 |-------------sum8-------sum(9)               
	//   10 11 12        0 |-------------sum(11)-----12
	//                           状态转移    sum(12) == sum9 + sum11 -sum8 + 12
	// 												根据sum表和图示可求 sum(8-12) = sum(12) - sum9 - sum11 + sum8
	for i := range sum {
		sum[i] = make([]int, len(matrix[0])+1)
	}
	for row := 1; row < len(sum); row++ {
		for col := 1; col < len(sum[row]); col++ {
			sum[row][col] = matrix[row-1][col-1] + sum[row-1][col] + sum[row][col-1] - sum[row-1][col-1]
		}
	}
	return NumMatrix{
		sum: sum,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	area := this.sum[row2+1][col2+1] - this.sum[row2+1][col1] - this.sum[row1][col2+1] + this.sum[row1][col1]
	return area
}

/**
* Your NumMatrix object will be instantiated and called as such:
* obj := Constructor(matrix);
* param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
