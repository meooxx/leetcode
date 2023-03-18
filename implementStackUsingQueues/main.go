package main

type Node struct {
	Val  int
	Next *Node
}
type MyStack struct {
	head *Node
}

func Constructor() MyStack {
	return MyStack{&Node{}}
}

func (this *MyStack) Push(x int) {
	next := this.head.Next
	n := &Node{Val: x, Next: next}
	this.head.Next = n
}

func (this *MyStack) Pop() int {
	next := this.head.Next
	this.head.Next = this.head.Next.Next
	next.Next = nil
	return next.Val
}

func (this *MyStack) Top() int {
	return this.head.Next.Val
}

func (this *MyStack) Empty() bool {
	return this.head.Next == nil
}

/**
* Your MyStack object will be instantiated and called as such:
* obj := Constructor();
* obj.Push(x);
* param_2 := obj.Pop();
* param_3 := obj.Top();
* param_4 := obj.Empty();
 */
