
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

function buildTree(inorder: number[], postorder: number[]): TreeNode | null {
	if (inorder.length == 0) return null
	if (inorder.length === 1) return new TreeNode(inorder[0])

	const head = postorder[postorder.length - 1]
	const headIndex = inorder.indexOf(head)
	const left = inorder.slice(0, headIndex)
	const right = inorder.slice(headIndex + 1)

	const headTree = new TreeNode(head)
	headTree.left = buildTree(
		left,
		postorder.slice(0, left.length)
	)
	headTree.right = buildTree(
		right,
		postorder.slice(left.length, postorder.length - 1)
	)
	return headTree

};