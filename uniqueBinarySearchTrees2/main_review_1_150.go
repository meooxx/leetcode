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

func generateTrees(n int) []*TreeNode {
	dp := make([][]*TreeNode, n+1)
	dp[0] = []*TreeNode{nil}
	for target := 1; target <= n; target++ {
		dp[target] = []*TreeNode{}
		for root := 1; root <= target; root++ {
			left := root - 1
			right := target - root
			for _, leftNode := range dp[left] {
				for _, rightNode := range dp[right] {
					dp[target] = append(dp[target], &TreeNode{
						Val:   root,
						Left:  leftNode,
						Right: copyNode(rightNode, root),
					})
				}
			}
		}
	}
	return dp[n]
}

func copyNode(n *TreeNode, offset int) *TreeNode {
	if n == nil {
		return nil
	}

	return &TreeNode{
		Val:   n.Val + offset,
		Left:  copyNode(n.Left, offset),
		Right: copyNode(n.Right, offset),
	}

}
