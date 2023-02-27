package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	result := &ListNode{}
	result.Next = head
	pre := result
	cur := pre.Next
	for cur != nil {
		if cur.Val == val {
			pre.Next = cur.Next
			cur = cur.Next
		} else {
			pre = pre.Next
			cur = pre.Next
		}

	}
	return result.Next
}
