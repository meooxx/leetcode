import { ListNode } from '../types.ts'

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

function removeElements(head: ListNode | null, val: number): ListNode | null {
	const dummy = new ListNode()
	let pre = dummy
	let cur = head
	while (cur != null) {
		if (cur.val !== val) {

			pre.next = cur
			pre = pre.next


		}
		cur = cur.next
	}
	pre.next = null
	return dummy.next

}