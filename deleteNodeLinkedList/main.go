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
// 1 -> 2 -> 3
// 2 -> 3
func deleteNode(node *ListNode) {
	pre := node
	for node.Next != nil {
		pre = node
		node.Val = node.Next.Val
		node = node.Next

	}
	pre.Next = nil
}
