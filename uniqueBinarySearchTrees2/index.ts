/**
 * Definition for a binary tree node.
 * class TreeNode {
 *     val: number
 *     left: TreeNode | null
 *     right: TreeNode | null
 *     constructor(val?: number, left?: TreeNode | null, right?: TreeNode | null) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.left = (left===undefined ? null : left)
 *         this.right = (right===undefined ? null : right)
 *     }
 * }
 */

class TreeNode {
	val: number;
	left: TreeNode | null;
	right: TreeNode | null;
	constructor(val?: number, left?: TreeNode | null, right?: TreeNode | null) {
		this.val = val === undefined ? 0 : val;
		this.left = left === undefined ? null : left;
		this.right = right === undefined ? null : right;
	}
}

function generateTrees(n: number): Array<TreeNode | null> {
	const dp: (TreeNode | null)[][] = [];
	dp[0] = [null];

	function copyNode(n: TreeNode | null, offset: number): TreeNode | null {
		if (n == null) return null;
		return new TreeNode(
			n.val + offset,
			copyNode(n.left, offset),
			copyNode(n.right, offset)
		);
	}

	for (let target = 1; target <= n; target++) {
		dp[target] = [];
		for (let root = 1; root <= target; root++) {
			const left = root - 1;
			const right = target - root;
			dp[left].forEach(leftNode => {
				dp[right].forEach(rightNode => {
					dp[target].push(
						new TreeNode(root, leftNode, copyNode(rightNode, root))
					);
				});
			});
		}
	}
	return dp[n];
}

