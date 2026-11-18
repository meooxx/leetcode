import { ListNode } from '../types.ts'

function insertionSortList(head: ListNode | null): ListNode | null {
	if (head === null) return head
	const dummy = new ListNode()
	let pre = dummy
	let curr = head
	while (curr !== null) {
		// 1 [4] 5    [3] 6
		while (pre.next != null && pre.next.val <= curr.val) {
			pre = pre.next
		}
		const next = curr.next
		curr.next = pre.next
		pre.next = curr
		pre = dummy
		curr = next
	}

	return dummy.next
};