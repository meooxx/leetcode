import { TreeNode } from '../types.ts'

function generateTrees(n: number): Array<TreeNode | null> {

	return getTrees(1, n)
};


function getTrees(start: number, end: number): Array<TreeNode | null> {
	if (start > end) {
		return [null]
	}

	const result: Array<TreeNode | null> = []



	for (let i = start; i <= end; i++) {
		const leftTrees = getTrees(1, i - 1)
		const rightTrees = getTrees(end - i, end)

		for (const left of leftTrees) {
			for (const right of rightTrees) {
				result.push(new TreeNode(i, left, right))
			}
		}

	}
	return result

}