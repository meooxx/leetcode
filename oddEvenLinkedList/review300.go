package main

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

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	cur := head.Next
	last := head
	for cur != nil && cur.Next != nil {
		next := cur.Next
		nextNext := next.Next
		// last  cur next  nextNext
		// 2      1  3     5
		cur.Next = nextNext
		next.Next = last.Next
		last.Next = next

		last = last.Next
		cur = cur.Next
	}
	return head
}
