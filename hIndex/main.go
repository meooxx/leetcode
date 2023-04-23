package main
// 3 0 6 5 1
// [0 1 0 1  2]
//  0 1 0 3 (5,6)
// 65310
func hIndex(citations []int) int {
	n := len(citations)
	bucket := make([]int, n+1)
	for _, v := range citations {
		if v >= n {
			bucket[n]++
		} else {
			bucket[v]++
		}
	}

	count := 0
	for i := n; i >= 0; i-- {
		count += bucket[i]
		if count >= i {
			return i
		}
	}
	return 0
}
