package main

import "container/heap"

type MinHeap [][]int

func (this MinHeap) Len() int           { return len(this) }
func (this MinHeap) Less(a, b int) bool { return this[a][0] < this[b][0] }
func (this MinHeap) Swap(a, b int) {
	this[a], this[b] = this[b], this[a]
}

func (this *MinHeap) Pop() any {
	l := len(*this) - 1
	val := (*this)[l]
	*this = (*this)[:l]
	return val
}

func (this *MinHeap) Push(v any) {
	*this = append(*this, v.([]int))
}

func kthSmallest(matrix [][]int, k int) int {
	minHeap := &MinHeap{}
	heap.Init(minHeap)
	for i := 0; i < len(matrix); i++ {
		heap.Push(minHeap, []int{matrix[i][0], i, 0})
	}
	anwser := 0
	for ; k > 0; k-- {
		v := heap.Pop(minHeap).([]int)
		anwser = v[0]
		if v[2]+1 < len(matrix[v[1]]) {
			heap.Push(minHeap, []int{matrix[v[1]][v[2]+1], v[1], v[2] + 1})
		}

	}
	return anwser
}


