import { ListNode } from '../types.ts'

function reorderList(head: ListNode | null): void {
	if (head === null) return
	let fast = head
	let slow = head
	while (fast.next != null && fast.next.next != null) {
		fast = fast.next.next!
		slow = slow.next!
	}
	// reverse the second half of linkNode
	let newHead = slow.next

	// let tail = slow
	let cur = newHead
	//   [1]  2 3 4
	while (cur != null && cur.next != null) {
		const next = cur.next
		cur.next = next!.next
		next.next = newHead
		newHead = next
	}
	slow.next = null
	cur = head
	// 1 2  -  4 3
	// 
	while (newHead != null) {
		const next: ListNode = cur.next!
		cur.next = newHead
		newHead = newHead!.next
		cur.next!.next = next
		cur = cur.next!.next
	}
};