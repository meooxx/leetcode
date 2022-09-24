package main

import "fmt"

func main() {
	fmt.Println(uniquePaths(3, 2))
	fmt.Println(uniquePaths(3, 7))
	fmt.Println(uniquePaths(11, 1))
}

// basic
//  fn(m,n) == f(m-1, n) + f(m, n-1)
// 1      1      	 1     1           advance1:  cur(j3) == pre[j3] + cur(j2)
// 1     2(1+1)      3     4           advance2: cur(j3) == cur[j3] + cur[j2]
// 1     3           6     10          新循环开始则 cur自动变成了pre所以 == 1 +  2
// 1

// advanced1
// pre[col] + cur[col-1]
// advance2
// cur[col] = cur[col] + cul[col-1]

func uniquePaths(m, n int) int {
	r := make([]int, m)
	for i := range r {
		r[i] = 1
	}
	for row := 1; row < n; row++ {
		for col := 1; col < m; col++ {
			r[col] = r[col] + r[col-1]
		}
	}
	return r[m-1]
}

func uniquePaths3(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}
	r := make([][]int, n)
	for index := range r {
		r[index] = make([]int, m)
		r[index][0] = 1
	}
	for i := range r[0] {
		r[0][i] = 1
	}
	for row := 1; row < len(r); row++ {
		for col := 1; col < len(r[row]); col++ {
			r[row][col] = r[row-1][col] + r[row][col-1]
		}
	}
	return r[n-1][m-1]
}

func uniquePaths2(m int, n int) int {

	count := 0
	next(1, 1, m, n, &count)
	return count
}

func next(x, y, m, n int, count *int) {
	if x == m && y == n {
		*count++
		return
	}
	if x < m {
		next(x+1, y, m, n, count)

	}
	if y < n {
		next(x, y+1, m, n, count)
	}

}
