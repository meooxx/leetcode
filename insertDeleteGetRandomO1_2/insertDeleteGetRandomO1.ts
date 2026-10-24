class RandomizedSet {
	map: Map<number, number>
	list: number[]
	constructor() {
		this.map = new Map<number, number>()
		this.list = [] as number[];
	}

	insert(val: number): boolean {
		if (this.map.has(val)) return false
		else {
			this.map.set(val, this.list.length)
			this.list.push(val)
			return true
		}
	}

	remove(val: number): boolean {
		if (!this.map.has(val)) return false
		const idx = this.map.get(val)!;
		const list = this.list
		const last = list.length - 1;

		[list[idx], list[last]] = [list[last], list[idx]];
		this.map.set(list[idx], idx)
		this.list.pop();
		this.map.delete(val)

		return true

	}

	getRandom(): number {

		// 0 - 1
		return this.list[~~(Math.random() * this.list.length)]
	}
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * var obj = new RandomizedSet()
 * var param_1 = obj.insert(val)
 * var param_2 = obj.remove(val)
 * var param_3 = obj.getRandom()
 */