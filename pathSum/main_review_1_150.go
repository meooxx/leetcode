/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {

	if root == nil {
		return false
	}
	rest := targetSum - root.Val
	if root.Left == nil && root.Right == nil {
		return rest == 0
	}

	return hasPathSum(root.Left, rest) || hasPathSum(root.Right, rest)
}
