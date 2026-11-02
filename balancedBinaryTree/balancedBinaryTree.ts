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

function isBalanced(root: TreeNode | null): boolean {
	if (root === null) return true
	const depth = getHeight(root)
	return depth == -1 ? false : true
};

function getHeight(node: TreeNode | null): number {
	if (node === null) return 0
	const leftDepth = getHeight(node.left)
	if (leftDepth == -1) return -1
	const rightDepth = getHeight(node.right)
	if (rightDepth === -1) return -1
	if (Math.abs(leftDepth - rightDepth) <= 1) {
		return 1 + Math.max(leftDepth, rightDepth)
	}
	return -1




}