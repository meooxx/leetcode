package main

import "fmt"

func main() {
	r1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	r2 := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 9,
			Left: &TreeNode{
				Val: 5,
			},
			Right: &TreeNode{
				Val: 1,
			},
		},
		Right: &TreeNode{
			Val: 0,
		},
	}

	fmt.Println(sumNumbers(r1))
	fmt.Println(sumNumbers(r2))

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {

	var getChildSum func(root *TreeNode, cur int) int
	getChildSum = func(root *TreeNode, cur int) int {
		if root == nil {
			return 0
		}
		if root.Left == nil && root.Right == nil {
			return cur*10 + root.Val
		}
		return getChildSum(root.Left, cur*10+root.Val) + getChildSum(root.Right, cur*10+root.Val)
	}

	return getChildSum(root, 0)
}

func sumNumbers1(root *TreeNode) int {
	sum := 0
	getChild(root, &sum, 0)
	return sum
}

func getChild(root *TreeNode, sum *int, cur int) {
	if root.Left == nil && root.Right == nil {
		*sum += cur*10 + root.Val
		return
	}
	t := cur*10 + root.Val
	if root.Left != nil {
		getChild(root.Left, sum, t)
	}
	if root.Right != nil {
		getChild(root.Right, sum, t)
	}
}
