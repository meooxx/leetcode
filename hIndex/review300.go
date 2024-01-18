package main


// hindex means total of citations > index
// 3  0  6  5   1
// 1 	1  0	1	  0  2 	 bucket 
//         [count>=3] got hindex
// 0  1  2  3   4  5    index
// +1   +1 +1


func hIndex(citations []int) int {
	bucket := make([]int, len(citations)+1)

	for i := range citations {
		if citations[i] > len(citations) {
			bucket[len(citations)]++
		} else {
			bucket[citations[i]]++
		}
	}
	count := 0

	for i := range bucket {
		count += bucket[i]
		if count >= i {
			return i
		}
	}
	return 0

}
