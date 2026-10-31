
import { TreeNode } from '../types.ts'
function isSymmetric1(root: TreeNode | null): boolean {
	if (root == null) return true
	let stack = [root.left, root.right]
	while (stack.length > 0) {
		const [first, second] = stack
		stack = stack.slice(2)

		if (first == null && second == null) continue
		if (first == null || second === null) return false

		if (first.val !== second.val) return false
		stack.push(first.left, second.right)
		stack.push(first.right, second.left)
	}
	return true


};

function isSymmetric(root: TreeNode | null): boolean {
	if (root == null) return root
	return next(root.left, root.right)

}
function next(n1: TreeNode | null, n2: TreeNode | null): boolean {
	if (n1 == null && n2 === null) return true
	if (n1 == null || n2 === null) return false
	if (n1.val !== n2.val) return false
	return next(n1.left, n2.right) && next(n1.right, n2.left)

}

//isSymmetric(createTreeNode([1, 2, 2, 3, 4, 4, 3]))
