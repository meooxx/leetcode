/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

package main

import "fmt"

func main() {
	h := &TreeNode{
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
		Right: &TreeNode{Val: 5, Right: &TreeNode{
			Val: 6,
		}},
	}
	flatten(h)
	fmt.Println(h)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten1(root *TreeNode) {
	if root == nil {
		return
	}
	if root.Left == nil {
		flatten(root.Right)
		return
	}
	right := root.Right
	root.Right = root.Left
	cur := root.Left

	for cur != nil && cur.Right != nil {

		cur = cur.Right
	}
	cur.Right = right
	root.Left = nil
	flatten(root.Right)
}

func flatten(root *TreeNode) {
	for root != nil {
		cur := root.Left
		if cur != nil {
			for cur != nil && cur.Right != nil {
				cur = cur.Right
			}
			cur.Right = root.Right
			root.Right = root.Left
			root.Left = nil
		}
		root = root.Right
	}
}
