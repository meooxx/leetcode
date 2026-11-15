
/**
 * Definition for _Node.

 
 */

class _Node {
	val: number
	neighbors: _Node[]

	constructor(val?: number, neighbors?: _Node[]) {
		this.val = (val === undefined ? 0 : val)
		this.neighbors = (neighbors === undefined ? [] : neighbors)
	}
}
function cloneGraph(node: _Node | null): _Node | null {
	const visited = new Map<number, _Node | null>();
	const clone = (n: _Node | null): _Node | null => {
		if (n === null) return n
		const newNode = new _Node()
		newNode.val = n.val
		visited.set(n.val, newNode)
		n.neighbors!.forEach(nei => {
			if (visited.has(nei.val)) {
				newNode.neighbors.push(visited.get(nei.val)!)
				return
			}
			newNode.neighbors.push(clone(nei)!)
		})
		return newNode
	}
	return clone(node)
};