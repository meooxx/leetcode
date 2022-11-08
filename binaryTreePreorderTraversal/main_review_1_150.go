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


//     1
//   4  3
// 2 5   7
// 先 left |
// 1 4 2 5   3 7

func preorderTraversal(root *TreeNode) []int {

	result := []int{}
	if root == nil {
		return result
	}
	queue := []*TreeNode{root}
	// 1 4 3 2 5 1  => 1 4 2 5 3 1
	for len(queue) > 0 {
		node := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		result = append(result, node.Val)
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
	}
	return result
}
