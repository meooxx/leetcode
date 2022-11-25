package main

// 本来打算next 时候计算下一个的, 可以但是很耗时
// 看了别人的答案发现, 不需要, 直接in-order 遍历
// 使用一个指针, 就行了
// main1 就是 next 的时候计算
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
	node  *TreeNode
	stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	cur := root
	stack := []*TreeNode{}
	for cur != nil {
		stack = append(stack, cur)
		cur = cur.Left
	}
	return BSTIterator{
		node:  cur,
		stack: stack,
	}
}

func (this *BSTIterator) Next() int {
	if this.node == nil {
		this.node = this.stack[len(this.stack)-1]
		this.stack = this.stack[:len(this.stack)-1]
		return this.node.Val
	} else {
		cur := this.node.Right
		if cur == nil {
			this.node = this.stack[len(this.stack)-1]
			this.stack = this.stack[:len(this.stack)-1]
			return this.node.Val
		} else {
			for cur != nil && cur.Left != nil {
				this.stack = append(this.stack, cur)
				cur = cur.Left
			}
			this.node = cur
			return this.node.Val
		}
	}

}

func (this *BSTIterator) HasNext() bool {
	return len(this.stack) > 0 || (this.node != nil && this.node.Right != nil)
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
