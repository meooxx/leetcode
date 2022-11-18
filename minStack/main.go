package main

type MinStack struct {
	head *Node
}

type Node struct {
	Val  int
	Min  int
	Next *Node
}
//  head = {val, minVal, next}
// push =>  new Node{val, minVal(val, head.Min), head}
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
	if this.head == nil {
		this.head = &Node{val, val, nil}
	} else {
		this.head = &Node{val, min(val, this.head.Min), this.head}
	}
}

func (this *MinStack) Pop() {
	this.head = this.head.Next
}

func (this *MinStack) Top() int {
	return this.head.Val
}

func (this *MinStack) GetMin() int {
	return this.head.Min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
