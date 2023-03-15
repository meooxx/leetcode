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

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// if root.Left == nil && root.Right == nil {
	//     return 1
	// }
	left := countNodes(root.Left)
	right := countNodes(root.Right)
	return left + right + 1
}
// less O(n)
// if height of left and right are equal, the anwser = 1<<height -1
// else 1 + count(left) + count(right)
func countNodes1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	hl, hr := 0, 0
	for cur := root; cur != nil; cur = cur.Left {
		hl++
	}
	for cur := root; cur != nil; cur = cur.Right {
		hr++
	}
	if hl == hr {
		return 1<<hl - 1
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)

}
