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
//     -1
//   -2   9   1 最大是8不是9因为 9是一个节点, 只能是边 -1到9即 8
//  2  3      2 如果最大出现在 3字节点下如 (3 10, 11), 只能修改使用外部变量保存
//   *10 *11       函数返回值不一定是最大值,函数返回值是顶点到路径最大值
//  *意味假设有这个值
func maxPathSum(root *TreeNode) int {
	max := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}
	maxSum := -10000
	var getPath func(n *TreeNode) int
	getPath = func(n *TreeNode) int {
		if n == nil {
			return 0
		}
		left := getPath(n.Left)
		right := getPath(n.Right)
		leftRight := left + right + n.Val
		maxLeftRight := max(left, right)
		sum := n.Val
		sum = max(sum, maxLeftRight+n.Val)
		sum = max(sum, leftRight)
		if sum > maxSum {
			maxSum = sum
		}

		return max(n.Val, maxLeftRight+n.Val)

	}
	getPath(root)
	return maxSum
}
