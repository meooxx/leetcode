import { ListNode } from '../types.ts'

class MinListNode extends ListNode {
	min: number = Infinity
	override next: MinListNode | null = null
}

class MinStack {
	head: MinListNode | null = null
	constructor() {

	}

	push(value: number): void {
		if (this.head == null) {
			this.head = new MinListNode(value)
			this.head.min = value
		} else {
			const n = new MinListNode(value)
			n.min = Math.min(this.head.min, value)
			n.next = this.head
			this.head = n
		}
	}

	pop(): void {
		this.head = this.head!.next
	}

	top(): number {
		return this.head!.val
	}

	getMin(): number {
		return this.head!.min
	}
}

/**
 * Your MinStack object will be instantiated and called as such:
 * var obj = new MinStack()
 * obj.push(value)
 * obj.pop()
 * var param_3 = obj.top()
 * var param_4 = obj.getMin()
 */