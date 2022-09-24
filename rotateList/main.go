package main

import "fmt"

func main() {
	l1 := ListNode{Val: 1}
	l2 := ListNode{Val: 2}
	l3 := ListNode{Val: 3}

	l1.Next = &l2
	l2.Next = &l3
	l := rotateRight(&l1, 2)
	for l != nil {
		fmt.Println(l)
		l = l.Next
	}

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {

	count := 0
	l := head
	var tail *ListNode
	for l != nil {
		count++
		if l.Next == nil {
			tail = l
		}
		l = l.Next

	}
	if count < 2 {
		return head
	}

	l = head
	k = k % count
	if k == 0 {
		return head
	}
	for i := 1; i != count-k; i++ {
		l = l.Next
	}

	next := l.Next
	l.Next = nil
	tail.Next = head

	return next
}
