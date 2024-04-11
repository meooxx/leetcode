package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	stack := []*TreeNode{root}
	result := []int{}
	if root == nil {
		return result
	}
	for len(stack) > 0 {
		result = append(result, stack[len(stack)-1].Val)
		size := len(stack)
		for i := 0; i < size; i++ {
			node := stack[0]
			stack = stack[1:]
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
		}
	}
	return result
}
