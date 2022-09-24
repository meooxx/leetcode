package main

import "fmt"

func main() {
	l1 := &ListNode{Val: 0}
	l2 := &ListNode{Val: 1}
	l3 := &ListNode{Val: 1}
	l4 := &ListNode{Val: 2}
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	fmt.Printf("%#v", deleteDuplicates(l1))
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {

	curr := head
	for curr != nil && curr.Next != nil {
		if curr.Val == curr.Next.Val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}

	}
	return head
}
