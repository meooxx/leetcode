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

function pathSum(root: TreeNode | null, targetSum: number): number[][] {

	const result: number[][] = []
	const nextNode = (n: TreeNode | null, sum: number, candidates: number[]) => {
		if (n == null) return
		if (n.val == sum && n.left == null && n.right == null) {

			result.push(candidates.concat(n.val))
			return
		}
		nextNode(n.left, sum - n.val, [...candidates, n.val])
		nextNode(n.right, sum - n.val, [...candidates, n.val])
	}
	nextNode(root, targetSum, [])
	return result
};

