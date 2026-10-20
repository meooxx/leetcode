import { TreeNode } from '../types.ts'


function isValidBST(root: TreeNode | null): boolean {
	const stack: Array<TreeNode | null> = []

	let cur = root
	let pre: TreeNode | null = null
	while (cur != null || stack.length > 0) {

		while (cur != null) {
			stack.push(cur)
			cur = cur.left
		}

		const leftRoot = stack.pop()!
		// pre equals leftChildren root or root
		//    6
		//  4  9
		// 3  8 
		// order is 3 4 6 8 9
		// pre could be 3,4 and then 6,8,9
		if (pre != null && pre.val >= leftRoot.val) {
			return false
		}
		pre = leftRoot
		cur = leftRoot.right


	}


}