package main

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

func postorderTraversal(root *TreeNode) []int {
	s := []*TreeNode{root}
	result := []int{}
	if root == nil {
		return result
	}
	for len(s) > 0 {
		node := s[len(s)-1]
		s = s[:len(s)-1]
		result = append(result, node.Val)
		if node.Left != nil {
			s = append(s, node.Left)
		}
		if node.Right != nil {
			s = append(s, node.Right)
		}
	}
	for left, right := 0, len(result)-1; left < right; {
		result[left], result[right] = result[right], result[left]
		left++
		right--
	}
	return result
}
