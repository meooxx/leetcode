package main

import "fmt"

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 7,
			},
			Right: &TreeNode{
				Val: 8,
			},
		},
	}
	fmt.Println(inorderTraversal(root))
	fmt.Println(inorderTraversal2(root))
}

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
func inorderTraversal(root *TreeNode) []int {
	result := []int{}

	var traverse func(r *TreeNode)
	traverse = func(r *TreeNode) {
		if r == nil {
			return
		}
		traverse(r.Left)
		result = append(result, r.Val)
		traverse(r.Right)

	}

	traverse(root)

	return result

}

func inorderTraversal2(root *TreeNode) []int {
	result := []int{}
	for root != nil {
		if root.Left == nil {
			result = append(result, root.Val)
			root = root.Right
		} else {
			pre := root.Left
			for pre.Right != nil && pre.Right != root {
				pre = pre.Right
			}
			// 第一次设置root
			if pre.Right == nil {
				pre.Right = root
				root = root.Left
			}
			// 第二次说明是从左节点返回
			// 保存val, 遍历右侧节点
			if pre.Right == root {
				result = append(result, root.Val)
				root = root.Right
				pre.Right = nil
			}
		}
	}

	return result

}
