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

	for i := range left {
		result = append(result, fmt.Sprintf("%d->%s", root.Val, left[i]))
	}

	for i := range right {
		result = append(result, fmt.Sprintf("%d->%s", root.Val, right[i]))
	}
	if len(result) == 0 {
		return []string{fmt.Sprint(root.Val)}
	}
	return result
}
