
import {
	ListNode
} from '../types.ts'

function removeNthFromEnd1(head: ListNode | null, n: number): ListNode | null {

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

function removeNthFromEnd2(head: ListNode | null, n: number): ListNode | null {
	let count = 0
	let curr = head
	while (curr !== null) {
		curr = curr.next
		count++
	}
	let index = 0
	let dummy = new ListNode()
	curr = dummy
	// 1-2-3
	while (head !== null) {
		if (index === count - n) {
			if (head.next !== null) {
				curr.next = head.next
			} else {
				// this is the last node
				curr.next = null
			}
			return dummy.next
		}
		curr.next = head
		curr = curr!.next
		head = head.next
		index++
	}
	return dummy.next
}

function removeNthFromEnd(head: ListNode | null, n: number): ListNode | null {

	let fast = head
	let slow = head
	//    quick: start from here 3
	//    when quick point reach the end
	// 1- 2- 3- 4
	//    slow: 1, ending at 2 position
	//      1   2  [3  4]
	for (let i = 0; i < n; i++) {
		if (fast !== null) {
			fast = fast.next
		}
	}
	if (fast === null) return head!.next
	while (fast.next !== null) {
		fast = fast.next
		slow = slow!.next
	}
	slow!.next = slow!.next!.next
	return head

}