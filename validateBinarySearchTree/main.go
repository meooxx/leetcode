package main

import (
	"fmt"
)

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

func main() {
	fmt.Println(isValidBST(&TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}))
}

var InfP = 1 << 31
var InfN = -InfP - 1

func isValidBST(root *TreeNode) bool {

	return isValid(root, InfN, InfP)
}

func isValid(root *TreeNode, min, max int) bool {

	if root == nil {
		return true
	}
	if root.Val <= min || root.Val >= max {
		return false
	}

	if !isValid(root.Left, min, root.Val) {
		return false
	}
	if !isValid(root.Right, root.Val, max) {
		return false
	}

	return true
}

// in-order traverse will get a sorted array
// 1 [2] 3  [5] 6 [7]  判断pre < 当前节点就可以了

//         5
//       2     7
//    1    3   6
//nil  nil
// push(5), push(2), push(1) 
//                   pop(1) 
// 					 1.right == nil 
//          pop(2)
//             2.right == 3
//          push(3)
//          pop(3)
//            3.right == nil
// pop(5)
//   
func isValidBST2(root *TreeNode) bool {
	stack := []*TreeNode{}
	var pre *TreeNode
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if pre != nil && pre.Val >= root.Val {
			return false
		}
		pre = root
		root = root.Right
	}
	return true
}
