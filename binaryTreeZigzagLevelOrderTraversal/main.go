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
func zigzagLevelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}
	stack := []*TreeNode{root}

	reverse := false
	for len(stack) > 0 {
		size := len(stack)
		nodes := []int{}
		for i := 0; i < size; i++ {
			n := stack[0]
			stack = stack[1:]
			nodes = append(nodes, n.Val)
			if n.Left != nil {
				stack = append(stack, n.Left)
			}
			if n.Right != nil {
				stack = append(stack, n.Right)
			}
		}
		if reverse {
			for left, right := 0, len(nodes)-1; left < right; {
				nodes[left], nodes[right] = nodes[right], nodes[left]
				left++
				right--
			}
		}
		reverse = !reverse
		result = append(result, nodes)
	}
	return result
}
