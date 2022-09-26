package main

import "fmt"

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

//     1
//   2   3
//  4   5 6
func postorderTraversal(root *TreeNode) []int {
	stack := []*TreeNode{root}
	result := []int{}
    if root == nil {
        return result
    }
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, top.Val)
		if top.Left != nil {
			stack = append(stack, top.Left)
		}
		if top.Right != nil {
			stack = append(stack, top.Right)
		}
	}
	for left, right := 0, len(result)-1; left < right; {
		result[left], result[right] = result[right], result[left]
        left++
        right--

	}
	return result

}