package main

import "fmt"

func main() {
	fmt.Println(insert([][]int{
		{1, 3},
		{4, 7},
		{10, 18},
	}, []int{9, 10}))
	fmt.Println(insert2([][]int{
		//{1, 3},
		{6, 7},
		{10, 18},
	}, []int{8, 9}))
}

func insert(intervals [][]int, newInterval []int) [][]int {
	result := [][]int{}
	i := 0
	for i < len(intervals) {
		if intervals[i][1] < newInterval[0] {
			result = append(result, intervals[i])
			i++
		} else {
			break
		}
	}
	for i < len(intervals) && intervals[i][0] <= newInterval[1] {

		if intervals[i][0] < newInterval[0] {
			newInterval[0] = intervals[i][0]
		}
		if intervals[i][1] > newInterval[1] {
			newInterval[1] = intervals[i][1]
		}
		i++

	}
	result = append(result, newInterval)
	if i < len(intervals) {
		result = append(result, intervals[i:]...)
	}
	return result
}
func insert2(intervals [][]int, newInterval []int) [][]int {
	

}
