package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}

	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		cur := root
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		node := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		if k == 1 {
			return node.Val
		}
		k--
		root = node.Right

	}
	return 0
}
