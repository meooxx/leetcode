package main

import "container/heap"

type IntHeap struct {
	Vals     []int
	LessFunc func(a, b int) bool
}

func (this IntHeap) Len() int {
	return len(this.Vals)
}

func (this *IntHeap) Swap(a, b int) {
	this.Vals[a], this.Vals[b] = this.Vals[b], this.Vals[a]
}
func (this *IntHeap) Push(a any) {
	this.Vals = append(this.Vals, a.(int))
}
func (this *IntHeap) Pop() any {
	v := this.Vals[len(this.Vals)-1]
	this.Vals = this.Vals[:len(this.Vals)-1]
	return v
}
func (this *IntHeap) Less(a, b int) bool {
	return this.LessFunc(this.Vals[a], this.Vals[b])

}
func (this *IntHeap) Peek() int {
	return this.Vals[0]
}

type MedianFinder struct {
	Size  int
	Large *IntHeap
	Small *IntHeap
}

func Constructor() MedianFinder {
	large := &IntHeap{
		Vals: []int{},
		LessFunc: func(a, b int) bool {
			return a < b
		},
	}
	small := &IntHeap{
		Vals: []int{},
		LessFunc: func(a, b int) bool {
			return a > b
		},
	}
	heap.Init(small)
	heap.Init(large)
	return MedianFinder{
		Size:  0,
		Large: large, // minHeap
		Small: small, // maxHeap
	}
}

func (this *MedianFinder) AddNum(num int) {
	if this.Size%2 == 0 {
		heap.Push(this.Small, num)
		heap.Push(this.Large, heap.Pop(this.Small))
	} else {
		heap.Push(this.Large, num)
		heap.Push(this.Small, heap.Pop(this.Large))
	}
	this.Size++
}

func (this *MedianFinder) FindMedian() float64 {
	
	if this.Size%2 == 0 {
		return float64(this.Large.Peek()+this.Small.Peek()) / 2
	} else {
		return float64(this.Large.Peek())
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
