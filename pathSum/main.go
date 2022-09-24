package main

import "fmt"


func main(){
	t := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: -2,
			Left: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: -1,
				},
				
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: -3,
			Left: &TreeNode{
				Val: -2,
			},
		},
	}
	fmt.Println(hasPathSum(t, -1))
}

//[1,-2,-3,1,3,-2,null,-1]
//           1
//		-2      -3
//	   1  3    -2 nil
//	-1

type TreeNode struct {
    Val int
    Left *TreeNode
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
 func hasPathSum(root *TreeNode, targetSum int) bool {
    return pathSum(root, targetSum)
}

func pathSum(root *TreeNode, targetSum int) bool {
    if root == nil {
        return false
    }
    
    rest := targetSum - root.Val

    if rest == 0 && root.Left == nil && root.Right == nil{
            return true
    }
    return pathSum(root.Left, rest) || pathSum(root.Right, rest)

}