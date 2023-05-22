package main

import "container/heap"

type IntHeap struct {
	Values   []int
	LessFunc func(int, int) bool
}
// low: maxHeap   3 2 1     3 2 1     
// high: minHeap  4 5       4 5 6
// if size is odd, push(low) , push(pop(low))
//                 median = low.Peek()
// if size is even, push(high), push(pop(hight)) 
//                 median = (low.Peek() + high.Peek()) / 2 
func (h IntHeap) Len() int           { return len(h.Values) }
func (h IntHeap) Less(i, j int) bool { return h.LessFunc(h.Values[i], h.Values[j]) }
func (h IntHeap) Swap(i, j int)      { h.Values[i], h.Values[j] = h.Values[j], h.Values[i] }
func (h IntHeap) Peek() int {
	return h.Values[0]
}

func (h *IntHeap) Pop() any {
	l := len(h.Values)
	v := h.Values[l-1]
	h.Values = h.Values[:l-1]
	return v
}
func (h *IntHeap) Push(n any) {
	h.Values = append(h.Values, n.(int))
}

type MedianFinder struct {
	MaxHeap *IntHeap
	MinHeap *IntHeap
	Size    int
}

func Constructor() MedianFinder {
	maxH := &IntHeap{
		[]int{},
		func(a, b int) bool {
			return a < b
		},
	}
	minH := &IntHeap{
		[]int{},
		func(a, b int) bool {
			return b < a
		},
	}
	return MedianFinder{
		MaxHeap: maxH,
		MinHeap: minH,
	}
}

// 2 1  						2 1
// 3 4  						3
func (this *MedianFinder) AddNum(num int) {

	if this.Size%2 == 0 {
		heap.Push(this.MaxHeap, num)
		heap.Push(this.MinHeap, heap.Pop(this.MaxHeap))
	} else {
		heap.Push(this.MinHeap, num)
		heap.Push(this.MaxHeap, heap.Pop(this.MinHeap))
	}
	this.Size++
}

func (this *MedianFinder) FindMedian() float64 {
	if this.Size%2 == 0 {
		return float64(this.MinHeap.Peek()+this.MaxHeap.Peek()) / 2
	} else {
		return float64(this.MinHeap.Peek())
	}
}

/**
* Your MedianFinder object will be instantiated and called as such:
* obj := Constructor();
* obj.AddNum(num);
* param_2 := obj.FindMedian();
 */
