package main

func main() {}

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}
// pre > curr
// 
func recoverTree(root *TreeNode) {
	var first *TreeNode
	var second *TreeNode
	var pre *TreeNode
	var inOrderTraverse func(root *TreeNode)

	inOrderTraverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		inOrderTraverse(root.Left)

		if pre != nil && pre.Val >= root.Val {
			if first == nil {
				first = pre
			}
			second = root
		}
		pre = root
		inOrderTraverse(root.Right)

	}
	first.Val, second.Val = second.Val, first.Val

}
func recoverTree2(root *TreeNode) {
	stack := []*TreeNode{}
	var pre *TreeNode
	var first *TreeNode
	var second *TreeNode
	node := root

	for node != nil || len(stack) != 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}

		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if pre != nil && pre.Val >= node.Val {
			if first == nil {
				first = pre
			}
			second = node
		}
		pre = node
		node = node.Right

	}
	first.Val, second.Val = second.Val, first.Val
}
