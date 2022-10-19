package main

import "fmt"

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
	}
	fmt.Println(isSymmetric(root))
	fmt.Println(isSymmetric1(root))

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
//     1
//   2  2      push left, right, [2, 2]
//  34 43      pop a,b [2, 2], 推入 a.left, b.right, a.right, b.left
//             [3,3 4,4], pop 3,3 -> pop 4,4
// passed 100% fast
func isSymmetric(root *TreeNode) bool {
	if root.Left == nil && root.Right == nil {
		return true
	}
	stack := []*TreeNode{
		root.Left, root.Right,
	}
	for len(stack) >= 2 {
		a := stack[0]
		b := stack[1]
		stack = stack[2:]
		if a == nil && b == nil {
			continue
		// 1 a or b is nil, if a = nil, b != nil(must be, cuz a!=nil && b!=nil)
		//   return false
		// if b = nil, a != nil(must be), return false
		// 2 or a.Val != b.Val
		} else if a == nil || b == nil || a.Val != b.Val {
			return false
		}
		stack = append(stack, a.Left, b.Right, a.Right, b.Left)
	}
	return true

}
// passed
func isSymmetric1(root *TreeNode) bool {
	var isSame func(a, b *TreeNode) bool
	isSame = func(a, b *TreeNode) bool {
		if a == nil || b == nil {
			return a == b
		}
		if a.Val == b.Val {
			return isSame(a.Left, b.Right) && isSame(a.Right, b.Left)
		}
		return false
	}
	return isSame(root.Left, root.Right)
}
