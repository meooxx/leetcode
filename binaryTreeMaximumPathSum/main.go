package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxPathSum(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}))
}

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

//    -1              2
//  2   3          4   5
// 4   2  5           3  6
// max 2+3+5      2+  5 + max(3,6),
// case1          case2
// max( 1 , sum(left)+1, 1+sum(right), 1 + left + right )
// case2 因为从上到下一级, 只能选择做或者右
//
func maxPathSum(root *TreeNode) int {
	var getPathSum func(n *TreeNode) int
	maxSum := math.MinInt32
	max := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}
	getPathSum = func(n *TreeNode) int {
		if n == nil {
			return 0
		}
		leftSum := getPathSum(n.Left)
		rightSum := getPathSum(n.Right)
		maxLeftRight := max(leftSum, rightSum)
		sum := n.Val
		sum = max(sum, leftSum+n.Val)
		sum = max(sum, rightSum+n.Val)
		sum = max(sum, leftSum+rightSum+n.Val)
		if sum > maxSum {
			maxSum = sum
		}
		return max(n.Val, maxLeftRight+n.Val)

	}
	getPathSum(root)
	return maxSum
}
