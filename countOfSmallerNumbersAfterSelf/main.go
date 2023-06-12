package main

import "fmt"

func main() {
	fmt.Println(countSmaller([]int{
		5, 2, 6, 1,
	}))
}

type ValWithIndex struct {
	val   int
	index int
}

func countSmaller(nums []int) []int {

	newNums := make([]ValWithIndex, len(nums))
	for i := range nums {
		newNums[i] = ValWithIndex{
			val:   nums[i],
			index: i,
		}
	}
	result := make([]int, len(nums))
	mergeSort(newNums, 0, len(nums)-1, result)

	anwser := make([]int, len(result))
	for i := range result {
		anwser[i] = result[i]
	}
	return anwser
}

func mergeSort(nums []ValWithIndex, start, end int, result []int) {
	if start >= end {
		return
	}

	mid := (end + start) / 2
	mergeSort(nums, start, mid, result)
	mergeSort(nums, mid+1, end, result)

	rightNum := 0
	merged := []ValWithIndex{}
	i, j := start, mid+1
	for i <= mid && j <= end {
		if nums[i].val > nums[j].val {
			rightNum++
			merged = append(merged, nums[j])
			j++
		} else {
			merged = append(merged, nums[i])
			result[nums[i].index] += rightNum
			i++

		}
	}
	for i <= mid {
		merged = append(merged, nums[i])
		result[nums[i].index] += rightNum
		i++

	}
	for j <= end {
		merged = append(merged, nums[j])
		j++

	}

	for i := 0; i < len(merged); i++ {
		nums[start+i] = merged[i]
	}

}
