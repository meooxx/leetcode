/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	result := &ListNode{Next: head}
	pre := result

	for pre != nil && pre.Next != nil {
		if pre.Next.Val == val {
			pre.Next = pre.Next.Next
		} else {
			pre = pre.Next
		}
	}
	return result.Next
}
