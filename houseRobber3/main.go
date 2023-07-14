package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rob(root *TreeNode) int {
	val0, val1 := robImpl(root)
	return max(val0, val1)

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
//^
//|          4                 4+8+0 ||  9+0        (12,0), (0,0)  
//|       1   null             1+8+0 ||  8+0    =>  (9,8),  (0,0)
//|      1  null              1+0+0  ||  3+5    =>  (1,8),  (0,0)
//|							 root+left(no r)+right(no r)	|| max(l1,l2) + max(r1, r2) 
//|   3  5          	                          =>  (3,0) (5,0)
//                    left(r, no r)       right(r, no r)

// 1 rob the root, sum = val + left(not rob) + right (not rob) no directly rob
// 2 not robbing the root , sum = max of left + max of right 
// 															= max(rob left root, not left root) + max(rob right root, no right root)
func robImpl(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	left0, left1 := robImpl(root.Left)
	right0, right1 := robImpl(root.Right)

	// rob root
	val0 := root.Val + left1 + right1
	// not robbing root
	val1 := max(left0, left1) + max(right0, right1)
	return val0, val1
}
