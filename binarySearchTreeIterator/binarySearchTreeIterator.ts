/**
 * Definition for a binary tree node.
 * class TreeNode {
 *     val: number
 *     left: TreeNode | null
 *     right: TreeNode | null
 *     constructor(val?: number, left?: TreeNode | null, right?: TreeNode | null) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.left = (left===undefined ? null : left)
 *         this.right = (right===undefined ? null : right)
 *     }
 * }
 */

class BSTIterator {
	stack: TreeNode[] = []
	head: TreeNode | null
	constructor(root: TreeNode | null) {
		this.head = root
	}

	next(): number {
		while (this.head) {
			const head = this.head
			this.stack.push(this.head)
			this.head = head.left
		}
		if (this.stack.length === 0) return null
		const root = this.stack.pop()

		this.head = root.right
		return root.val


	}

	hasNext(): boolean {
		return this.head != null || this.stack.length > 0
	}
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * var obj = new BSTIterator(root)
 * var param_1 = obj.next()
 * var param_2 = obj.hasNext()
 */