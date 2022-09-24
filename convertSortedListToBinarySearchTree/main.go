package main

import "fmt"

func main() {
	fmt.Println()
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	return toBst(head, nil)
}
// -10, -3, 0 , 3, 5
func toBst(head, tail *ListNode) *TreeNode {
	
	if head == tail {
		return nil
	}
	if head.Next == tail {
		return &TreeNode{Val:head.Val}
	}
	fast := head
	slow := head
	for fast != tail && fast.Next != tail {
		fast = fast.Next.Next
		slow = slow.Next
	}

	n := &TreeNode{Val: slow.Val}
	n.Left = toBst(head, slow)
	n.Right = toBst(slow.Next, tail) 
	return n

}
