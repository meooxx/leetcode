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
import { TreeNode, createTreeNode } from '../types.ts'
function isSameTree1(p: TreeNode | null, q: TreeNode | null): boolean {
	if (p == null || q == null) {
		if (p === q) return true
		if (p !== null || q !== null) return false
	}

	return p!.val == q!.val && isSameTree(p!.left, q!.left) && isSameTree(p.right, q.right)


};

function isSymmetric(root: TreeNode | null): boolean {
	if (root == null) return true
	let stack = [root.left, root.right]
	while (stack.length > 0) {
		const [first, second] = stack
		stack = stack.slice(2)

		if (first == null && second == null) continue
		if (first == null || second === null) return false

		if (first.val !== second.val) return false
		stack.push(first.left, second.right)
		stack.push(first.right, second.left)
	}
	return true


};
isSymmetric(createTreeNode([1, 2, 2, 3, 4, 4, 3]))

