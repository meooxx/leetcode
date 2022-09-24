package main

import (
	"fmt"
)

func main() {
	l1 := &ListNode{Val: 1}
	l2 := &ListNode{Val: 2}
	l3 := &ListNode{Val: 3}
	l4 := &ListNode{Val: 4}

	l1.Next = l2
	l2.Next = l3
	l3.Next = l4

	list := swapPairs(l1)

	for list != nil {
		fmt.Println(list)
		list = list.Next
	}

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	result := &ListNode{}
	last := result
	current := head
	next := current.Next
	nextNext := next.Next
	// 1 2 3 4
	for {
		last.Next = next
		next.Next = current
		current.Next = nextNext

		if nextNext == nil || nextNext.Next == nil {
			break
		}
		// next current nextNext
		// 2     1      3         4  ,..
		last = current
		current = nextNext
		next = nextNext.Next
		nextNext = next.Next
	}

	return result.Next
}
