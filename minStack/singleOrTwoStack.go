package main

import "fmt"

func main() {
	s := Constructor()
	s.Push(5)
	fmt.Println(s)
	s.Push(1)
	fmt.Println(s)
	fmt.Println(s.GetMin())
	fmt.Println(s.Top())
}

type MinStack struct {
	stack    []int
	// 或者可以用一个二元组来储存数据如 [val, min]
	minStack []int
}

func Constructor() MinStack {
	return MinStack{}
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	if this.minStack == nil {
		this.minStack = append(this.minStack, val)
	} else {
		this.minStack = append(this.minStack, min(val, this.minStack[len(this.minStack)-1]))
	}
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
