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

function maxPathSum(root: TreeNode | null): number {
	let answer = -Infinity
	const maxPath = (n: TreeNode | null) => {
		if (n == null) return 0
		const maxLeft = maxPath(n.left)
		const maxRight = maxPath(n.right)
		const maxLR = Math.max(maxLeft, maxRight)
		const oneNode = Math.max(n.val, n.val + maxLR)
		const allNode = Math.max(oneNode, n.val + maxLeft + maxRight)

		answer = Math.max(answer, allNode)

		return oneNode
	}
	maxPath(root)
	return answer
};