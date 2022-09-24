/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	fmt.Println(buildTree2([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}))
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	rootIndex := 0
	for index, v := range inorder {
		if v == root.Val {
			rootIndex = index
		}
	}
	left := buildTree(preorder[1:rootIndex+1], inorder[:rootIndex])
	right := buildTree(preorder[rootIndex+1:], inorder[rootIndex+1:])
	root.Left = left
	root.Right = right
	return root
}

func buildTree2(preorder []int, inorder []int) *TreeNode {
	mp := map[int]int{}
	for index, v := range inorder {
		mp[v] = index
	}
	var helper func(preStart, preEnd, inStart, inEnd int) *TreeNode
	helper = func(preStart, preEnd, inStart, inEnd int) *TreeNode {
		if preStart > preEnd {
			return nil
		}
		root := &TreeNode{Val: preorder[preStart]}
		rootIndex := mp[preorder[preStart]]
		left := helper(preStart+1, preStart+rootIndex-inStart, inStart, rootIndex-1)
		right := helper(preStart+rootIndex-inStart+1, preEnd, rootIndex+1, inEnd)
		root.Left = left
		root.Right = right
		return root
	}
	return helper(0, len(preorder)-1, 0, len(inorder)-1)
}