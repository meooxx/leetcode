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

func traverse(root *TreeNode, parent *TreeNode, roadmap *map[int]*TreeNode) {
	rd := *roadmap
	if root == nil {
		return
	}
	rd[root.Val] = parent
	traverse(root.Left, root, roadmap)
	traverse(root.Right, root, roadmap)
}

func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	roadmap := map[int]*TreeNode{}
	if p.Val == q.Val {
		return p
	}
	traverse(root.Left, root, &roadmap)
	traverse(root.Right, root, &roadmap)
	cur := p
	path1 := []*TreeNode{cur}
	for cur != nil {
		v, ok := roadmap[cur.Val]
		if !ok {
			break
		}
		path1 = append(path1, v)
		cur = v
	}
	cur = q
	path2 := []*TreeNode{cur}
	for cur != nil {
		v, ok := roadmap[cur.Val]
		if !ok {
			break
		}
		path2 = append(path2, v)
		cur = v
	}
	index1 := len(path1) - 1
	index2 := len(path2) - 1
	var parent *TreeNode
	for index1 >= 0 && index2 >= 0 {
		if path1[index1].Val != path2[index2].Val {
			break
		}
		parent = path1[index1]
		index1--
		index2--
	}
	return parent
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	//  2     2,3   2,1 ||  1   1,3
	// 3 1
	
	if root == nil || root == q || root == p {
		return root
	}
	leftNode := lowestCommonAncestor(root.Left, p, q)
	rightNode := lowestCommonAncestor(root.Right, p, q)

	if leftNode != nil && rightNode != nil {
		return root
	} else if leftNode != nil {
		return leftNode
	} else if rightNode != nil {
		return rightNode
	} else {
		return nil
	}

}
