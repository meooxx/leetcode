import { ListNode } from '../types.ts'


function reverseKGroup(head: ListNode | null, k: number): ListNode | null {
	if (head === null) return null

	let curr: ListNode | null = head
	let dummy = new ListNode(0, head)
	let pre = dummy
	// 1 2 3 
	while (true) {
		let tail = curr as ListNode | null
		let count = 0

		while (count < k && tail !== null) {
			tail = tail.next
			count++
		}
		if (count < k) {
			break
		}
		// const tailNext = tail!.next
		// pre 1 2 3 4 -> pre 2 1 3 4
		// 1 2 3 4 5
		while (count > 1) {
			const next = curr!.next
			const nextNext = next!.next

			curr!.next = nextNext
			next!.next = pre.next
			pre.next = next
			count--

		}
		pre = curr!
		curr = curr!.next
	}
	return dummy.next
};

reverseKGroup(new ListNode(1, new ListNode(2, new ListNode(3, new ListNode(4, new ListNode(5))))), 3)