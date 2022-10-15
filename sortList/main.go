package main

import "fmt"

func main() {

	head := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val: 6,
						},
					},
				},
			},
		},
	}
	head = sortList(head)
	for cur := head; cur != nil; {
		fmt.Print(cur.Val)
		cur = cur.Next
	}
}

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

func sortList(head *ListNode) *ListNode {
	return sort(head)

}

func mergeSort(left, right *ListNode) *ListNode {
	result := &ListNode{}
	cur := result
	for left != nil && right != nil {
		if left.Val <= right.Val {
			cur.Next = left
			cur = cur.Next
			left = left.Next
		} else {
			cur.Next = right
			cur = cur.Next
			right = right.Next
		}

	}
	if left != nil {
		cur.Next = left
	}
	if right != nil {
		cur.Next = right
	}
	return result.Next
}

func sort(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	slow := head
	fast := head

	// var switchNode func()
	// switchNode := func() {}

	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	rightHead := slow.Next
	slow.Next = nil
	left := sort(head)
	right := sort(rightHead)
	list := mergeSort(left, right)
	return list
	// mid
	// 1 2 3    2
	// 1 2 3 4  2
	// eg. 4 1 2 3 5 6

}
