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

function buildTree(preorder: number[], inorder: number[]): TreeNode | null {
	if (preorder.length == 0) return null
	if (preorder.length === 1) return new TreeNode(preorder[0])
	const head = preorder[0]
	const headIndex = inorder.indexOf(head)
	const leftPart = inorder.slice(0, headIndex)
	const rightPart = inorder.slice(headIndex + 1)
	const headerTree = new TreeNode(head)
	headerTree.left = buildTree(preorder.slice(1, 1 + leftPart.length), leftPart)
	headerTree.right = buildTree(preorder.slice(1 + leftPart.length), rightPart)
	return headerTree

};