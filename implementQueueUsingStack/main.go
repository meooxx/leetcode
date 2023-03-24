package main

type MyQueue struct {
	Stack []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.Stack = append(this.Stack, x)
}

func (this *MyQueue) Pop() int {
	x := this.Stack[0]
	this.Stack = this.Stack[1:]
	return x
}

func (this *MyQueue) Peek() int {
	return this.Stack[0]
}

func (this *MyQueue) Empty() bool {
	return len(this.Stack) == 0
}

/**
* Your MyQueue object will be instantiated and called as such:
* obj := Constructor();
* obj.Push(x);
* param_2 := obj.Pop();
* param_3 := obj.Peek();
* param_4 := obj.Empty();
 */
