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
func flatten2(root *TreeNode) {

	var postOrderTraverse func(t *TreeNode)
	var pre *TreeNode

	postOrderTraverse = func(t *TreeNode) {
		if t == nil {
			return
		}
		postOrderTraverse(t.Right)
		postOrderTraverse(t.Left)
		t.Right = pre
		t.Left = nil
		pre = root

	}
	postOrderTraverse(root)
}

//  1        1     1       1     1
// 2 3      2      2       2      2
//5   4    5 3    5  3    5        5
//            4      4     3        3 
//                           4       4
func flatten3(root *TreeNode) {
	for root != nil {
		cur := root.Left
		for cur != nil && cur.Right != nil {
			cur = cur.Right
		}
		cur.Right = root.Right
		root.Left = nil
		root.Right = cur
		root = root.Right

	}
}

func flatten(root *TreeNode) {

	getNode(root)
}

func getNode(n *TreeNode) (*TreeNode, *TreeNode) {
	cur := n
	if cur.Left == nil && cur.Right == nil {
		return cur, cur
	}
	if cur.Left == nil {
		_, tail := getNode(cur.Right)
		return cur, tail
	}
	if cur.Right == nil {
		right, tail := getNode(cur.Left)
		cur.Right = right
		return cur, tail
	}

	leftHead, tailLeft := getNode(n.Left)
	rightHead, tailRight := getNode(n.Right)

	cur.Left = nil
	cur.Right = leftHead
	tailLeft.Right = rightHead

	return cur, tailRight
}
