import { runLRU } from '../types.ts'



class LRUNode {
	constructor(label?: string) {
		this._label_ = label || ''
	}
	next: LRUNode | null = null
	last: LRUNode | null = null
	val: any = null
	key: any = null
	_label_: string = ''

}

class LRUCache {
	capacity: number
	size: number
	head: LRUNode
	tail: LRUNode
	constructor(capacity: number) {
		this.capacity = capacity
		this.size = 0
		this.head = new LRUNode('head')
		this.tail = new LRUNode('tail')

		this.head.next = this.tail
		this.tail.last = this.head

	}

	get(key: number): number {
		let cur: LRUNode | null = this.head
		while (cur != null) {
			if (cur.key == key) {
				this.moveToHead(cur)
				return cur.val
			} else {
				cur = cur.next
			}
		}
		return -1
	}

	put(key: number, value: number): void {
		if (this.get(key) == -1) {
			const n = new LRUNode()
			n.key = key
			n.val = value
			if (this.size < this.capacity) {
				this.size++
			} else {
				// a b  tail
				this.tail.last!.last!.next = this.tail
				this.tail.last = this.tail.last!.last

			}
			this.head!.next!.last = n
			n.next = this.head.next

			n.last = this.head
			this.head.next = n
		} else {
			this.head.next!.val = value
		}
	}
	moveToHead(node: LRUNode) {

		node.last!.next = node.next
		node.next!.last = node.last


		node.next = this.head.next
		this.head.next!.last = node

		node.last = this.head
		this.head.next = node


	}
}


runLRU(LRUCache, ["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"], [[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]])

/**
 * Your LRUCache object will be instantiated and called as such:
 * var obj = new LRUCache(capacity)
 * var param_1 = obj.get(key)
 * obj.put(key,value)
 */

