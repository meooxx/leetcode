package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	minV := min(p.Val, q.Val)
	maxV := max(p.Val, q.Val)
	if maxV < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	} else if minV > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	} else {
		return root
	}
}
