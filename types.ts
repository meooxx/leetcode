export class ListNode {
	val: number
	next: ListNode | null
	constructor(val?: number, next?: ListNode | null) {
		this.val = (val === undefined ? 0 : val)
		this.next = (next === undefined ? null : next)
	}
}

export function createNodeList(nodes: number[]): ListNode {
	const dummy = new ListNode()
	let pre = dummy
	for (const n of nodes) {
		pre.next = new ListNode(n)
		pre = pre.next
	}
	return dummy.next!
}



