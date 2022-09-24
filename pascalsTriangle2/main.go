package main

func main() {

}
func getRow(rowIndex int) []int {
	result := []int{1}
	rIndex := 1
	// 121
	//  121
	for ; rIndex <= rowIndex; rIndex++ {
		// 从后往前
		// 因为后面的值依赖前面的值, 从前往后会修改前面的值
		for index := len(result) - 1; index >= 1; index-- {
			result[index] += result[index-1]
		}
		result = append(result, 1)

	}
	return result
}
