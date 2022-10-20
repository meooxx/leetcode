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
func buildTree(preorder []int, inorder []int) *TreeNode {
	mp := map[int]int{}
	for index, v := range inorder {
		mp[v] = index
	}

	// 1 2 3 preorder
	// 2 1 3 
	var constructTree func(preorder []int, inorder []int, pStart, pEnd, inStart, inEnd int) *TreeNode
	constructTree = func(preorder, inorder []int, pStart, pEnd, inStart, inEnd int) *TreeNode {
		if pStart > pEnd || inStart > inEnd {
			return nil
		}
		rootVal := preorder[pStart]
		rootIndex := mp[rootVal]

		diff := rootIndex - inStart

		root := &TreeNode{
			Val:   rootVal,
			Left:  constructTree(preorder, inorder, pStart+1, pStart+diff, inStart, rootIndex-1),
			Right: constructTree(preorder, inorder, pStart+diff+1, pEnd, rootIndex+1, inEnd),
		}
		return root

	}

	return constructTree(preorder, inorder, 0, len(preorder)-1, 0, len(preorder))
}
