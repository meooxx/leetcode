import { TreeNode } from '../types.ts'

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

function rightSideView(root: TreeNode | null): number[] {
	if (root === null) return []
	const stack: TreeNode[] = [root]
	const answer: number[] = []
	while (stack.length > 0) {
		let size = stack.length
		const idx = answer.length
		while (size-- > 0) {
			const r = stack.shift()!
			answer[idx] = r.val
			if (r.left) stack.push(r.left)
			if (r.right) stack.push(r.right)
		}
	}
	return answer
};