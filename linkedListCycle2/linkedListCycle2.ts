import { ListNode } from '../types.ts'


function detectCycle(head: ListNode | null): ListNode | null {
	//  2(x + y) = x +2y + z
	// x = z

	if (head === null) return head
	if (head.next == null) return null
	let fast = head
	let slow = head
	while (fast !== null && fast.next !== null) {
		fast = fast.next.next!
		slow = slow.next!
		if (fast == slow) break


	}
	if (fast !== slow) return null
	let cur = head
	while (cur != slow) {
		cur = cur.next!
		slow = slow.next!
	}
	return cur
};