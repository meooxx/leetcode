package main

import "container/heap"

//  a       b      c        i
//  d       e      f        j
// ad ,     bd,    cd
//  |      |      |
// push(ae) p(be)  p(ce, cf)
//  |       |
// push(af) p(bf)
// at first push all (ad, bd, cd,) nums1[i], nums2[0]
// pop from MinHeap got minium val as cur, e.g. ad
// push cur[2](j index +1) == nums1[i], nums2[j+1], so on be ce, af, bf, cf

type PriorityQueue [][]int

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(a, b int) bool {
	return pq[a][0]+pq[a][1] < (pq[b][0] + pq[b][1])
}

func (pq PriorityQueue) Swap(i, j int) {

	pq[i], pq[j] = pq[j], pq[i]

}

func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.([]int))
}
func (pq *PriorityQueue) Pop() any {
	first := (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]
	return first
}
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	pq := &PriorityQueue{}
	heap.Init(pq)
	for i := range nums1 {
		heap.Push(pq, []int{nums1[i], nums2[0], 0})
	}
	result := [][]int{}
	for k > 0 && pq.Len() > 0 {
		cur := heap.Pop(pq).([]int)
		result = append(result, []int{cur[0], cur[1]})
		k--
		if cur[2] < len(nums2)-1 {
			heap.Push(pq, []int{cur[0], nums2[cur[2]+1], cur[2] + 1})
		}
	}

	return result
}
