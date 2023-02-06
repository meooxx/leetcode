package main

type Node struct {
	head *Node
	val  int
	min  int
	next *Node
}
type MinStack struct {
	head *Node
}

func Constructor() MinStack {
	return MinStack{}
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func (this *MinStack) Push(val int) {
	if this.head == nil {
		this.head = &Node{
			val: val,
			min: val,
		}
	} else {
		this.head = &Node{
			val:  val,
			min:  min(this.head.min, val),
			next: this.head,
		}
	}
}

func (this *MinStack) Pop() {
	if this.head == nil {
		return
	}
	this.head = this.head.next
}

func (this *MinStack) Top() int {
	return this.head.val
}

func (this *MinStack) GetMin() int {
	return this.head.min
}
