
import { ListNode } from '../types.ts'

function rotateRight(head: ListNode | null, k: number): ListNode | null {
	if (head === null) return head
	let fast: ListNode | null = head
	let count = 0
	while (fast != null) {
		count++
		fast = fast.next
	}
	fast = head
	let n = k % count
	// 12 345   2
	// 123
	while (n > 0 && fast != null) {
		fast = fast.next
		n--
	}

	let slow: ListNode | null = head
	while (fast!.next != null) {
		slow = slow!.next
		fast = fast!.next
	}
	fast!.next = head
	const next = slow?.next || null
	slow!.next = null

	return next


};