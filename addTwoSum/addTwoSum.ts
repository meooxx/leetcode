/**
 * Definition for singly-linked list.
 * class ListNode {
 *     val: number
 *     next: ListNode | null
 *     constructor(val?: number, next?: ListNode | null) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.next = (next===undefined ? null : next)
 *     }
 * }
 */
class ListNode {
	val: number
	next: ListNode | null
	constructor(val?: number, next?: ListNode | null) {
		this.val = (val === undefined ? 0 : val)
		this.next = (next === undefined ? null : next)
	}
}

function addTwoNumbers(l1: ListNode | null, l2: ListNode | null): ListNode | null {
	let head1 = l1
	let head2 = l2
	const result = new ListNode()
	result.next = new ListNode()
	let pre = result
	let rem = 0
	while (head1 != null || head2 != null) {
		pre.next = new ListNode()
		head1 = head1 ?? new ListNode()
		head2 = head2 ?? new ListNode()
		pre.next.val = (head1.val + head2.val + rem) % 10 
		rem = Math.floor((head1.val + head2.val + rem) / 10)
		head1 = head1.next
		head2 = head2.next
		pre = pre.next

	}
	
	if (rem != 0) {
		pre.next = new ListNode(rem)
	}

	return result.next
};