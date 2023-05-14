package main

import "fmt"

type MaxHeap struct {
	values []int
	index  int
}

func (this *MaxHeap) Push(n int) {
	this.values = append(this.values, n)
	this.index++
	this.Liftup(this.index)

}

func (this *MaxHeap) Peek(n int) int {
	return this.values[1]
}

func (this *MaxHeap) Pop() int {
	v := this.values[1]
	this.values[1] = this.values[this.index]
	this.values = this.values[:this.index]
	this.index--
	this.shitDown(1)
	return v
}

//	 				100(1)
//	50(2)   				60(3)
//
// 10(4) 20(5) 		56(6) 57(7)
func (this *MaxHeap) Liftup(pos int) {
	parent := pos / 2
	for parent > 0 && this.values[pos] > this.values[parent] {
		this.values[parent], this.values[pos] = this.values[pos], this.values[parent]
		pos = parent
		parent = pos / 2
	}
}
// x 4 2 1 0
func (this *MaxHeap) shitDown(pos int) {
	for pos <= this.index {
		left := pos * 2
		right := pos*2 + 1
		parent := pos
		if left <= this.index && this.values[pos] < this.values[left] {
			pos = left
		}
		if right <= this.index && this.values[pos] < this.values[right] {
			pos = right
		}
		if parent != pos {
			this.values[pos], this.values[parent] = this.values[parent], this.values[pos]
		} else {
			break
		}
	}

}

func main() {
	h := *&MaxHeap{values: []int{0}}
	h.Push(1)
	h.Push(4)
	h.Push(2)
	h.Push(5)
	fmt.Print(h.values)
	h.Pop()
	fmt.Print(h.values)
	h.Pop()
	fmt.Print(h.values)
	h.Push(3)
	fmt.Print(h.values)

}
