/**
 * Definition for _Node.
 * class _Node {
 *     val: number
 *     next: _Node | null
 *     random: _Node | null
 * 
 *     constructor(val?: number, next?: _Node, random?: _Node) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.next = (next===undefined ? null : next)
 *         this.random = (random===undefined ? null : random)
 *     }
 * }
 */
class _Node {
	val: number
	next: _Node | null
	random: _Node | null

	constructor(val?: number, next?: _Node, random?: _Node) {
		this.val = (val === undefined ? 0 : val)
		this.next = (next === undefined ? null : next)
		this.random = (random === undefined ? null : random)
	}
}

function copyRandomList(head: _Node | null): _Node | null {
	if (head === null) return head
	let cur = head
	while (cur !== null) {
		const newNode = new _Node()
		const next = cur.next!
		newNode.val = cur.val

		newNode.next = next
		cur.next = newNode
		cur = next
	}
	cur = head
	while (cur != null) {
		if (cur.random) {
			cur.next!.random = cur.random.next
		}
		cur = cur.next!.next!
	}


	const result = new _Node
	cur = head
	let pre = result
	while (cur != null && cur.next != null) {
		const next = cur.next
		cur.next = next.next
		pre.next = next

		pre = pre.next

		cur = cur.next!
	}
	return result.next
};