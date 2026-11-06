import { _Node } from '../types.ts'
/**
 * Definition for _Node.
 * class _Node {
 *     val: number
 *     left: _Node | null
 *     right: _Node | null
 *     next: _Node | null
 * 
 *     constructor(val?: number, left?: _Node, right?: _Node, next?: _Node) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.left = (left===undefined ? null : left)
 *         this.right = (right===undefined ? null : right)
 *         this.next = (next===undefined ? null : next)
 *     }
 * }
 */


function connect(root: _Node | null): _Node | null {
	if (root == null) return root
	const heads: Array<_Node | null> = [root]
	// const stack: Array<_Node | null> = []
	let index = 0
	while (index < heads.length) {
		let size = heads.length
		let pre = null
		while (index < size) {
			const top = heads[index]!
			index++
			if (pre) pre.next = top
			pre = top
			if (top.left) heads.push(top.left)
			if (top.right) heads.push(top.right)
		}
	}
	return root
};