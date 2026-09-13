import { ListNode } from '../types.ts'

function swapPairs(head: ListNode | null): ListNode | null {
	if (head === null) return null
	if (head.next === null) return head
	const dummy = new ListNode()
	let pre = dummy
	let curr = head as ListNode | null

	// c next
	// 1  2    3 4
	while (curr !== null && curr.next !== null) {
		const next = curr.next
		const nextNext = next.next ?? null

		pre.next = next
		next.next = curr
		curr.next = nextNext

		pre = curr
		curr = curr.next





	}
	return dummy.next
};