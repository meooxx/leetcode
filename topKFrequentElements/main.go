func topKFrequent(nums []int, k int) []int {
	mp := map[int]int{}
	for _, v := range nums {
		mp[v]++
	}
	bucket := make([][]int, len(nums)+1)
	for key, val := range mp {
		if bucket[val] == nil {
			bucket[val] = []int{}
		}
		bucket[val] = append(bucket[val], key)
	}
	result := []int{}

	// consider 1,1,1 2,2,2 3 k =1, len(result) > k ?
	// NO, CUZ' this is illegal testcase, result can be either 1 or 2
	for i := len(bucket) - 1; i >= 0 && len(result) < k; i-- {
		if bucket[i] != nil {
			result = append(result, bucket[i]...)
		}
	}
	return result
}