package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	cur := head

	// a b c
	for cur != nil && cur.Next != nil {
		next := cur.Next
		nextNext := next.Next
		cur.Next = nextNext
		next.Next = head
		head = next
	}
	return head
}
