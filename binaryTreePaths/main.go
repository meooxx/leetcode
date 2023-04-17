package main

import "fmt"

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

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	left := binaryTreePaths(root.Left)
	right := binaryTreePaths(root.Right)
	result := []string{}

	for _, i := range left {
		result = append(result, fmt.Sprintf("%d->%s", root.Val, i))
	}
	for _, i := range right {
		result = append(result, fmt.Sprintf("%d->%s", root.Val, i))
	}
	if len(result) == 0 {
		result = append(result, fmt.Sprint(root.Val))
	}
	return result
}
