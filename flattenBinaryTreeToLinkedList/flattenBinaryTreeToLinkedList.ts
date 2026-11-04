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

//   1 
// 2 3
//     4
//   pre: 4 =>3 =>2=>1
// 走到2, 3的时候

// 先是 findRight(n.right) 所以此时pre == 3
// 再走findRight(n.left), 设置node(2).right = pre
//   2 
//    3
//     4

function flatten(root: TreeNode | null): void {
	let pre: TreeNode | null = null
	const findRight = (n: TreeNode | null): void => {
		if (n == null) return
		findRight(n.right) // 先right, meaning pre = right
		findRight(n.left) //  then left, meaning lastLeft.right = pre 
		n.right = pre
		n.left = null
		pre = n
	}
	findRight(root)
};


function nextNode(n: TreeNode | null): TreeNode | null {
	if (n === null) return null
	if (n.left == null && n.right === null) return n
	const left = nextNode(n.left)
	const right = nextNode(n.right)

	if (left == null) return right
	left.right = n.right
	n.right = n.left
	n.left = null
	return right || left

}