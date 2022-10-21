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
func main() {}

func pathSum(root *TreeNode, targetSum int) [][]int {
	result := [][]int{}

	var searchPath func(node *TreeNode, current []int, rest int)
	searchPath = func(node *TreeNode, current []int, rest int) {
		if node == nil {
			return
		}
		current = append(current, node.Val)
		if rest-node.Val == 0 && node.Left == nil && node.Right == nil {

			item := make([]int, len(current))
			copy(item, current)
			result = append(result, item)

			return
		}
		searchPath(node.Left, current, rest-node.Val)
		searchPath(node.Right, current, rest-node.Val)
	}
	searchPath(root, []int{}, targetSum)

	return result

}
