import { TreeNode } from '../types.ts'
function levelOrder(root: TreeNode | null): number[][] {
	const stack = [root]
	const answer: number[][] = []
	if (root == null) return answer
	while (stack.length > 0) {
		let size = stack.length
		const nodes: number[] = []
		while (size-- > 0) {
			const root = stack.shift()!
			if (!root) continue
			nodes.push(root.val)
			if (root.left || root.right) {
				stack.push(root.left, root.right)
			}


		}
		answer.push(nodes)
	}
	return answer
};