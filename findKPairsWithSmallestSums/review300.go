package main
import "container/heap"
//  a       b      c
//  d       e      f
// ad ,     bd,    cd
//  | 1     |      |
// push(ae) p(be)  p(ce, cf)
//  |       |
// push(af) p(bf)

type PriorityQueue [][]int

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(a, b int) bool {
	return pq[a][0]+pq[a][1]-(pq[b][0]+pq[b][1]) < 0
}

func (pq PriorityQueue) Swap(i, j int) {

	pq[i], pq[j] = pq[j], pq[i]

}

func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.([]int))
}
func (pq *PriorityQueue) Pop() any {
	last := (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]
	return last
}
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	pq := &PriorityQueue{}
	heap.Init(pq)
	for i := range nums1 {
		heap.Push(pq, []int{nums1[i], nums2[0], 0})
	}
	result := [][]int{}
	for k > 0 && len(*pq) > 0 {
		minv := heap.Pop(pq).([]int)
		result = append(result, []int{minv[0], minv[1]})
		k--
		index := minv[2] + 1
		if index < len(nums2) {
			heap.Push(pq, []int{minv[0], nums2[index], index})
		}
	}
	return result
}
