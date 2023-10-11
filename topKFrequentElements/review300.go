package main

func topKFrequent(nums []int, k int) []int {
	mp := map[int]int{}

	for _, v := range nums {
		mp[v]++
	}

	bucket := make([][]int, len(nums)+1)

	for v, count := range mp {
		if bucket[count] == nil {
			bucket[count] = []int{}
		}
		bucket[count] = append(bucket[count], v)
	}
	result := []int{}
	for i := len(bucket) - 1; i >= 0 && len(result) < k; i-- {
		if bucket[i] != nil {
			result = append(result, bucket[i]...)
		}
	}
	return result
}
