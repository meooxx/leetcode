import { ListNode } from '../types.ts'

function mergeKLists(lists: Array<ListNode | null>): ListNode | null {
	if (lists.length === 0) return null
	if (lists.length === 1) return lists[0]
	const left = mergeKLists(lists.slice(0, ~~(lists.length / 2)))
	const right = mergeKLists(lists.slice(~~(lists.length / 2)))
	return mergeTwoLists(left, right)
};

function mergeTwoLists(l1: ListNode | null, l2: ListNode | null): ListNode | null {
	if (l1 === null && l2 === null) {
		return null
	}
	if (l1 === null) {
		return l2
	}
	if (l2 === null) {
		return l1
	}
	const dummy = new ListNode()
	let curr = dummy
	while (l1 !== null && l2 !== null) {
		if (l1.val <= l2.val) {
			curr.next = l1
			curr = curr.next
			l1 = l1.next
		} else {
			curr.next = l2
			curr = curr.next
			l2 = l2.next
		}

	}
	if (l1 !== null) {
		curr.next = l1
	}
	if (l2 !== null) {
		curr.next = l2
	}

	return dummy.next
}