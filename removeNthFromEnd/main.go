package main

import (
	"fmt"
)

func main() {
	last := &ListNode{
		3, nil}
	/* n:= &ListNode{
	2,last} */

	head := &ListNode{
		1,
		last,
	}

	fmt.Printf("%v", removeNthFromEnd(head, 1))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//              index from end
// 1->2->3      2   1-3
// 1 -> 2 ->3   3,  2-3
// 1            1   nil
func removeNthFromEnd(head *ListNode, n int) *ListNode {

	var first *ListNode = head
	size := 1
	for first.Next != nil {
		first = first.Next
		size++
	}
	first = head
	last := &ListNode{Next: head}
	result := last
	for i := 0; i < size; i++ {
		if i == size-n {
			next := first.Next
			last.Next = next
		} else {
			last = first
			first = first.Next
		}

	}
	return result.Next
}
