package main

import (
	"fmt"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */
func main() {
	root := &TreeNode{
		Val: 6,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 3,
				},
				Right: &TreeNode{
					Val: 5,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 7,
			},
			Right: &TreeNode{
				Val: 9,
			},
		},
	}
	root2 := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
	}
	// fmt.Println(lowestCommonAncestor1(root2, root2, root2.Left))
	fmt.Println(lowestCommonAncestor(root2, root2, root2.Left))
	fmt.Println(lowestCommonAncestor(root, root.Left, root.Right))
}

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

// there are three cases:
// let large = max(p, q), small = min(p, q)
// if root.Val > large, then the lcp is in the right tree
// if root.Val < small, then  the lcp is in the left tree
// if small <= root.Left && root.Right >= large
//    root is the lcp
// 
//       6                     6                   4
//  2       8           		2      8           3(p)  5(q)
// 4(p) 3                      7(p)  9(q)
//       5(q)
// 6 > max(4,5)               6 < min(7, 9)       3<4<5

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	large := max(p.Val, q.Val)
	small := min(p.Val, q.Val)
	for {
		if large >= root.Val && small <= root.Val {
			return root
		} else if small < root.Val {
			root = root.Left
		} else {
			root = root.Right
		}
	}
}
