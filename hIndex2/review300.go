package main

// 3  5  10  28
// 0  1  2   3
// 4	3  2	 1				 len - index
// so just find the first element that >= index
// return len(citations) - left

// 如果为倒叙找到最后一个 >= index
// 这里则找到第一个 >= index
func hIndex(citations []int) int {

	left, right := 0, len(citations)-1
	// e.g [0]
	for left <= right {
		midIndex := (right-left)/2 + left
		if citations[midIndex] > len(citations)-midIndex {
			right = midIndex - 1
		} else if citations[midIndex] < len(citations)-midIndex {
			left = midIndex + 1
		} else {
			return len(citations) - midIndex
		}

	}
	return len(citations) - left

}
