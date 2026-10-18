import { TreeNode } from '../types.ts'

function inorderTraversal(root: TreeNode | null): number[] {
	const stack: TreeNode[] = []
	if (root == null) return []
	const result: number[] = []
	let cur: TreeNode | null = root
	while (stack.length > 0 || cur !== null) {
		while (cur != null) {
			stack.push(cur)
			cur = cur.left
		}
		const root = stack.pop()
		result.push(root!.val)
		cur = root!.right


	}
	return result
};