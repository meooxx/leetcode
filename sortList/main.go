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
	// h := sortList(head)
	// fmt.Print("sortList")
	// for cur := h; cur != nil; {
	// 	fmt.Print(cur.Val)
	// 	cur = cur.Next
	// }
	h2 := sortList1(head)
	fmt.Println("sortList1")
	for cur := h2; cur != nil; {
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

// merge sort
func sortList(head *ListNode) *ListNode {
	return sort(head)

}

// bottom line sort
func sortList1(head *ListNode) *ListNode {
	length := 0

	for cur := head; cur != nil; {
		length++
		cur = cur.Next
	}

	split := func(node *ListNode, n int) *ListNode {
		if node == nil {
			return node
		}
		for i := 1; i < n && node.Next != nil; i++ {
			node = node.Next
		}
		next := node.Next
		node.Next = nil
		return next
	}

	merge := func(left, right, pre *ListNode) *ListNode {
		for left != nil && right != nil {
			if left.Val <= right.Val {
				pre.Next = left
				left = left.Next
				pre = pre.Next
			} else {
				pre.Next = right
				pre = pre.Next
				right = right.Next
			}
		}
		if left != nil {
			pre.Next = left
		}
		if right != nil {
			pre.Next = right
		}
		for pre.Next != nil {
			pre = pre.Next
		}
		return pre
	}
	// left, right 先各取一个合并
	// 再各取俩个 合并
	// 取 4 , 取 8...
	// 
	//  l r
	//  4 3 1 5 6 2
	//  4 3                1
	///     1 5
	//          6 2
	// 34 15               2
	//          26
	// 1345     26         4
	result := &ListNode{Next: head}
	for step := 1; step < length; step <<= 1 {
		cur := result.Next
		pre := result

		for cur != nil {
			left := cur
			right := split(left, step)
			cur = split(right, step)
			pre = merge(left, right, pre)
		}
	}
	return result.Next
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
}
