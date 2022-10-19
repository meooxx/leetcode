package main

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
func sortedArrayToBST(nums []int) *TreeNode {
	var constructTree func(nums []int, left, right int) *TreeNode

	constructTree = func(nums []int, left, right int) *TreeNode {
		if left > right {
			return nil
		}
		mid := (right-left)/2 + left
		node := &TreeNode{
			Val: nums[mid],
		}
		node.Left = constructTree(nums, left, mid-1)
		node.Right = constructTree(nums, mid+1, right)
		return node
	}
	return constructTree(nums, 0, len(nums)-1)

}

func sortedArrayToBST1(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	half := len(nums) / 2
	node := &TreeNode{Val: nums[half]}

	node.Left = sortedArrayToBST(nums[:half])
	node.Right = sortedArrayToBST(nums[half+1:])
	return node
}
