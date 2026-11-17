import { TreeNode } from '../types.ts'

function preorderTraversal(root: TreeNode | null): number[] {
	if (root == null) return []
	let cur = root
	const stack: TreeNode[] = []
	const result: number[] = []
	while (stack.length > 0 || cur != null) {
		while (cur != null) {
			result.push(cur.val)
			stack.push(cur)
			cur = cur.left!
		}
		const r = stack.pop()!
		cur = r.right!
	}
	return result
};