package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
	fast := head
	slow := head
	// 1 2 3 -> 2
	// 1 2 3 4 -> 2
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	rightHead := slow.Next
	slow.Next = nil
	cur := rightHead
	for cur != nil && cur.Next != nil {
		next := cur.Next
		cur.Next = next.Next

		next.Next = rightHead
		rightHead = next
	}

	cur = head
	for rightHead != nil {
		next := cur.Next
		cur.Next = rightHead
		rightHead = rightHead.Next
		cur.Next.Next = next
		cur = next
	}

}
