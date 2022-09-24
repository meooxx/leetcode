package main

func main() {}
// 1 先左右俩边放1
// 2 随后按位计算
// [1331]
// [14641]
func generate(numRows int) [][]int {
	result := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		result[i] = make([]int, i+1)
		result[i][0] = 1
		result[i][i] = 1
		for j := 1; j < i; j++ {
			result[i][j] = result[i-1][j-1] + result[i-1][j]
		}
	}
	return result
}
func generate2(numRows int) [][]int {
	result := [][]int{{1}}
	if numRows <= 1 {
		return result
	}
	for i := 2; i <= numRows; i++ {
		pre := result[len(result)-1]
		item := make([]int, len(pre)+1)
		item[0] = 1
		item[len(item)-1] = 1

		for j := 1; j <= len(item)-2; j++ {
			item[j] = pre[j-1] + pre[j]
		}
		result = append(result, item)

	}
	return result
}
