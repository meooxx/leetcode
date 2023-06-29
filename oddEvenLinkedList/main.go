package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	last := head
	cur := head.Next

	// 2 1 3 5
	// 2 3 1 5
	for cur != nil && cur.Next != nil {
		next := cur.Next
		cur.Next = next.Next
		next.Next = last.Next
		last.Next = next

		last = last.Next
		cur = cur.Next

	}

	return head
}
