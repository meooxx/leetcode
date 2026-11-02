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

function minDepth(root: TreeNode | null): number {
	if (root == null) return 0
	if (root.left == null && root.right == null) return 1
	if (root.left == null || root.right == null) {
		return 1 + minDepth(root.left || root.right)
	}
	return 1 + Math.min(minDepth(root.left), minDepth(root.right))

};