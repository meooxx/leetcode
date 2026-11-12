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

function sumNumbers(root: TreeNode | null): number {
	const answer = { sum: 0 }
	getNext(root, answer, 0)
	return answer.sum
};

function getNext(n: TreeNode | null, ans: { sum: number }, pre: number) {

	if (n == null) return null
	if (n.left == null && n.right === null) {
		ans.sum += pre * 10 + n.val
		return
	}
	getNext(n.left, ans, pre * 10 + n.val)
	getNext(n.right, ans, pre * 10 + n.val)


}