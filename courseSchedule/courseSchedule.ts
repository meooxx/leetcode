// construct a structure of 
// edges: number, [1,0], [1,2] => 1 require 0 and 2, {1: 2}
// dep: 0:[1], 2: [1], 1 relies on 0, and 2
// when remove 0 or 2, we can simply reduct 1 edge from 1

function canFinish(numCourses: number, prerequisites: number[][]): boolean {
	const edges = new Map<number, number>()
	const dep = new Map<number, Set<number>>()
	for (let i = 0; i < numCourses; i++) {
		edges.set(i, 0)
		dep.set(i, new Set)
	}
	for (const pre of prerequisites) {
		edges.set(pre[0], edges.get(pre[0])! + 1)
		dep.get(pre[1])!.add(pre[0])
	}
	const stack: number[] = [];
	for (const [key, val] of edges) {
		if (val == 0) {
			stack.push(key)
		}

	}
	while (stack.length > 0) {
		const c = stack.shift()

		for (const val of dep.get(c)!) {
			edges.set(val, edges.get(val)! - 1)
			if (edges.get(val) === 0) {
				stack.push(val)
			}
		}



	}
	for (const v of edges.values()) {
		if (v > 0) return false
	}
	return true
};