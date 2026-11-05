import { _Node } from '../types.ts'

/**
 * Definition for _Node.
 * class _Node {
 *     val: number
 *     left: _Node | null
 *     right: _Node | null
 *     next: _Node | null
 *     constructor(val?: number, left?: _Node, right?: _Node, next?: _Node) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.left = (left===undefined ? null : left)
 *         this.right = (right===undefined ? null : right)
 *         this.next = (next===undefined ? null : next)
 *     }
 * }
 */

function connect2(root: _Node | null): _Node | null {
	if (root == null) return root
	const heads = [root]
	// const stack: Array<_Node | null> = []
	let index = 0

	while (index < heads.length) {
		let size = heads.length
		let pre = null
		while (index < size) {
			const top = heads[index]
			index++
			if (pre) pre.next = top
			if (top.left) {

				heads.push(top.left)

			}
			if (top.right) {
				if (pre) pre.next = top
				heads.push(top.right)

			}
			pre = top

			// stack.push(top.left, top.right)
		}
	}
	return root
};

///         1
//  2(root)    3 
//4   [5]      [6] 7  
//left  right  next.left
// connect left - right
// 1 root.left.next =root.right
// connect right -> neigbour.left
// 2 root.right.next = root.next.left

function connect(root: _Node | null): _Node | null {
	if (root == null) return root
	let head: _Node | null = root
	while (head !== null) {
		let curr: _Node | null = head
		while (curr != null) {
			if (curr.left) {
				curr.left.next = curr.right
				if (curr.next) {
					curr.right!.next = curr.next.left
				}

			}
			curr = curr.next

		}
		head = head.left
	}
	return root
}