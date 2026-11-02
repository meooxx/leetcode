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

function levelOrderBottom(root: TreeNode | null): number[][] {
	if (root == null) return []

	const stack: Array<TreeNode | null> = [root]

	const result: number[][] = []
	let pointer = 0

	while (pointer < stack.length) {
		const size = stack.length
		const levelNodes = []

		while (pointer < size) {
			const root = stack[pointer++]

			if (root) {
				levelNodes.push(root.val)

				if (root.left || root.right) {
					stack.push(root.left, root.right)
				}
			}
		}
		result.push(levelNodes)
	}


	return result.reverse()

};