package main

import "fmt"

func main() {
	nums := []int{1,-3,6,8, 9}
	fmt.Println(countRangeSum(nums, 2, 10))
}

func countRangeSum(nums []int, lower int, upper int) int {
	sums := make([]int, len(nums)+1)
	for i := range nums {
		sums[i+1] = sums[i] + nums[i]
	}
	return mergeSortCount(sums, 0, len(sums)-1, lower, upper)
}
func mergeSortCount(sums []int, left, right int, lower, upper int) int {
	if left >= right {
		return 0
	}
	mid := (right-left)/2 + left
	leftCount := mergeSortCount(sums, left, mid, lower, upper)
	rightCount := mergeSortCount(sums, mid+1, right, lower, upper)
	count := leftCount + rightCount
	// valid index start
	startIndex := mid + 1
	// vaid index end
	endIndex := mid + 1
	// val of index that is smaller that sums[i]
	minIndex := mid + 1
	// cache index
	index := 0
	cache := make([]int, right-left+1)
	for i := left; i <= mid; i++ {
		for startIndex <= right && sums[startIndex]-sums[i] < lower {
			startIndex++
		}
		for endIndex <= right && sums[endIndex]-sums[i] <= upper {
			endIndex++
		}
		for minIndex <= right && sums[minIndex] < sums[i] {
			cache[index] = sums[minIndex]
			minIndex++
			index++
		}
		cache[index] = sums[i]
		index++
		count += endIndex - startIndex
	}
	for minIndex <= right {
		cache[index] = sums[minIndex]
		minIndex++
		index++
	}
	for i := 0; i <= len(cache)-1; i++ {
		sums[left+i] = cache[i]
	}
	return count
}
