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
func levelOrderBottom(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		nodes := []int{}
		for range stack {
			node := stack[0]
			stack = stack[1:]
			nodes = append(nodes, node.Val)
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
		}
		result = append(result, nodes)
	}
	for left, right := 0, len(result)-1; left < right; {
		result[left], result[right] = result[right], result[left]
		left++
		right--

	}
	return result
}
