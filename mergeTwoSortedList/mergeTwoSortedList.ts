import { ListNode } from '../types.ts'

function mergeTwoLists(list1: ListNode | null, list2: ListNode | null): ListNode | null {
	let h1 = list1
	let h2 = list2
	const dummy = new ListNode()
	let curr = dummy
	while (h1 !== null && h2 !== null) {
		if (h1.val < h2.val) {
			curr.next = h1
			h1 = h1.next

		} else {
			curr.next = h2
			h2 = h2.next
		}
		curr = curr.next
	}
	while (h1 !== null) {
		curr.next = h1
		curr = curr.next
		h1 = h1.next
	}
	while (h2 !== null) {
		curr.next = h2
		curr = curr.next
		h2 = h2.next
	}
	return dummy.next
};