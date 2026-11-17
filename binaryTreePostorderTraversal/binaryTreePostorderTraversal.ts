import { TreeNode } from '../types.ts'

function postorderTraversal(root: TreeNode | null): number[] {
	const answer: number[] = []
	if (root == null) return answer

	const stack: TreeNode[] = [root]
	while (stack.length > 0) {
		const top = stack.pop()!
		answer.unshift(top.val)
		if (top.left) stack.push(top.left)
		if (top.right) stack.push(top.right)

	}
	return answer
};

function postorderTraversal2(root: TreeNode | null): number[] {
	const answer: number[] = []
	const traversal = (n: TreeNode | null) => {
		if (n === null) return
		traversal(n.left)
		traversal(n.right)
		answer.push(n.val)
	}
	traversal(root)
	return answer
};