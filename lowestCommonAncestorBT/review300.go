package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// if root == q || root == q that means p or q is root
	// which is the lowest common ancestor
	// it's unnessary recurse the left and right, 
	// the anwser is the ROOT
	//  3 root == 3, p == 3, q == 4
	// 4
	if root == nil || root == p || root == q {
		return root
	}

	n1 := lowestCommonAncestor(root.Left, p, q)
	n2 := lowestCommonAncestor(root.Right, p, q)

	if n1 != nil && n2 != nil {
		return root
	} else if n1 != nil {
		return n1
	} else if n2 != nil {
		return n2
	} else {
		return nil
	}
}
