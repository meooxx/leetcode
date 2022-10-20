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
func buildTree(inorder []int, postorder []int) *TreeNode {
	var constructTree func(inStart, inEnd, posStart, posEnd int) *TreeNode
	mp := map[int]int{}
	for index, v := range inorder {
		mp[v] = index
	}
	constructTree = func(inStart, inEnd, posStart, posEnd int) *TreeNode {
		if inStart > inEnd {
			return nil
		}
		rootVal := postorder[posEnd]
		rootIndex := mp[rootVal]
		n := rootIndex - inStart - 1
		root := &TreeNode{Val: rootVal}
		root.Left = constructTree(inStart, rootIndex-1, posStart, posStart+n)
		root.Right = constructTree(rootIndex+1, inEnd, posStart+n+1, posEnd-1)
		return root
	}
	return constructTree(0, len(inorder)-1, 0, len(postorder)-1)

}
