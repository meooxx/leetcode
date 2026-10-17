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
import { ListNode, createNodeList } from '../types.ts'

function reverseBetween(head: ListNode | null, left: number, right: number): ListNode | null {

	if (right - left === 0) return head
	let reverseHead = new ListNode()
	reverseHead.next = head

	let pre = reverseHead
	let count = 0
	// p   t
	// 1   2  3 4 5, 2, 3
	while (pre != null && count < left - 1) {
		count++
		pre = pre.next!
	}
	const tail = pre.next!


	while (pre != null && left < right) {
		left++
		const next = pre.next
		pre.next = tail.next
		tail.next = tail.next!.next
		pre.next!.next = next

	}




	return reverseHead.next
};

reverseBetween(createNodeList([1, 2, 3, 4, 5]), 2, 3)