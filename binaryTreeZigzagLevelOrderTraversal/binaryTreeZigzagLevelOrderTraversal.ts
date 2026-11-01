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

function zigzagLevelOrder(root: TreeNode | null): number[][] {
	if (root === null) return [] as number[][]
	const stack: Array<TreeNode | null> = [root]
	let direction = false
	const result: number[][] = []
	let pointer = 0

	while (pointer < stack.length) {
		let size = stack.length
		const levelNodes: number[] = []
		while (pointer < size) {
			const top = stack[pointer++]
			if (top !== null) {
				if (direction) {
					levelNodes.unshift(top.val)
				} else {
					levelNodes.push(top.val)
				}


				if (top.left || top.right) {
					stack.push(top.left, top.right)
				}
			}
		}
		result.push(levelNodes)
		direction = !direction
	}
	return result
};