package main

func hIndex1(citations []int) int {
	index := 0
	for i := len(citations) - 1; i >= 0; i-- {
		if citations[i] < len(citations)-i {
			return index
		}
		index++
	}
	return index
}

func hIndex(citations []int) int {
	l := len(citations)
	left := 0
	right := l - 1
	// 1 2 3
	for left <= right {
		middle := (right - left/2) + left
		if citations[middle] < l-middle {
			left = middle + 1
		} else if citations[middle] > l-middle {
			right = middle - 1
		} else {
			return l - middle
		}
	}
	return l - left

}
