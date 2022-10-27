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

func sumNumbers(root *TreeNode) int {
	//   1
	//  2  3
	var getSum func(root *TreeNode, sum int) int
	getSum = func(n *TreeNode, sum int) int {
		if n == nil {
			return 0
		}
		sum = sum*10 + n.Val
		if n.Left == nil && n.Right == nil {
			return sum
		}
		return getSum(n.Left, sum) + getSum(n.Right, sum)
	}
	return getSum(root, 0)
}
