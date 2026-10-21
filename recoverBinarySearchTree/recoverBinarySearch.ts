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

/**
 Do not return anything, modify root in-place instead.
 */
function recoverTree(root: TreeNode | null): void {
	let firstNode: TreeNode | null = null
	let secondNode: TreeNode | null = null

	let cur = root
	let pre: TreeNode | null = null
	const stack: TreeNode[] = []
	while (stack.length > 0 || cur != null) {
		while (cur != null) {
			stack.push(cur)
			cur = cur.left
		}
		const r = stack.pop()!
		// 4 6, 5 
		// => got pre 6, node 4
		if (pre != null && pre.val! >= r!.val) {
			// default to the first pair
			if (firstNode === null) {
				// find a bigger node
				firstNode = pre
			}
			// 3 2 1, 3 > 2 => f:3, second:2
			//              =>      1 < 2 => second:1
			// find the smaller node
			secondNode = r

		}
		pre = r
		cur = r.right
	}

	[firstNode!.val, secondNode!.val] = [secondNode!.val, firstNode!.val]
};