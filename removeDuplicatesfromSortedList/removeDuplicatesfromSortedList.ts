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

function deleteDuplicates(head: ListNode | null): ListNode | null {
	let answer = new ListNode()
	let current: ListNode | null = head
	let pre = answer
	pre.next = head

	// 1 2  2 3 3 null
	while (current != null) {
		while (current.next != null && current.val == current.next.val) {

			current = current.next

		}
		// indicating that we didn't find the duplicates
		if (pre.next === current) {
			pre = pre.next
		} else {
			// find a duplicate and point to the next
			pre.next = current.next
		}
		current = current.next

	}
	return answer.next

};

deleteDuplicates(new ListNode(
	1, new ListNode(
		2, new ListNode(3,
			new ListNode(
				3, new ListNode(4,
					new ListNode(4)
				)
			)
		)
	)
))