function findOrder(numCourses: number, prerequisites: number[][]): number[] {

	const graph: Map<number, Set<number>> = new Map()
	const degree: Map<number, number> = new Map
	for (let i = 0; i < numCourses; i++) {
		degree.set(i, 0)
	}

	for (const course of prerequisites) {
		const [c, r] = course
		if (!graph.has(c)) {
			graph.set(c, new Set())
		}
		graph.get(c)!.add(r)
		degree.set(c, degree.get(c)! + 1)
	}
	const stack: number[] = [];
	[...degree.entries()].forEach(v => {
		const [c, count] = v
		if (count == 0) {
			stack.push(c)
		}
	})
	const result = []
	while (stack.length > 0) {
		const cur = stack.shift()!;
		result.push(cur);
		[...graph.entries()].forEach(v => {
			const [k, r] = v
			if (r.has(cur)) {
				r.delete(cur)
				degree.set(k, degree.get(k)! - 1)
				if (degree.get(k) === 0) {
					stack.push(k)
				}
			}
		})
	}
	if (result.length < numCourses) {
		return []
	}
	return result
};