import { TreeNode } from '../types.ts'
/**
 * Definition for singly-linked list.
 * class ListNode {
 *     val: number
 *     next: ListNode | null
 *     constructor(val?: number, next?: ListNode | null) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.next = (next===undefined ? null : next)
 *     }
 * }
 */

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

function sortedListToBST(head: ListNode | null): TreeNode | null {
	if (head == null) return null
	if (head.next == null) return new TreeNode(head.val)

	let fast = head
	let slow = head
	let pre = null
	while (fast != null && fast.next !== null) {
		fast = fast.next.next
		pre = slow
		slow = slow.next

	}
	const root = new TreeNode(slow.val)
	root.right = sortedListToBST(slow.next)
	pre.next = null
	root.left = sortedListToBST(head)
	return root
};