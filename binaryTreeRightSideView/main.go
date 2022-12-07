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

//    1     <----- 
// 2    3   <-----   sight 
//  7   4   <-----
//   8   5  <-----
// 	  9     <-----
// it  equals the rightest of  evey level
// 1 3 4 5 9
func rightSideView(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}

	stack := []*TreeNode{root}

	for len(stack) > 0 {
		size := len(stack)
		result = append(result, stack[size-1].Val)
		for i := 0; i < size; i++ {
			node := stack[0]
			stack = stack[1:]
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
		}
	}

	return result

}
