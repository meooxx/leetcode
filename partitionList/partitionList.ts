
import { ListNode, createNodeList } from '../types.ts'
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

function partition(head: ListNode | null, x: number): ListNode | null {
	const node1 = new ListNode()
	const node2 = new ListNode()
	let tail1 = node1
	let tail2 = node2
	let cur = head
	// tail->xxx
	//    3 [1]  5
	// pre -> 1 
	//        3 ->5
	while (cur != null) {
		if (cur.val < x) {

			tail1.next = cur
			tail1 = tail1.next

		} else {
			tail2.next = cur
			tail2 = tail2.next

		}
		cur = cur.next

	}
	tail2.next = null
	tail1.next = node2.next
	return node1.next

};


partition(
	createNodeList([1, 4, 3, 2, 5, 2]), 3
)