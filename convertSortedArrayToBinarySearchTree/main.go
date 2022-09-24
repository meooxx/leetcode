package main
func main(){}
type TreeNode struct {
    Val int
    Left *TreeNode
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
    return constructTree(nums)
    return constructTree2(nums, 0 , len(nums)-1)
}

func constructTree(nums []int) *TreeNode{
    if len(nums) == 0 {
        return nil
    }
    mid := len(nums) / 2
    root := nums[mid]
    node := &TreeNode{Val:root}
    left := constructTree(nums[:mid])
    right := constructTree(nums[mid+1:])
    node.Left = left
    node.Right = right
    return node
}

func constructTree2(nums []int, left, right int) *TreeNode{
    if left > right {
        return nil
    }
    midIndex := (right-left)/2 + left
    root := nums[midIndex]
    node := &TreeNode{Val:root}
    l := constructTree2(nums, left, midIndex-1)
    r := constructTree2(nums, midIndex+1, right)
    node.Left = l
    node.Right = r
    return node
}
