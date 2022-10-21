package main

func main() {}

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

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	mp := map[*TreeNode]int{}

	var getDepth func(node *TreeNode) int
	getDepth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		if v, ok := mp[node]; ok {
			return v
		}
		left := getDepth(node.Left)
		right := getDepth(node.Right)
		if left >= right {
			mp[node] = 1 + left
			return 1 + left
		}
		mp[node] = 1 + right
		return 1 + right
	}
	left := getDepth(root.Left)
	right := getDepth(root.Right)
	if left == right || left+1 == right || left-1 == right {
		return isBalanced(root.Left) && isBalanced(root.Right)
	}
	return false

}
