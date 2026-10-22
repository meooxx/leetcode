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



export class TreeNode {
	val: number
	left: TreeNode | null
	right: TreeNode | null
	constructor(val?: number, left?: TreeNode | null, right?: TreeNode | null) {
		this.val = (val === undefined ? 0 : val)
		this.left = (left === undefined ? null : left)
		this.right = (right === undefined ? null : right)
	}
}

type LRUCache = {
	get(key: number): number
	put(key: number, val: number): void
}
type LRUConstructor = new (size: number) => LRUCache

export function runLRU(LRU: LRUConstructor, action: string[], val: number[][]): Array<number | null> {
	let lru: LRUCache | null = null
	const result: Array<number | null> = []
	action.forEach((a, index) => {
		switch (a) {
			case 'LRUCache':
				lru = new LRU(val[index][0])
				result.push(null)
				break
			case 'get': {
				const v = lru!.get(val[index][0])
				result.push(v)

				break
			}
			case 'put': {
				lru!.put(val[index][0], val[index][1])
				result.push(null)
				break
			}
		}
	})
	return result

}