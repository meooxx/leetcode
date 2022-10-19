package main

func main() {}

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
func levelOrder(root *TreeNode) [][]int {

	result := [][]int{}
	if root == nil {
		return result
	}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		size := len(stack)
		levelNodes := []int{}
		for i := 0; i < size; i++ {
			node := stack[0]
			stack = stack[1:]
			levelNodes = append(levelNodes, node.Val)
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
		}
		result = append(result, levelNodes)
	}
	return result
}
