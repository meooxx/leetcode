package main

func main() {}

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
type ListNode struct {
	Val  int
	Next *ListNode
}
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 1 2 3     2
// 1 2 3 4   2
func sortedListToBST1(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return &TreeNode{Val: head.Val}
	}
	slow := head
	fast := head
	pre := head
	for fast != nil && fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	pre.Next = nil
	root := &TreeNode{
		Val:   slow.Val,
		Left:  sortedListToBST(head),
		Right: sortedListToBST(slow.Next),
	}
	return root
}

// not change origin head
func sortedListToBST(head *ListNode) *TreeNode {
	var toBst func(head, end *ListNode) *TreeNode
	toBst = func(head, end *ListNode) *TreeNode {
		// 1 3 2(end)
		// root == 3
		// .Left = toBts(1,3)
		// .Right = toBts(3.Next == 2,2)
		if head == end {
			return nil
		}
		if head.Next == end {
			return &TreeNode{Val: head.Val}
		}
		slow := head
		fast := head
		for fast != end && fast.Next != end {
			slow = slow.Next
			fast = fast.Next.Next
		}
		root := &TreeNode{
			Val:   slow.Val,
			Left:  toBst(head, slow),
			Right: toBst(slow.Next, end),
		}
		return root
	}
	return toBst(head, nil)

}
