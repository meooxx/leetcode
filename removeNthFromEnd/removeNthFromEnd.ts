
import {
	ListNode
} from '../types.ts'

function removeNthFromEnd(head: ListNode | null, n: number): ListNode | null {

	const nodes: ListNode[] = []
	let curr = head
	while (curr != null) {
		const temp = curr.next
		curr.next = null
		nodes.push(curr)
		curr = temp

	}
	const dump = new ListNode()
	curr = dump
	nodes.forEach((node, index) => {
		if (index === nodes.length - n) {
			return
		}
		curr!.next = node
		curr = curr!.next
	})
	return dump.next
};