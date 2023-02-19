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
type BSTIterator struct {
	stack []int
	index int
}

func Constructor(root *TreeNode) BSTIterator {
	cur := root
	stack := []*TreeNode{}
	result := []int{}
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, cur.Val)
		cur = cur.Right
	}
	return BSTIterator{result, 0}
}

func (this *BSTIterator) Next() int {

	v := this.stack[this.index]
	this.index++
	return v
}

func (this *BSTIterator) HasNext() bool {
	return this.index < len(this.stack)
}

/**
* Your BSTIterator object will be instantiated and called as such:
* obj := Constructor(root);
* param_1 := obj.Next();
* param_2 := obj.HasNext();
 */
