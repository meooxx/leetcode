package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rob(root *TreeNode) int {
	a, b := robImpl(root)
	return max(a, b)
}

func robImpl(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	l0, l1 := robImpl(root.Left)
	r0, r1 := robImpl(root.Right)

	val1 := root.Val + root.Left.Val + root.Right.Val
	val2 := max(l0, l1) + max(r0, r1)
	return val1, val2
}
