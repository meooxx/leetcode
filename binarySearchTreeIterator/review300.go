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
	list  []int
	index int
}

func Constructor(root *TreeNode) BSTIterator {
	stack := []*TreeNode{}
	cur := root
	list := []int{}

	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		list = append(list, cur.Val)
		cur = cur.Right

	}
	return BSTIterator{
		list:  list,
		index: 0,
	}
}

func (this *BSTIterator) Next() int {
	val := this.list[this.index]
	this.index++
	return val
}

func (this *BSTIterator) HasNext() bool {
	return this.index < len(this.list)
}

/**
* Your BSTIterator object will be instantiated and called as such:
* obj := Constructor(root);
* param_1 := obj.Next();
* param_2 := obj.HasNext();
 */
